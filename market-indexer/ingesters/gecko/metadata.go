package gecko

import (
	"github.com/skip-mev/connect/v2/providers/apis/defi/curve"
	"github.com/skip-mev/connect/v2/providers/apis/defi/uniswapv3"
)

func MakeMetadata(dex string) interface{} {
	switch dex {
	case GeckoVenueUniswapEth:
		return uniswapv3.PoolConfig{}
	case GeckoVenueCurve:
		return curve.CurveMetadata{}
	default:
		return nil
	}
}

func MakeNetwork(network string) string {
	switch network {
	case "eth":
		return "ethereum"
	default:
		return ""
	}
}
