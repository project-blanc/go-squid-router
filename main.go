package main

import (
	"fmt"

	v2 "github.com/project-blanc/go-squid-router/v2"
)

func main() {
	cli := v2.NewClient("123")
	_, err := cli.Route(v2.RouteRequestParameters{
		FromChain:   "42161",
		ToChain:     "42220",
		FromToken:   "0xfd086bc7cd5c481dcc9c85ebe478a1c0b69fcbb9",
		ToToken:     "0x765de816845861e75a25fca122bb6898b8b1282a",
		ToAddress:   "0x4254EbB758d60e3f15A93D53BDa6a760F0Ae24f9",
		FromAmount:  "10000000",
		FromAddress: "0xcc70cbe5e7669E5c72E6C1c7b73a0da0855050AB",
		SlippageConfig: v2.SlippageConfig{
			Slippage: 1,
			AutoMode: v2.AGGRESSIVE,
		},
	})

	if err != nil {
		fmt.Println(err)
	}
}
