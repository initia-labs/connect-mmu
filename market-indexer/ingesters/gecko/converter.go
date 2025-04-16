package gecko

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
