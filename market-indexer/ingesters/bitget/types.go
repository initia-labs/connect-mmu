package bitget

import (
	"fmt"
	"github.com/skip-mev/connect-mmu/lib/symbols"
	"github.com/skip-mev/connect-mmu/store/provider"
	"strconv"
)

//{
//	"code": "00000",
//	"msg": "success",
//	"requestTime": 1745502012392,
//	"data": [
//		{
//			"open": "93706.78",
//			"symbol": "BTCUSDT",
//			"high24h": "94207.48",
//			"low24h": "91666.67",
//			"lastPr": "92889.08",
//			"quoteVolume": "873656389.988141",
//			"baseVolume": "9401.942602",
//			"usdtVolume": "873656389.98814094",
//			"ts": "1745502010943",
//			"bidPr": "92889.08",
//			"askPr": "92889.09",
//			"bidSz": "0.387634",
//			"askSz": "1.340655",
//			"openUtc": "93683.73",
//			"changeUtc24h": "-0.00848",
//			"change24h": "-0.00873"
//		},
//		{
//			"open": "0.03",
//			"symbol": "INITUSDT",
//			"high24h": "0.8",
//			"low24h": "0.03",
//			"lastPr": "0.60332",
//			"quoteVolume": "38776551.6",
//			"baseVolume": "60696049.03",
//			"usdtVolume": "38776551.5967272",
//			"ts": "1745502010782",
//			"bidPr": "0.60255",
//			"askPr": "0.60332",
//			"bidSz": "353.14",
//			"askSz": "4771.97",
//			"openUtc": "0.03",
//			"changeUtc24h": "19.11067",
//			"change24h": "19.11067"
//		}, ...
//	]
//}

type Response struct {
	Code        string `json:"code"`
	Msg         string `json:"msg"`
	RequestTime int64  `json:"requestTime"`
}

type BitgetTickersResponse struct {
	Response
	Data []TickerData `json:"data"`
}

type TickerData struct {
	Open         string `json:"open"`
	Symbol       string `json:"symbol"`
	High24H      string `json:"high24h"`
	Low24H       string `json:"low24h"`
	LastPr       string `json:"lastPr"`
	QuoteVolume  string `json:"quoteVolume"`
	BaseVolume   string `json:"baseVolume"`
	UsdtVolume   string `json:"usdtVolume"`
	Ts           string `json:"ts"`
	BidPr        string `json:"bidPr"`
	AskPr        string `json:"askPr"`
	BidSz        string `json:"bidSz"`
	AskSz        string `json:"askSz"`
	OpenUtc      string `json:"openUtc"`
	ChangeUtc24H string `json:"changeUtc24h"`
	Change24H    string `json:"change24h"`
}

//{
//	"code": "00000",
//	"msg": "success",
//	"requestTime": 1745506794688,
//	"data": [
//		{
//			"symbol": "BTCUSDT",
//			"baseCoin": "BTC",
//			"quoteCoin": "USDT",
//			"minTradeAmount": "0",
//			"maxTradeAmount": "900000000000000000000",
//			"takerFeeRate": "0.002",
//			"makerFeeRate": "0.002",
//			"pricePrecision": "2",
//			"quantityPrecision": "6",
//			"quotePrecision": "8",
//			"status": "online",
//			"minTradeUSDT": "1",
//			"buyLimitPriceRatio": "0.05",
//			"sellLimitPriceRatio": "0.05",
//			"areaSymbol": "no",
//			"orderQuantity": "200",
//			"openTime": "1532454360000",
//			"offTime": ""
//		},
//		{
//			"symbol": "INITUSDT",
//			"baseCoin": "INIT",
//			"quoteCoin": "USDT",
//			"minTradeAmount": "0",
//			"maxTradeAmount": "900000000000000000000",
//			"takerFeeRate": "0.001",
//			"makerFeeRate": "0.001",
//			"pricePrecision": "5",
//			"quantityPrecision": "2",
//			"quotePrecision": "7",
//			"status": "online",
//			"minTradeUSDT": "1",
//			"buyLimitPriceRatio": "0.02",
//			"sellLimitPriceRatio": "0.02",
//			"areaSymbol": "no",
//			"orderQuantity": "200",
//			"openTime": "1745492400000",
//			"offTime": ""
//		},
//	]
//}

type BitgetSymbolsResponse struct {
	Response
	SymbolData []SymbolData `json:"data"`
}

type SymbolData struct {
	Symbol              string `json:"symbol"`
	BaseCoin            string `json:"baseCoin"`
	QuoteCoin           string `json:"quoteCoin"`
	MinTradeAmount      string `json:"minTradeAmount"`
	MaxTradeAmount      string `json:"maxTradeAmount"`
	TakerFeeRate        string `json:"takerFeeRate"`
	MakerFeeRate        string `json:"makerFeeRate"`
	PricePrecision      string `json:"pricePrecision"`
	QuantityPrecision   string `json:"quantityPrecision"`
	QuotePrecision      string `json:"quotePrecision"`
	Status              string `json:"status"`
	MinTradeUSDT        string `json:"minTradeUSDT"`
	BuyLimitPriceRatio  string `json:"buyLimitPriceRatio"`
	SellLimitPriceRatio string `json:"sellLimitPriceRatio"`
	AreaSymbol          string `json:"areaSymbol"`
	OrderQuantity       string `json:"orderQuantity"`
	OpenTime            string `json:"openTime"`
	OffTime             string `json:"offTime"`
}

func (tr *BitgetTickersResponse) toMap() map[string]TickerData {
	m := make(map[string]TickerData, len(tr.Data))

	for _, data := range tr.Data {
		m[data.Symbol] = data
	}

	return m
}

func (sd *SymbolData) toProviderMarket(tickerData TickerData) (provider.CreateProviderMarket, error) {
	quoteVol, err := strconv.ParseFloat(tickerData.QuoteVolume, 64)
	if err != nil {
		return provider.CreateProviderMarket{}, err
	}
	refPrice, err := strconv.ParseFloat(tickerData.LastPr, 64)
	if err != nil {
		return provider.CreateProviderMarket{}, fmt.Errorf("failed to convert lastPr: %w", err)
	}
	targetBase, err := symbols.ToTickerString(sd.BaseCoin)
	if err != nil {
		return provider.CreateProviderMarket{}, err
	}
	targetQuote, err := symbols.ToTickerString(sd.QuoteCoin)
	if err != nil {
		return provider.CreateProviderMarket{}, err
	}

	pm := provider.CreateProviderMarket{
		Create: provider.CreateProviderMarketParams{
			TargetBase:     targetBase,
			TargetQuote:    targetQuote,
			OffChainTicker: sd.Symbol,
			ProviderName:   ProviderName,
			QuoteVolume:    quoteVol,
			ReferencePrice: refPrice,
		},
	}
	return pm, pm.ValidateBasic()
}
