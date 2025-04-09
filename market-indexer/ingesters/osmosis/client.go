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
			return response, err
		}

		var pageresponse OsmosisMarketResponse
		if err := json.NewDecoder(resp.Body).Decode(&pageresponse); err != nil {
			return response, err
		}
		defer resp.Body.Close()

		allData = append(allData, pageresponse.Data...)

		if len(pageresponse.Data) < limit {
			break
		}

		scrollID = pageresponse.Data[len(pageresponse.Data)-1].ScrollID
	}

	response.Data = allData

	return response, nil
}
