package initia

//{
//	"pools": [
//	{
//		"lp": "move/543b35a39cfadad3da3c23249c474455d15efd2f94f849473226dee8a3c7a9e1",
//		"lp_metadata": "0x543b35a39cfadad3da3c23249c474455d15efd2f94f849473226dee8a3c7a9e1",
//		"is_minitswap_pool": false,
//		"pool_type": "BALANCER",
//		"coins": [
//		{
//			"denom": "ibc/6490A7EAB61059BFC1CDDEB05917DD70BDF3A611654162A1A47DB930D40D8AF4",
//			"metadata": "0xe0e9394b24e53775d6af87934ac02d73536ad58b7894f6ccff3f5e7c0d548e55",
//			"weight": "0.2"
//		},
//		{
//			"denom": "uinit",
//			"metadata": "0x8e4733bdabcf7d4afc3d14f0dd46c9bf52fb0fce9e4b996c939e195b8bc891d9",
//			"weight": "0.8"
//		}],
//		"swap_fee_apr": 0.3717892141152,
//		"staking_apr": 4.29469906132925,
//		"total_apr": 4.66648827544445,
//		"volume_24h": 2640289890742.88,
//		"liquidity": 8246860048845.37,
//		"value_per_lp": 0.852691576175358,
//		"swap_fee_rate": 0.003
//	}, ... ],
//	"pagination": {
//	"next_key": null,
//	"total": "9"
//	}
//}

type InitiaDexResponse struct {
	Pools      []PoolsResponse    `json:"pools"`
	Pagination PaginationResponse `json:"pagination"`
}

type PoolsResponse struct {
	Lp              string          `json:"lp"`
	LpMetadata      string          `json:"lp_metadata"`
	IsMinitswapPool bool            `json:"is_minitswap_pool"`
	PoolType        string          `json:"pool_type"`
	Coins           []CoinsResponse `json:"coins"`
	SwapFeeApr      float64         `json:"swap_fee_apr"`
	StakingApr      float64         `json:"staking_apr"`
	TotalApr        float64         `json:"total_apr"`
	Volume24H       float64         `json:"volume_24h"`
	Liquidity       float64         `json:"liquidity"`
	ValuePerLp      float64         `json:"value_per_lp"`
	SwapFeeRate     float64         `json:"swap_fee_rate"`
}

type CoinsResponse struct {
	Denom    string `json:"denom"`
	Metadata string `json:"metadata"`
	Weight   string `json:"weight"`
}

type PaginationResponse struct {
	NextKey interface{} `json:"next_key"`
	Total   string      `json:"total"`
}

type InitiaRegistryResponse struct {
	Schema    string `json:"$schema"`
	ChainName string `json:"chain_name"`
	Assets    []struct {
		Description string `json:"description"`
		DenomUnits  []struct {
			Denom    string `json:"denom"`
			Exponent int    `json:"exponent"`
		} `json:"denom_units"`
		Base        string `json:"base"`
		Display     string `json:"display"`
		Name        string `json:"name"`
		Symbol      string `json:"symbol"`
		CoingeckoId string `json:"coingecko_id"`
		Images      []struct {
			Png string `json:"png"`
		} `json:"images"`
		LogoURIs struct {
			Png string `json:"png"`
		} `json:"logo_URIs"`
		Traces []struct {
			Type         string `json:"type"`
			Counterparty struct {
				ChainName string `json:"chain_name"`
				BaseDenom string `json:"base_denom"`
				ChannelId string `json:"channel_id"`
			} `json:"counterparty"`
			Chain struct {
				ChannelId string `json:"channel_id"`
				Path      string `json:"path"`
			} `json:"chain"`
		} `json:"traces,omitempty"`
	} `json:"assets"`
}

func (r *InitiaRegistryResponse) toMap() map[string]string {
	m := make(map[string]string, len(r.Assets))
	for _, asset := range r.Assets {
		m[asset.Base] = asset.Symbol
	}
	return m
}
