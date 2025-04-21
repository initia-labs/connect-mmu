package transformer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/skip-mev/connect-mmu/market-indexer/ingesters/gecko"
	"github.com/skip-mev/connect/v2/providers/apis/defi/curve"
	"slices"
	"strings"

	connecttypes "github.com/skip-mev/connect/v2/pkg/types"
	mmtypes "github.com/skip-mev/connect/v2/x/marketmap/types"
	"go.uber.org/zap"

	"github.com/skip-mev/connect-mmu/config"
	"github.com/skip-mev/connect-mmu/generator/filter"
	"github.com/skip-mev/connect-mmu/generator/types"
)

// TransformMarketMap is a function that performs some transformation on a marketmap.
type TransformMarketMap func(ctx context.Context, logger *zap.Logger, cfg config.GenerateConfig, mm mmtypes.MarketMap) (mmtypes.MarketMap, types.RemovalReasons, error)

// OverrideMarkets applies the GenerateConfig's market overrides. Note: if there is no market to override, this function
// will add the market to the market map.
func OverrideMarkets() TransformMarketMap {
	return func(_ context.Context, logger *zap.Logger, cfg config.GenerateConfig, mm mmtypes.MarketMap) (mmtypes.MarketMap, types.RemovalReasons, error) {
		logger.Info("overriding MarketMap")
		for name, market := range cfg.MarketMapOverride.Markets {
			logger.Info("overriding market", zap.String("name", name))
			mm.Markets[name] = market
		}

		logger.Info("market size after overrides", zap.Int("size", len(mm.Markets)))
		return mm, nil, nil
	}
}

// OverrideMinProviderCount will wholesale replace the MinProviderCount value for each Market's Ticker.
// This would be run before the OverrideMarkets transform so that specific MinProviderCount values in overridden markets
// are preserved.
func OverrideMinProviderCount() TransformMarketMap {
	return func(_ context.Context, logger *zap.Logger, cfg config.GenerateConfig, mm mmtypes.MarketMap) (mmtypes.MarketMap, types.RemovalReasons, error) {
		if cfg.MinProviderCountOverride == 0 {
			return mm, nil, nil
		}
		logger.Info("overriding min provider count")
		for name, market := range mm.Markets {
			minCount := filter.MinProvidersThreshold
			if minCount < len(market.ProviderConfigs) {
				market.Ticker.MinProviderCount = uint64(minCount)
			} else {
				market.Ticker.MinProviderCount = uint64(len(market.ProviderConfigs))
			}
			mm.Markets[name] = market
		}
		return mm, nil, nil
	}
}

// PruneInsufficientlyProvidedMarkets removes markets that did not have the minimum amount of providers.
func PruneInsufficientlyProvidedMarkets() TransformMarketMap {
	return func(_ context.Context, logger *zap.Logger, cfg config.GenerateConfig, mm mmtypes.MarketMap) (mmtypes.MarketMap, types.RemovalReasons, error) {
		logger.Info("pruning insufficiently provided markets")

		removals := types.NewRemovalReasons()
		for key, market := range mm.Markets {

			providers := uint64(0)
			for _, provider := range market.ProviderConfigs {
				providerConfig := cfg.Providers[provider.Name]
				if !providerConfig.IsSupplemental {
					providers++
				}
			}
			if providers < market.Ticker.MinProviderCount {
				logger.Debug("pruning market with insufficient providers", zap.String("name", key))
				providerNames := make([]string, len(market.ProviderConfigs))
				for i, providerConfig := range market.ProviderConfigs {
					providerNames[i] = providerConfig.Name
				}
				removals.AddRemovalReasonFromMarket(market, market.Ticker.CurrencyPair.String(),
					fmt.Sprintf("PruneInsufficientlyProvidedMarkets: insufficient # of providers: %s, min: %d", strings.Join(providerNames, ","), market.Ticker.MinProviderCount))
				delete(mm.Markets, key)
			}
		}

		logger.Info("market size after pruning insufficiently provided markets", zap.Int("size", len(mm.Markets)))
		return mm, removals, nil
	}
}

// PruneMarkets removes currency pairs that are not allowed in the configuration. This is decided by the
// ExcludeCurrencyPairs set, and the AllowedCurrencyPairs set. See method IsCurrencyPairAllowed for more details.
func PruneMarkets() TransformMarketMap {
	return func(_ context.Context, logger *zap.Logger, cfg config.GenerateConfig, mm mmtypes.MarketMap) (mmtypes.MarketMap, types.RemovalReasons, error) {
		logger.Info("pruning disallowed markets")
		removals := types.NewRemovalReasons()
		for name, market := range mm.Markets {
			if !cfg.IsCurrencyPairAllowed(market.Ticker.CurrencyPair) {
				logger.Debug("removing market", zap.String("name", name))
				removals.AddRemovalReasonFromMarket(market, market.Ticker.CurrencyPair.String(),
					fmt.Sprintf("PruneMarkets: disallowed currency pair: %s", market.Ticker.CurrencyPair.String()))
				delete(mm.Markets, name)
			}
		}

		logger.Info("market size after pruning disallowed markets", zap.Int("size", len(mm.Markets)))
		return mm, removals, nil
	}
}

// RemoveDisabledProviders will remove providers from markets in the marketmap that are specified in the DisableProviders
// field of the GenerateConfig.
func RemoveDisabledProviders() TransformMarketMap {
	return func(_ context.Context, logger *zap.Logger, cfg config.GenerateConfig, mm mmtypes.MarketMap) (mmtypes.MarketMap, types.RemovalReasons, error) {
		removals := types.NewRemovalReasons()
		for ticker, providersToRemove := range cfg.DisableProviders {
			m, ok := mm.Markets[ticker]
			if !ok {
				logger.Debug("RemoveDisabledProviders: market does not exist", zap.String("market", ticker))
				continue
			}
			updatedProviders := make([]mmtypes.ProviderConfig, 0)
			for _, provider := range m.ProviderConfigs {
				// only append to updateProviders is this provider is _NOT_ in the providersToRemove slice.
				if !slices.Contains(providersToRemove, provider.Name) {
					updatedProviders = append(updatedProviders, provider)
				} else {
					logger.Debug("RemoveDisabledProviders: removing provider", zap.String("provider", provider.Name), zap.String("market", ticker))
					removals.AddRemovalReasonFromMarket(m, provider.Name, fmt.Sprintf("RemoveDisabledProviders: provider %q is disabled for market %q", provider.Name, ticker))
				}
			}
			m.ProviderConfigs = updatedProviders
			mm.Markets[ticker] = m
		}
		return mm, removals, nil
	}
}

// EnableMarkets enabled markets based on the GenerateConfig rules.
func EnableMarkets() TransformMarketMap {
	return func(_ context.Context, logger *zap.Logger, cfg config.GenerateConfig,
		mm mmtypes.MarketMap,
	) (mmtypes.MarketMap, types.RemovalReasons, error) {
		if cfg.EnableAll {
			logger.Info("enabling all markets")
			for name, market := range mm.Markets {
				market.Ticker.Enabled = true
				mm.Markets[name] = market
			}

			return mm, nil, nil
		}

		return mm, nil, nil
	}
}

// replaceNormalizeBy finds all instances of oldNormalizeBy and replaces them with newNormalizeBy in the marketmap.
func replaceNormalizeBy(mm mmtypes.MarketMap, oldNormalizeBy, newNormalizeBy connecttypes.CurrencyPair) mmtypes.MarketMap {
	for key, market := range mm.Markets {
		for i, pc := range market.ProviderConfigs {
			if pc.NormalizeByPair != nil {
				if pc.NormalizeByPair.Equal(oldNormalizeBy) {
					pc.NormalizeByPair = &newNormalizeBy
				}
			}

			market.ProviderConfigs[i] = pc
		}

		mm.Markets[key] = market
	}

	return mm
}

// ProcessDefiMarkets adds defi tickers to markets that are defi and only have one provider.
func ProcessDefiMarkets() TransformMarketMap {
	return func(_ context.Context, logger *zap.Logger, cfg config.GenerateConfig,
		mm mmtypes.MarketMap,
	) (mmtypes.MarketMap, types.RemovalReasons, error) {
		logger.Info("processing defi markets", zap.Int("num markets", len(mm.Markets)))

		for name, market := range mm.Markets {
			if len(market.ProviderConfigs) == 1 && cfg.IsProviderDefi(market.ProviderConfigs[0].Name) {
				// cache the old ticker string and use to resolve any dangling normalizations
				oldCp := market.Ticker.CurrencyPair

				pc := market.ProviderConfigs[0]
				cp, err := connecttypes.CurrencyPairFromString(pc.OffChainTicker)
				if err != nil {
					logger.Debug("failed to create currency pair", zap.String("offchain tickers", pc.OffChainTicker),
						zap.Error(err))
					delete(mm.Markets, name)
					continue
				}

				// put market under new key
				delete(mm.Markets, name)

				if pc.Invert {
					cp = cp.Invert()
				}

				if pc.NormalizeByPair != nil {
					cp.Quote = pc.NormalizeByPair.Quote
				}

				market.Ticker.CurrencyPair = cp
				mm.Markets[market.Ticker.String()] = market
				mm = replaceNormalizeBy(mm, oldCp, cp)
			}
		}

		logger.Info("processed defi markets", zap.Int("num markets", len(mm.Markets)))
		return mm, nil, nil
	}
}

func ProcessDefiInvert() TransformMarketMap {
	return func(_ context.Context, logger *zap.Logger, cfg config.GenerateConfig,
		mm mmtypes.MarketMap,
	) (mmtypes.MarketMap, types.RemovalReasons, error) {
		logger.Info("processing defi invert", zap.Int("num markets", len(mm.Markets)))
		count := 0
		for name, market := range mm.Markets {
			for i, pc := range market.ProviderConfigs {
				if cfg.IsProviderDefi(pc.Name) && pc.Name == gecko.ProviderNameCurve {
					var metadata curve.CurveMetadata
					if err := json.Unmarshal([]byte(pc.Metadata_JSON), &metadata); err != nil {
						logger.Debug("failed to parse metadata JSON", zap.String("metadata", pc.Metadata_JSON), zap.Error(err))
						continue
					}
					if pc.Invert {
						//newMetadata := curve.CurveMetadata{
						//	Network:           metadata.Network,
						//	PoolID:            metadata.PoolID,
						//	BaseTokenAddress:  metadata.QuoteTokenAddress,
						//	QuoteTokenAddress: metadata.BaseTokenAddress,
						//}
						//if newMetadataJson, err := json.Marshal(newMetadata); err == nil {
						//	market.ProviderConfigs[i].OffChainTicker = metadata.QuoteTokenAddress
						//	market.ProviderConfigs[i].Metadata_JSON = string(newMetadataJson)
						//	logger.Debug("metadata updated",
						//		zap.String("market", name),
						//		zap.String("new_metadata", string(newMetadataJson)))
						//} else {
						//	logger.Error("failed to marshal metadata",
						//		zap.String("market", name),
						//		zap.Error(err))
						//
						//	count++
						//}
						market.ProviderConfigs[i].OffChainTicker = metadata.QuoteTokenAddress
						count++
					} else {
						market.ProviderConfigs[i].OffChainTicker = metadata.BaseTokenAddress
					}

					mm.Markets[name] = market
				}
			}
		}
		logger.Info("processed defi invert", zap.Int("num markets", count))
		return mm, nil, nil
	}
}
