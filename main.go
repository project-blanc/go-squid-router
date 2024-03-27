package main

import (
	"fmt"

	squid "github.com/project-blanc/go-squid-router/v2"
)

func main() {
	s := squid.NewClient("")

	r, err := s.Route(squid.RouteRequestParameters{
		FromChain:   "42161",
		ToChain:     "42220",
		FromToken:   "0xfd086bc7cd5c481dcc9c85ebe478a1c0b69fcbb9",
		ToToken:     "0x765de816845861e75a25fca122bb6898b8b1282a",
		ToAddress:   "0x4254EbB758d60e3f15A93D53BDa6a760F0Ae24f9",
		FromAmount:  "10000000",
		FromAddress: "0xcc70cbe5e7669E5c72E6C1c7b73a0da0855050AB",
		Prefer: []squid.DexName{
			squid.PANCAKESWAP,
			squid.UNISWAP_V2,
			squid.CURVE_V2,
		},
		SlippageConfig: squid.SlippageConfig{
			Slippage: 1,
			AutoMode: squid.NORMAL,
		},
	})

	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", r)
}
