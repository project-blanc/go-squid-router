package v2

import v1 "github.com/project-blanc/go-squid-router/v1"

type SlippageMode int

const (
	AGGRESSIVE SlippageMode = iota
	NORMAL
	FLEXIBLE
)

type RouteRequestParameters struct {
	// FromChain is the from chain id
	FromChain string `json:"fromChain"`
	// ToChain is the to chain id
	ToChain string `json:"toChain"`
	// FromToken is the from token address
	FromToken string `json:"fromToken"`
	// ToToken is the to token address
	ToToken string `json:"toToken"`
	// FromAmount is the amount to be sent from To Chain
	FromAmount string `json:"fromAmount"`
	// FromAddress EVM address for EVM to Cosmos, and Cosmos address for Cosmos to EVM.
	FromAddress string `json:"fromAddress"`
	// ToAddress is the ToChain recipient address
	ToAddress string `json:"toAddress"`
	// QuoteOnly returns only the route quote and excludes all call data
	QuoteOnly   bool `json:"quoteOnly"`
	EnableBoost bool `json:"enableBoost"`
	// Prefer array of supported DEXs for this trade
	Prefer []v1.DexName `json:"prefer"`
	// ReceiveGasOnDestination receive gas on destination chain
	ReceiveGasOnDestination bool `json:"receiveGasOnDestination"`

	SlippageConfig struct {
		Slippage float64      `json:"slippage"`
		AutoMode SlippageMode `json:"autoMode"`
	} `json:"slippageConfig"`
}
