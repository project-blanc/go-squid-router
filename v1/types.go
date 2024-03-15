package v1

import (
	"fmt"
	"net/url"
	"strconv"
)

type DexName string

const (
	AGNI_V3               DexName = "Agni_v3"
	AERODROME             DexName = "Aerodrome"
	UNISWAP_V2            DexName = "UniswapV2"
	UNISWAP_V3            DexName = "UniswapV3"
	UBESWAP               DexName = "Ubeswap"
	SPOOKYSWAP            DexName = "Spookyswap"
	SUSHISWAP             DexName = "Sushiswap"
	SUSHISWAP_V3          DexName = "SushiswapV3"
	STELLASWAP            DexName = "Stellaswap"
	STELLASWAP_V3         DexName = "StellaswapV3"
	PANGOLIN              DexName = "Pangolin"
	CURVE_V2              DexName = "Curve_v2"
	CURVE_V2_POOL         DexName = "Curve_v2_Pool"
	ELLIPSIS              DexName = "Ellipsis"
	FUSIONX_V2            DexName = "FusionX_v2"
	FUSIONX_V3            DexName = "FusionX_v3"
	KINETIX_V3            DexName = "Kinetix_v3"
	QUICKSWAP             DexName = "Quickswap"
	STELLASWAP_SADDLE     DexName = "Stellaswap_Saddle"
	VELODROME_V2          DexName = "Velodrome_V2"
	EQUALIZER             DexName = "Equalizer"
	EQUILIBRE             DexName = "Equilibre"
	QUICKSWAP_V3          DexName = "Quickswap_v3"
	PANCAKESWAP           DexName = "Pancakeswap"
	PANCAKESWAP_V3        DexName = "Pancakeswap_v3"
	PANCAKESWAP_STABLE    DexName = "Pancakeswap_stable"
	TRADERJOE             DexName = "TraderJoe"
	TRIDENT               DexName = "Trident"
	PLATYPUS              DexName = "Platypus"
	WOMBAT                DexName = "Wombat"
	ZYBERSWAP             DexName = "Zyberswap"
	KYBERSWAP             DexName = "KyberSwap"
	KYBERSWAP_AGGREGATOR  DexName = "KyberSwap_Aggregator"
	GMX                   DexName = "GMX"
	APESWAP               DexName = "Apeswap"
	OPENOCEAN             DexName = "OpenOcean"
	OSMOSIS               DexName = "Osmosis"
	THENA_V3              DexName = "Thena_v3"
	THENA                 DexName = "Thena"
	UBESWAP_V3            DexName = "Ubeswap_v3"
	ZYBERSWAP_V3          DexName = "Zyberswap_v3"
	BEAMSWAP              DexName = "Beamswap"
	BEAMSWAP_V2_SADDLE    DexName = "Beamswap_v2_Saddle"
	SWAPBASED             DexName = "SwapBased"
	HORIZON_V3            DexName = "Horizon_v3"
	BASESWAP              DexName = "Baseswap"
	SYNTHSWAP_V2          DexName = "SynthSwap_v2"
	SYNTHSWAP_V3          DexName = "SynthSwap_v3"
	SKYDROME              DexName = "Skydrome"
	VELOCIMETER           DexName = "Velocimeter"
	CAMELOT               DexName = "Camelot"
	THRUSTER_V3           DexName = "Thruster_v3"
	THRUSTER_V2_1_PERCENT DexName = "Thruster_v2_1_percent"
	THRUSTER_V2_30_BPS    DexName = "Thruster_v2_30_bps"
	RAMSES_V3             DexName = "ramses_v3"
)

type CallType int

const (
	CallTypeDefault CallType = iota
	CallTypeFullTokenBalance
	CallTypeFullNativeBalance
	CollectTokenBalance
)

type CustomContractCall struct {
	// CallType squid call type
	CallType CallType
	// Target is the address of the contract to be called
	Target string
	// Value is the amount of native coin, in most scenarios should be "0"
	Value string
	// CallData is the contract call encoded call data
	CallData string
	// EstimatedGas is the amount of gas of the call
	EstimatedGas string
	Payload      struct {
		// TokenAddress is the address of the ERC20 token
		TokenAddress string
		// InputPosition is the position of the amount argument in the contract call to set the balance dynamically
		InputPosition string
	}
}

type RouteRequestParameters struct {
	// FromChain is the from chain id
	FromChain string
	// ToChain is the to chain id
	ToChain string
	// FromToken is the from token address
	FromToken string
	// ToToken is the to token address
	ToToken string
	// FromAmount is the amount to be sent from To Chain
	FromAmount string
	// FromAddress EVM address for EVM to Cosmos, and Cosmos address for Cosmos to EVM.
	FromAddress string
	// ToAddress is the ToChain recipient address
	ToAddress string
	// Slippage must be between 0 and 99.99
	Slippage float64
	// QuoteOnly returns only the route quote and excludes all call data
	QuoteOnly bool
	// EnableExpress enables express feature
	EnableExpress bool
	// CustomContractCalls array of custom contract calls
	CustomContractCalls []CustomContractCall
	// Prefer array of supported DEXs for this trade
	Prefer []DexName
	// ReceiveGasOnDestination receive gas on destination chain
	ReceiveGasOnDestination bool
}

func (p RouteRequestParameters) Query() url.Values {
	query := url.Values{}
	query.Add("fromChain", p.FromChain)
	query.Add("toChain", p.ToChain)
	query.Add("fromToken", p.FromToken)
	query.Add("toToken", p.ToToken)
	query.Add("fromAmount", p.FromAmount)

	if p.FromAddress != "" {
		query.Add("fromAddress", p.FromAddress)
	}

	query.Add("toAddress", p.ToAddress)

	query.Add("slippage", strconv.FormatFloat(p.Slippage, 'f', -1, 64))
	query.Add("quoteOnly", strconv.FormatBool(p.QuoteOnly))
	query.Add("enableExpress", strconv.FormatBool(p.EnableExpress))

	for i, call := range p.CustomContractCalls {
		query.Add(fmt.Sprintf("customContractCalls[%d][callType]", i), strconv.Itoa(int(call.CallType)))
		query.Add(fmt.Sprintf("customContractCalls[%d][target]", i), call.Target)
		query.Add(fmt.Sprintf("customContractCalls[%d][value]", i), call.Value)
		query.Add(fmt.Sprintf("customContractCalls[%d][callData]", i), call.CallData)
		query.Add(fmt.Sprintf("customContractCalls[%d][estimatedGas]", i), call.EstimatedGas)
		query.Add(fmt.Sprintf("customContractCalls[%d][payload][tokenAddress]", i), call.Payload.TokenAddress)
		query.Add(fmt.Sprintf("customContractCalls[%d][payload][inputPos]", i), call.Payload.InputPosition)
	}

	for i, prefer := range p.Prefer {
		query.Add(fmt.Sprintf("prefer[%d]", i), string(prefer))
	}

	query.Add("receiveGasOnDestination", strconv.FormatBool(p.ReceiveGasOnDestination))

	return query
}

type TransactionRequest struct {
	RouteType            string `json:"routeType"`
	TargetAddress        string `json:"targetAddress"`
	Data                 string `json:"data"`
	Value                string `json:"value"`
	GasLimit             string `json:"gasLimit"`
	GasPrice             string `json:"gasPrice"`
	MaxFeePerGas         string `json:"maxFeePerGas"`
	MaxPriorityFeePerGas string `json:"maxPriorityFeePerGas"`
}

type RouteResponse struct {
	// @TODO estimate
	// @TodO params
	TransactionRequest TransactionRequest `json:"transactionRequest"`
}

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorType string `json:"errorType"`
}

func (e ErrorResponse) Err() error {
	return fmt.Errorf("squid err: %s message: %s", e.ErrorType, e.Message)
}
