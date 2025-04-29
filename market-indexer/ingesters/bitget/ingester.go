package bitget

import (
	"context"
	"fmt"

	"go.uber.org/zap"

	"github.com/skip-mev/connect-mmu/market-indexer/ingesters"
	"github.com/skip-mev/connect-mmu/market-indexer/ingesters/types"
	"github.com/skip-mev/connect-mmu/store/provider"
)

const (
	Name         = "bitget"
	ProviderName = Name + types.ProviderNameSuffixWS

	StatusOnline = "online"
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
	i.logger.Info("fetching data")

	symbols, err := i.client.Symbols(ctx)
	if err != nil {
		return nil, err
	}

	tickers, err := i.client.Tickers(ctx)
	if err != nil {
		return nil, err
	}

	tickerMap := tickers.toMap()

	pms := make([]provider.CreateProviderMarket, 0, len(symbols.SymbolData))
	for _, item := range symbols.SymbolData {
		if item.Status != StatusOnline {
			continue
		}

		ticker, found := tickerMap[item.Symbol]
		if !found {
			return nil, fmt.Errorf("ticker not found for symbol %s", item.Symbol)
		}

		i.logger.Debug("ticker", zap.Any("data", ticker))

		pm, err := item.toProviderMarket(ticker)
		if err != nil {
			return nil, err
		}
		pms = append(pms, pm)
	}

	return pms, nil
}
