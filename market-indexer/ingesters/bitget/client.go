package bitget

import (
	"context"
	"encoding/json"
	"github.com/skip-mev/connect-mmu/lib/http"
	"go.uber.org/zap"
)

const (
	EndpointSymbols = "https://api.bitget.com/api/v2/spot/public/symbols"

	EndpointTickers = "https://api.bitget.com/api/v2/spot/market/tickers"
)

var _ Client = &httpClient{}

type Client interface {
	// Symbols gets all symbols from Bitget
	Symbols(ctx context.Context) (BitgetSymbolsResponse, error)

	// Tickers gets all tickers from Bitget
	Tickers(ctx context.Context) (BitgetTickersResponse, error)
}

type httpClient struct {
	client *http.Client
	logger *zap.Logger
}

func NewHttpClient(logger *zap.Logger) Client {
	return &httpClient{
		client: http.NewClient(),
		logger: logger,
	}
}

func (h *httpClient) Symbols(ctx context.Context) (BitgetSymbolsResponse, error) {
	resp, err := h.client.GetWithContext(ctx, EndpointSymbols)
	if err != nil {
		h.logger.Error("failed to get symbols", zap.Error(err))
		return BitgetSymbolsResponse{}, err
	}
	defer resp.Body.Close()

	var symbols BitgetSymbolsResponse
	if err := json.NewDecoder(resp.Body).Decode(&symbols); err != nil {
		h.logger.Error("failed to decode symbols", zap.Error(err))
		return BitgetSymbolsResponse{}, err
	}

	return symbols, nil
}

func (h *httpClient) Tickers(ctx context.Context) (BitgetTickersResponse, error) {
	resp, err := h.client.GetWithContext(ctx, EndpointTickers)
	if err != nil {
		h.logger.Error("failed to get tickers", zap.Error(err))
		return BitgetTickersResponse{}, err
	}
	defer resp.Body.Close()

	var tickers BitgetTickersResponse
	if err := json.NewDecoder(resp.Body).Decode(&tickers); err != nil {
		h.logger.Error("failed to decode tickers", zap.Error(err))
		return BitgetTickersResponse{}, err
	}

	return tickers, nil
}
