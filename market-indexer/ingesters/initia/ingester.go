package initia

import (
	"context"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"

	"github.com/skip-mev/connect-mmu/lib/symbols"
	"github.com/skip-mev/connect-mmu/market-indexer/ingesters"
	"github.com/skip-mev/connect-mmu/market-indexer/ingesters/types"
	"github.com/skip-mev/connect-mmu/store/provider"
	"github.com/skip-mev/connect/v2/providers/apis/defi/initia"
)

const (
	Name         = "initia"
	ProviderName = Name + types.ProviderNameSuffixAPI
)

var _ ingesters.Ingester = &Ingester{}

type Ingester struct {
	logger *zap.Logger
	client Client
}

func (i *Ingester) Name() string {
	return Name
}

func New(logger *zap.Logger) *Ingester {
	if logger == nil {
		panic("cannot set nil logger")
	}
	return &Ingester{
		logger: logger.With(zap.String("ingester", Name)),
		client: NewHttpClient(logger),
	}
}

func (i *Ingester) GetProviderMarkets(ctx context.Context) ([]provider.CreateProviderMarket, error) {
	i.logger.Info("fetching initia dex pools")
	pools, err := i.client.Pools(ctx)
	if err != nil {
		i.logger.Error("failed to fetch initia dex pools", zap.Error(err))
		return nil, err
	}

	i.logger.Info("fetching initia dex registries")
	registries, err := i.client.Registries(ctx)
	if err != nil {
		i.logger.Error("failed to fetch initia dex registries", zap.Error(err))
		return nil, err
	}

	denomToSymbol := registries.toMap()
	markets := make([]provider.CreateProviderMarket, 0, len(pools.Pools))

	for _, pool := range pools.Pools {
		if len(pool.Coins) != 2 {
			continue
		}

		baseDenom := pool.Coins[0].Denom
		quoteDenom := pool.Coins[1].Denom

		baseSymbol, findBase := denomToSymbol[baseDenom]
		quoteSymbol, findQuote := denomToSymbol[quoteDenom]
		if !findBase || !findQuote {
			i.logger.Error("failed to find symbol for denom", zap.String("base", baseDenom), zap.String("quote", quoteDenom))
			continue
		}

		targetBase, err := symbols.ToTickerString(baseSymbol)
		if err != nil {
			i.logger.Error("failed to convert symbol to ticker string", zap.String("base symbol", baseSymbol), zap.Error(err))
			continue
		}
		targetQuote, err := symbols.ToTickerString(quoteSymbol)
		if err != nil {
			i.logger.Error("failed to convert symbol to ticker string", zap.String("quote symbol", quoteSymbol), zap.Error(err))
			continue
		}

		metadata := initia.InitiaMetadata{
			LPDenom:         pool.Lp,
			BaseTokenDenom:  baseDenom,
			QuoteTokenDenom: quoteDenom,
		}

		metaDataBz, err := json.Marshal(metadata)
		if err != nil {
			i.logger.Error("failed to marshal metadata", zap.Error(err))
			continue
		}

		pm := provider.CreateProviderMarket{
			Create: provider.CreateProviderMarketParams{
				TargetBase:     targetBase,
				TargetQuote:    targetQuote,
				OffChainTicker: fmt.Sprintf("%s/%s", targetBase, targetQuote),
				ProviderName:   ProviderName,
				MetadataJSON:   metaDataBz,
				QuoteVolume:    pool.Volume24H,
				ReferencePrice: 1,
			},
			BaseAddress:  baseDenom,
			QuoteAddress: quoteDenom,
		}

		if err := pm.ValidateBasic(); err != nil {
			i.logger.Debug("failed to validate provider market", zap.Error(err))
			continue
		}

		markets = append(markets, pm)
	}

	return markets, nil
}
