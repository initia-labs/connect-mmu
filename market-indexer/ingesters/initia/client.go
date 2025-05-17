package initia

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/skip-mev/connect-mmu/lib/http"
	"go.uber.org/zap"
	"strconv"
)

const (
	EndpointInitiaPools = "https://dex-api.initia.xyz/indexer/dex/v1/pools"

	EndpointInitiaRegistry = "https://registry.initia.xyz/chains/initia/assetlist.json"
)

var _ Client = &httpClient{}

type Client interface {
	Pools(ctx context.Context) (InitiaDexResponse, error)

	Registries(ctx context.Context) (InitiaRegistryResponse, error)
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

func (h *httpClient) Pools(ctx context.Context) (InitiaDexResponse, error) {
	h.logger.Info("get initia dex pools")

	var response InitiaDexResponse
	var allData []PoolsResponse

	limit := 100
	nextKey := ""

	for {
		opts := []http.GetOptions{
			http.WithJSONAccept(),
			http.WithQueryParam("pagination.limit", fmt.Sprintf("%d", limit)),
			http.WithQueryParam("pagination.count_total", strconv.FormatBool(true)),
			http.WithQueryParam("type", "ALL"),
		}

		if nextKey != "" {
			opts = append(opts, http.WithQueryParam("pagination.key", nextKey))
		}

		resp, err := h.client.GetWithContext(ctx, EndpointInitiaPools, opts...)
		if err != nil {
			h.logger.Error("failed to get initia dex pools", zap.Error(err))
			return response, err
		}
		defer resp.Body.Close()

		var pageResponse InitiaDexResponse
		if err := json.NewDecoder(resp.Body).Decode(&pageResponse); err != nil {
			h.logger.Error("failed to decode initia dex pools", zap.Error(err))
			return response, err
		}

		allData = append(allData, pageResponse.Pools...)

		if pageResponse.Pagination.NextKey == nil {
			break
		}

		switch v := pageResponse.Pagination.NextKey.(type) {
		case string:
			if v == "" {
				return response, fmt.Errorf("pagination.next_key is empty")
			}
			nextKey = v
		default:
			return response, fmt.Errorf("unexpected type for pagination.next_key: %T", v)
		}

		h.logger.Debug("fetching next page with next_key", zap.String("next_key", nextKey))
	}

	response.Pools = allData

	h.logger.Info("fetched initia dex pools", zap.Int("total", len(response.Pools)))
	return response, nil
}

func (h *httpClient) Registries(ctx context.Context) (InitiaRegistryResponse, error) {
	h.logger.Info("get initia dex registries")

	resp, err := h.client.GetWithContext(ctx, EndpointInitiaRegistry)
	if err != nil {
		h.logger.Error("failed to get initia dex registries", zap.Error(err))
		return InitiaRegistryResponse{}, err
	}
	defer resp.Body.Close()

	var registries InitiaRegistryResponse
	if err := json.NewDecoder(resp.Body).Decode(&registries); err != nil {
		h.logger.Error("failed to decode initia dex registries", zap.Error(err))
		return InitiaRegistryResponse{}, err
	}

	return registries, nil
}
