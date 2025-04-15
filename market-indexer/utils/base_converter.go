package utils

func ConvertTargetBase(base string) string {
	mapping := map[string]string{
		"USDE": "USDe",
	}

	if newBase, exists := mapping[base]; exists {
		return newBase
	}
	return base
}
