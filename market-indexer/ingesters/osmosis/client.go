package osmosis

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/skip-mev/connect-mmu/lib/http"
	"go.uber.org/zap"
)

const (
	EndpointOsmosisMarkets = "https://pro-api.coinmarketcap.com/v4/dex/spot-pairs/latest?dex_id=1447&network_id=90"
)

var _ Client = &httpClient{}

type Client interface {
	OsmosisMarkets(ctx context.Context) (OsmosisMarketResponse, error)
}

type httpClient struct {
	client *http.Client
	logger *zap.Logger
	apiKey string
}

func newClient(logger *zap.Logger, apiKey string) Client {
	return &httpClient{
		client: http.NewClient(),
		logger: logger,
		apiKey: apiKey,
	}
}

func (c *httpClient) OsmosisMarkets(ctx context.Context) (OsmosisMarketResponse, error) {
	c.logger.Info("get osmosis market data")

	var response OsmosisMarketResponse
	var allData []OsmosisMarketData

	limit := 50
	var scrollID string

	for {
		opts := []http.GetOptions{
			http.WithHeader("X-CMC_PRO_API_KEY", c.apiKey),
			http.WithJSONAccept(),
			http.WithQueryParam("limit", fmt.Sprintf("%d", limit)),
			http.WithQueryParam("scroll_id", scrollID),
		}

		resp, err := c.client.GetWithContext(ctx, EndpointOsmosisMarkets, opts...)
		if err != nil {
			c.logger.Error("failed to get osmosis market data", zap.Error(err))
			return response, err
		}
		defer resp.Body.Close()

		var pageresponse OsmosisMarketResponse
		if err := json.NewDecoder(resp.Body).Decode(&pageresponse); err != nil {
			c.logger.Error("failed to decode osmosis market response", zap.Error(err))
			return response, err
		}

		allData = append(allData, pageresponse.Data...)

		if len(pageresponse.Data) < limit {
			break
		}

		scrollID = pageresponse.Data[len(pageresponse.Data)-1].ScrollID
		c.logger.Debug("fetching next page with scroll ID", zap.String("scroll_id", scrollID))
	}

	response.Data = allData
	c.logger.Info("fetched osmosis market data", zap.Int("total_markets", len(allData)))

	return response, nil
}
