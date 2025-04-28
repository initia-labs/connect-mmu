package utils

func ConvertTargetBase(base string) string {
	mapping := map[string]string{
		"USDE":  "USDe",
		"WEETH": "weETH",
		"SUSDE": "sUSDe",
	}

	if newBase, exists := mapping[base]; exists {
		return newBase
	}
	return base
}
