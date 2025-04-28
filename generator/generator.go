package generator

import (
	"context"

	"github.com/skip-mev/connect-mmu/config"
	"github.com/skip-mev/connect-mmu/generator/filter"
	"github.com/skip-mev/connect-mmu/generator/querier"
	"github.com/skip-mev/connect-mmu/generator/transformer"
	"github.com/skip-mev/connect-mmu/generator/types"
	"github.com/skip-mev/connect-mmu/store/provider"
	mmtypes "github.com/skip-mev/connect/v2/x/marketmap/types"
	"go.uber.org/zap"
)

type Generator struct {
	logger *zap.Logger

	q querier.Querier
	t transformer.Transformer
}

func New(logger *zap.Logger, providerStore provider.Store) Generator {
	return Generator{
		logger: logger.With(zap.String("service", "generator")),
		q:      querier.New(logger, providerStore),
		t:      transformer.New(logger),
	}
}

func (g *Generator) GenerateMarketMap(
	ctx context.Context,
	cfg config.GenerateConfig,
) (mmtypes.MarketMap, types.RemovalReasons, error) {
	feeds, err := g.q.Feeds(ctx, cfg)
	if err != nil {
		g.logger.Error("Unable to query", zap.Error(err))
		return mmtypes.MarketMap{}, nil, err
	}

	g.logger.Info("queried", zap.Int("feeds", len(feeds)))

	transformed, dropped, err := g.t.TransformFeeds(ctx, cfg, feeds)
	if err != nil {
		g.logger.Error("Unable to transform feeds", zap.Error(err))
		return mmtypes.MarketMap{}, nil, err
	}

	g.logger.Info("feed transforms complete", zap.Int("remaining feeds", len(transformed)))

	mm, err := transformed.ToMarketMap()
	if err != nil {
		g.logger.Error("Unable to transform feeds to a MarketMap", zap.Error(err))
		return mmtypes.MarketMap{}, nil, err
	}

	mm, droppedMarkets, err := g.t.TransformMarketMap(ctx, cfg, mm)
	if err != nil {
		g.logger.Error("Unable to transform market map", zap.Error(err))
		return mm, nil, err
	}
	dropped.Merge(droppedMarkets)

	g.logger.Info("apply market map filter")
	list := filter.GetMarketMapList()
	filteredCount := 0
	for key, market := range mm.Markets {
		if _, exists := list[market.Ticker.CurrencyPair.Base]; !exists {
			g.logger.Debug("filtering out market",
				zap.String("ticker", key),
				zap.String("base", market.Ticker.CurrencyPair.Base),
				zap.String("quote", market.Ticker.CurrencyPair.Quote))
			delete(mm.Markets, key)
			filteredCount++
		}
	}

	g.logger.Info("market filtering complete", zap.Int("filtered_out", filteredCount))
	g.logger.Info("final market", zap.Int("size", len(mm.Markets)))

	return mm, dropped, nil
}
