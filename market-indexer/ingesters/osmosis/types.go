package osmosis

import (
	"encoding/json"
	"fmt"
	"github.com/skip-mev/connect-mmu/store/provider"
	"github.com/skip-mev/connect/v2/providers/apis/defi/osmosis"
	"strconv"
	"time"
)

//"data": [
//	{
//		"quote": [
//			{
//				"convert_id": "2781",
//				"price": 83697.6129195127,
//				"price_by_quote_asset": 83636.35363536354,
//				"last_updated": "2025-03-19T12:18:28.452Z",
//				"volume_24h": 1521246.106461655,
//				"percent_change_price_1h": 0.0041188775,
//				"percent_change_price_24h": 0.0147362063,
//				"liquidity": 1373022.4966304917,
//				"fully_diluted_value": null
//			}
//		],
//		"scroll_id": "50",
//		"contract_address": "1943",
//		"name": "BTC/USDC",
//		"base_asset_id": "26291295",
//		"base_asset_ucid": null,
//		"base_asset_name": "Bitcoin",
//		"base_asset_symbol": "BTC",
//		"base_asset_contract_address": "factory/osmo1z6r6qdknhgsc0zeracktgpcxf43j6sekq07nw8sxduc9lg0qjjlqfu25e3/alloyed/allBTC",
//		"quote_asset_id": "23631603",
//		"quote_asset_ucid": "3408",
//		"quote_asset_name": "USDC",
//		"quote_asset_symbol": "USDC",
//		"quote_asset_contract_address": "ibc/498A0751C798A0D9A389AA3691123DADA57DAA4FE165D5C75894505B876BA6E4",
//		"dex_id": "1447",
//		"dex_slug": "osmosis",
//		"network_id": "90",
//		"network_slug": "Osmosis",
//		"last_updated": "2025-03-19T12:18:28.452Z",
//		"created_at": "2024-07-29T13:59:01.000Z"
//	},
//	{
//		"quote": [
//			{
//				"convert_id": "2781",
//				"price": 4.641623604826516,
//				"price_by_quote_asset": 4.638226344983575,
//				"last_updated": "2025-03-19T12:18:21.781Z",
//				"volume_24h": 1461199.7787608455,
//				"percent_change_price_1h": 0.0077219809,
//				"percent_change_price_24h": -0.0216360703,
//				"liquidity": 189722.7194721077,
//				"fully_diluted_value": null
//			}
//		],
//		"scroll_id": "50",
//		"contract_address": "1282",
//		"name": "ATOM/USDC",
//		"base_asset_id": "22222825",
//		"base_asset_ucid": "3794",
//		"base_asset_name": "Cosmos Hub",
//		"base_asset_symbol": "ATOM",
//		"base_asset_contract_address": "ibc/27394FB092D2ECCD56123C74F36E4C1F926001CEADA9CA97EA622B25F41E5EB2",
//		"quote_asset_id": "23631603",
//		"quote_asset_ucid": "3408",
//		"quote_asset_name": "USDC",
//		"quote_asset_symbol": "USDC",
//		"quote_asset_contract_address": "ibc/498A0751C798A0D9A389AA3691123DADA57DAA4FE165D5C75894505B876BA6E4",
//		"dex_id": "1447",
//		"dex_slug": "osmosis",
//		"network_id": "90",
//		"network_slug": "Osmosis",
//		"last_updated": "2025-03-19T12:18:21.781Z",
//		"created_at": "2023-11-10T18:06:40.000Z"
//	}
//]

type OsmosisMarketResponse struct {
	Data   []OsmosisMarketData `json:"data"`
	Status Status              `json:"status"`
}

type OsmosisMarketData struct {
	Quote                     []OsmosisMarketQuote `json:"quote"`
	ScrollID                  string               `json:"scroll_id"`
	ContractAddress           string               `json:"contract_address"`
	Name                      string               `json:"name"`
	BaseAssetID               string               `json:"base_asset_id"`
	BaseAssetUcid             *string              `json:"base_asset_ucid"`
	BaseAssetName             string               `json:"base_asset_name"`
	BaseAssetSymbol           string               `json:"base_asset_symbol"`
	BaseAssetContractAddress  string               `json:"base_asset_contract_address"`
	QuoteAssetID              string               `json:"quote_asset_id"`
	QuoteAssetUcid            *string              `json:"quote_asset_ucid"`
	QuoteAssetName            string               `json:"quote_asset_name"`
	QuoteAssetSymbol          string               `json:"quote_asset_symbol"`
	QuoteAssetContractAddress string               `json:"quote_asset_contract_address"`
	DexID                     string               `json:"dex_id"`
	DexSlug                   string               `json:"dex_slug"`
	NetworkID                 string               `json:"network_id"`
	NetworkSlug               string               `json:"network_slug"`
	LastUpdated               time.Time            `json:"last_updated"`
	CreatedAt                 time.Time            `json:"created_at"`
}

type OsmosisMarketQuote struct {
	ConvertID             string      `json:"convert_id"`
	Price                 float64     `json:"price"`
	PriceByQuoteAsset     float64     `json:"price_by_quote_asset"`
	LastUpdated           time.Time   `json:"last_updated"`
	Volume24H             float64     `json:"volume_24h"`
	PercentChangePrice1H  float64     `json:"percent_change_price_1h"`
	PercentChangePrice24H float64     `json:"percent_change_price_24h"`
	Liquidity             float64     `json:"liquidity"`
	FullyDilutedValue     interface{} `json:"fully_diluted_value"`
}

type Status struct {
	Timestamp    time.Time `json:"timestamp"`
	ErrorCode    string    `json:"error_code"`
	ErrorMessage string    `json:"error_message"`
	Elapsed      int       `json:"elapsed"`
	CreditCount  int       `json:"credit_count"`
}

func (o *OsmosisMarketData) toProvideMarket() (provider.CreateProviderMarket, error) {
	contractAddress, _ := strconv.ParseUint(o.ContractAddress, 10, 64)
	metaData := osmosis.TickerMetadata{
		PoolID:          contractAddress,
		BaseTokenDenom:  o.BaseAssetContractAddress,
		QuoteTokenDenom: o.QuoteAssetContractAddress,
	}

	metaDataBz, err := json.Marshal(metaData)
	if err != nil {
		return provider.CreateProviderMarket{}, fmt.Errorf("osmosis client: failed to marshal metadata: %w", err)
	}

	providerMarket := provider.CreateProviderMarket{
		Create: provider.CreateProviderMarketParams{
			TargetBase:       o.BaseAssetSymbol,
			TargetQuote:      o.QuoteAssetSymbol,
			OffChainTicker:   o.Name,
			ProviderName:     ProviderName,
			QuoteVolume:      o.Quote[0].Volume24H,
			MetadataJSON:     metaDataBz,
			ReferencePrice:   o.Quote[0].Price,
			NegativeDepthTwo: o.Quote[0].Liquidity / 2,
			PositiveDepthTwo: o.Quote[0].Liquidity / 2,
		},
		BaseAddress:  o.BaseAssetContractAddress,
		QuoteAddress: o.QuoteAssetContractAddress,
	}

	return providerMarket, nil
}
