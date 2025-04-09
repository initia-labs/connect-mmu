package filter

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

func GetSkipList() map[string]bool {
	return map[string]bool{
		"milkTIA/TIA": true,
	}
}
