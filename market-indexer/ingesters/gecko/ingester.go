package gecko

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"go.uber.org/zap"
	"gopkg.in/typ.v4/maps"

	"github.com/skip-mev/connect-mmu/config"
	"github.com/skip-mev/connect-mmu/market-indexer/ingesters"
	"github.com/skip-mev/connect-mmu/market-indexer/ingesters/types"
	"github.com/skip-mev/connect-mmu/store/provider"
	"github.com/skip-mev/connect/v2/providers/apis/defi/curve"
	"github.com/skip-mev/connect/v2/providers/apis/defi/uniswapv3"
	"github.com/skip-mev/connect/v2/providers/apis/geckoterminal"
)

const (
	Name         = "gecko_terminal"
	ProviderName = Name + types.ProviderNameSuffixAPI
	QuoteUSD     = "USD"
)

var _ ingesters.Ingester = &Ingester{}

type Ingester struct {
	logger *zap.Logger
	client Client
	pairs  []config.GeckoNetworkDexPair
}

// New returns a new gecko terminal ingester. Options may be used to specify more networks and dexes to query.
// The default ingester only queries uniswap v3 on the ethereum network.
func New(logger *zap.Logger, marketConfig config.MarketConfig) *Ingester {
	if err := validatePairs(marketConfig.GeckoNetworkDexPairs); err != nil {
		panic("invalid pairs: " + err.Error())
	}
	ing := &Ingester{
		logger: logger.With(zap.String("ingester", Name)),
		pairs:  marketConfig.GeckoNetworkDexPairs,
		client: newClient(logger, BaseEndpoint, marketConfig.CoinGeckoConfig.APIKey),
	}
	return ing
}

func (ig *Ingester) Name() string {
	return Name
}

func (ig *Ingester) GetProviderMarkets(ctx context.Context) ([]provider.CreateProviderMarket, error) {
	ig.logger.Info("fetching data")

	providerMarkets := make([]provider.CreateProviderMarket, 0)

	// for each network+dex pair:
	for _, pair := range ig.pairs {
		// get the top pools in that network + dex.
		pair.Dex = ConvertDexName(pair.Dex)
		pools, err := ig.topPools(ctx, pair.Network, pair.Dex)
		if err != nil {
			return nil, err
		}

		ig.logger.Info("fetched data", zap.Int("pools", len(pools)), zap.String("dex", pair.Dex))

		// extract a set of tokens from the top pools.
		tokenSet := make(map[string]struct{})
		for _, pool := range pools {
			base, quote := pool.GetBaseAndQuoteTokenAddress()
			tokenSet[base] = struct{}{}
			tokenSet[quote] = struct{}{}
		}

		// turn the set of unique tokens into a slice.
		tokens := maps.NewSetFromKeys(tokenSet).Slice()

		// get the token data for each of these tokens.
		tokensRes, err := ig.client.MultiToken(ctx, pair.Network, tokens)
		if err != nil {
			return nil, err
		}
		// make a mapping of token address -> tokenData
		tokensData := make(map[string]TokenData)
		for _, tokenData := range tokensRes.Data {
			tokensData[tokenData.Attributes.Address] = tokenData
		}

		// map of venue address -> Base
		type venueInfo struct {
			baseAddress string
			targetBase  string
		}
		venueAddresses := make(map[string]venueInfo)

		// iterate over all the pools, and create provider market params using the token + pool data.
		for _, pool := range pools {
			baseData := tokensData[pool.BaseAddress()]
			quoteData := tokensData[pool.QuoteAddress()]
			quoteVol, err := pool.QuoteVolume()
			if err != nil {
				return nil, fmt.Errorf("gecko client: failed to get quote volume for pool %q, quote %q: %w", pool.ID, quoteData.Attributes.Symbol, err)
			}
			quoteVolF64, _ := quoteVol.Float64()
			if !isValidFloat64(quoteVolF64) {
				ig.logger.Debug(
					"gecko client: unable to fit quote volume in float64",
					zap.String("quote", quoteData.Attributes.Symbol),
					zap.String("pool", pool.Attributes.Name),
					zap.String("quote volume", quoteVol.String()),
				)
				continue
			}

			// uniswap sets the token order based on the token addresses.
			// that means, pools may appear as 0xFOO/0xBAR, however, since sorting these strings would place
			// BAR first, that means the pool will be priced as 0xBAR/0xFOO.
			//
			// so if the base is greater(1) than the quote (meaning quote goes first in sorted order), we know we must invert.
			//
			// TODO: we currently need to do the opposite of the above, however, as there is a bug in Connect's uniswap code.
			// it will actually invert the price when invert == false, and not invert it when invert == true.
			invert := strings.Compare(pool.BaseAddress(), pool.QuoteAddress()) == 1

			metaData := MakeMetadata(pair.Dex)
			if metaData == nil {
				ig.logger.Debug("unsupported DEX type", zap.String("dex", pair.Dex))
				continue
			}

			switch m := metaData.(type) {
			case uniswapv3.PoolConfig:
				m.Address = pool.VenueAddress()
				m.BaseDecimals = int64(baseData.Decimals())
				m.QuoteDecimals = int64(quoteData.Decimals())
				m.Invert = invert
				metaData = m
			case curve.CurveMetadata:
				m.Network = MakeNetwork(pair.Network)
				m.PoolID = pool.VenueAddress()
				m.BaseTokenAddress = pool.BaseAddress()
				m.QuoteTokenAddress = pool.QuoteAddress()
				metaData = m
			}

			metaDataBz, err := json.Marshal(metaData)
			if err != nil {
				return nil, fmt.Errorf("gecko client: failed to marshal metadata: %w", err)
			}
			refPrice, err := pool.ReferencePrice()
			if err != nil {
				ig.logger.Debug("failed to get reference price for pool", zap.Error(err))
				continue
			}

			targetBase, err := pool.Base()
			if err != nil {
				ig.logger.Debug("failed to convert target base to ticker string - skipping", zap.Error(err))
				continue
			}

			targetQuote, err := pool.Quote()
			if err != nil {
				ig.logger.Debug("failed to convert target quote to ticker string - skipping", zap.Error(err))
				continue
			}

			offChainTicker, err := pool.OffChainTicker()
			if err != nil {
				ig.logger.Debug("gecko client: failed to convert off chain ticker to ticker string", zap.Error(err))
				continue
			}

			market := provider.CreateProviderMarket{
				Create: provider.CreateProviderMarketParams{
					TargetBase:     targetBase,
					TargetQuote:    targetQuote,
					OffChainTicker: offChainTicker,
					ProviderName:   geckoDexToConnectDex(pool.Venue()),
					QuoteVolume:    quoteVolF64,
					MetadataJSON:   metaDataBz,
					ReferencePrice: refPrice,
				},
				BaseAddress:  pool.BaseAddress(),
				QuoteAddress: pool.QuoteAddress(),
			}
			providerMarkets = append(providerMarkets, market)

			// add the venue address to the map if it doesn't exist.
			if _, ok := venueAddresses[pair.Network]; !ok {
				venueAddresses[pool.VenueAddress()] = struct {
					baseAddress string
					targetBase  string
				}{
					baseAddress: pool.BaseAddress(),
					targetBase:  targetBase,
				}
			}
		}

		// add gecko_terminal_api provider markets
		for _, info := range venueAddresses {
			metaData := geckoterminal.GeckoterminalMetadata{
				Network: pair.Network,
			}
			metaDataBz, err := json.Marshal(metaData)
			if err != nil {
				return nil, fmt.Errorf("gecko client: failed to marshal metadata: %w", err)
			}
			providerMarkets = append(providerMarkets, provider.CreateProviderMarket{
				Create: provider.CreateProviderMarketParams{
					TargetBase:     info.targetBase,
					TargetQuote:    QuoteUSD,
					OffChainTicker: info.baseAddress,
					ProviderName:   ProviderName,
					MetadataJSON:   metaDataBz,
					QuoteVolume:    0,
					ReferencePrice: 0,
				},
				BaseAddress:  info.baseAddress,
				QuoteAddress: "",
			})
		}
	}
	ig.logger.Info("fetched data", zap.Int("markets", len(providerMarkets)))

	return providerMarkets, nil
}

// topPools simply calls the TopPools method, facilitating the pagination.
func (ig *Ingester) topPools(ctx context.Context, network, dex string) ([]PoolData, error) {
	poolsData := make([]PoolData, 0)
	for i := 1; i <= maxPages; i++ {
		pools, err := ig.client.TopPools(ctx, network, dex, i)
		if err != nil {
			return nil, fmt.Errorf("failed to get pools: page %d: %w", i, err)
		}
		poolsData = append(poolsData, pools.Data...)
	}
	return poolsData, nil
}
