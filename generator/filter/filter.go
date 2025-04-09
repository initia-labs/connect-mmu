package filter

// GetMarketMapList returns a map of allowed market symbols
func GetMarketMapList() map[string]bool {
	return map[string]bool{
		"ATOM":    true,
		"OSMO":    true,
		"STARS":   true, // Stargaze
		"NTRN":    true,
		"TIA":     true,
		"MILKTIA": true,
		"BNB":     true,
		"BTC":     true,
		"LBTC":    true, // Lombard Staked BTC
		"ETH":     true,
		"WETH":    true,
		"WEETH":   true,
		"ETHFI":   true,
		"USDT":    true,
		"USDC":    true,
		"SUSDE":   true,
		"USDE":    true, // Ethena USDe
		"ENA":     true,
		"HYPE":    true, // Hyperliquid
		"BERA":    true,
		"SOL":     true,
		"SUI":     true,
		"APT":     true,
	}
}

// GetSkipList returns a map of market pairs that should be skipped during processing.
// These pairs are excluded from market operations for various reasons:
//   - milkTIA/TIA: This pair is skipped because milkTIA is not verified by CoinMarketCap.
//     The token's logo and social links are from third-party sources, requiring special handling
//     and verification before it can be properly integrated into the market map.
func GetSkipList() map[string]bool {
	return map[string]bool{
		"milkTIA/TIA": true, // Skip this pair as it requires special handling
	}
}
