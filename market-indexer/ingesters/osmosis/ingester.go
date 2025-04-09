package osmosis

import (
	"context"
	"github.com/skip-mev/connect-mmu/config"
	"github.com/skip-mev/connect-mmu/market-indexer/ingesters"
	"github.com/skip-mev/connect-mmu/market-indexer/ingesters/types"
	"github.com/skip-mev/connect-mmu/store/provider"
	"go.uber.org/zap"
)

const (
	Name         = "osmosis"
	ProviderName = Name + types.ProviderNameSuffixAPI
)

var _ ingesters.Ingester = &Ingester{}

type Ingester struct {
	logger *zap.Logger
	client Client
}

func New(logger *zap.Logger, marketConfig config.MarketConfig) *Ingester {
	ing := &Ingester{
		logger: logger.With(zap.String("ingester", Name)),
		client: newClient(logger, marketConfig.CoinMarketCapConfig.APIKey),
	}
	return ing
}

func (ing *Ingester) Name() string {
	return Name
}

func (i *Ingester) GetProviderMarkets(ctx context.Context) ([]provider.CreateProviderMarket, error) {
	i.logger.Info("fetching osmosis market data")
	marketPairs, err := i.client.OsmosisMarkets(ctx)
	if err != nil {
		i.logger.Error("failed to fetch osmosis market data", zap.Error(err))
		return nil, err
	}

	providerMarkets := make([]provider.CreateProviderMarket, 0, len(marketPairs.Data))
	for _, providerMarket := range marketPairs.Data {
		providerMarket, err := providerMarket.toProvideMarket()
		if err != nil {
			i.logger.Error("failed to convert provider market", zap.Error(err))
			return nil, err
		}

		providerMarkets = append(providerMarkets, providerMarket)
	}
	
	i.logger.Info("fetched osmosis market data", zap.Int("markets", len(providerMarkets)))
	return providerMarkets, nil
}
