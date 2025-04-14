package gecko

// TickerMap defines the mapping from targetBase to off-chain ticker
var TickerMap = map[string]string{
	"ATOM":    "cosmos",
	"OSMO":    "osmosis",
	"STARS":   "stargaze",
	"NTRN":    "neutron-3",
	"TIA":     "celestia",
	"MILKTIA": "milkyway-staked-tia",
	"BNB":     "binancecoin",
	"BTC":     "bitcoin",
	"LBTC":    "lombard-staked-btc",
	"ETH":     "ethereum",
	"WETH":    "weth",
	"WEETH":   "wrapped-eeth",
	"ETHFI":   "ether-fi",
	"USDT":    "tether",
	"USDC":    "usd-coin",
	"SUSDE":   "ethena-staked-usde",
	"USDE":    "ethena-usde",
	"ENA":     "ethena",
	"HYPE":    "hyperliquid",
	"BERA":    "berachain-bera",
	"SOL":     "solana",
	"SUI":     "sui",
	"APT":     "aptos",
}

// ConvertTicker takes a targetBase and returns the corresponding coingecko API ID with /usd (vs_currencies).
func ConvertTicker(targetBase string) (string, bool) {
	ticker, exists := TickerMap[targetBase]
	return ticker + "/usd", exists
}

func ConvertDexName(dex string) string {
	// Define a map to hold the conversions
	dexMap := map[string]string{
		"uniswap-v4":    "uniswap-v4-ethereum",
		"curve-finance": "curve",
	}

	// Check if the given dex name exists in the map and return the mapped value
	// If it doesn't exist, return the original dex name
	if converted, exists := dexMap[dex]; exists {
		return converted
	}
	return dex
}
