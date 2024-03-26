package v2

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type DexName string

const (
	AGNI_V3              DexName = "Agni-v3"
	AERODROME            DexName = "Aerodrome"
	APESWAP              DexName = "Apeswap"
	APESWAP_V3           DexName = "Apeswap-v3"
	BALANCER_V2          DexName = "Balancer-v2"
	BASESWAP             DexName = "Baseswap"
	BASESWAP_V3          DexName = "Baseswap-v3"
	BEAMSWAP             DexName = "Beamswap"
	BEAMSWAP_V3          DexName = "Beamswap-v3"
	CAMELOT              DexName = "Camelot"
	CURVE_V2             DexName = "Curve-v2"
	ELLIPSIS             DexName = "Ellipsis"
	EQUILIBRE            DexName = "Equilibre"
	EQUALIZER            DexName = "Equalizer"
	FUSIONX_V2           DexName = "FusionX-v2"
	FUSIONX_V3           DexName = "FusionX-v3"
	GMX                  DexName = "GMX"
	HORIZON              DexName = "Horizon"
	HORIZON_V3           DexName = "Horizon-v3"
	KYBERSWAP            DexName = "KyberSwap"
	KYBERSWAP_AGGREGATOR DexName = "kyberswap-aggregator"
	KINETIX_V3           DexName = "Kinetix-v3"
	MENTO_V2             DexName = "Mento-v2"
	OPENOCEAN            DexName = "OpenOcean"
	OSMOSIS              DexName = "Osmosis"
	PANCAKESWAP          DexName = "Pancakeswap"
	PANCAKESWAP_V3       DexName = "Pancakeswap-v3"
	PANCAKESWAP_STABLE   DexName = "Pancakeswap-stable"
	PANGOLIN             DexName = "Pangolin"
	PLATYPUS             DexName = "Platypus"
	QUICKSWAP            DexName = "Quickswap"
	QUICKSWAP_V3         DexName = "Quickswap-v3"
	RAMSES               DexName = "Ramses"
	SPOOKYSWAP           DexName = "Spookyswap"
	STELLASWAP           DexName = "Stellaswap"
	STELLASWAP_V3        DexName = "Stellaswap-v3"
	STELLASWAP_SADDLE    DexName = "Stellaswap-saddle"
	SUSHISWAP            DexName = "Sushiswap"
	SUSHISWAP_V3         DexName = "Sushiswap-v3"
	SWAPBASED            DexName = "SwapBased"
	SYNTHSWAP_V2         DexName = "SynthSwap-v2"
	SYNTHSWAP_V3         DexName = "SynthSwap-v3"
	SKYDROME             DexName = "Skydrome"
	THENA                DexName = "Thena"
	THENA_V3             DexName = "Thena-v3"
	TRADERJOE            DexName = "TraderJoe"
	TRIDENT              DexName = "Trident"
	UBESWAP              DexName = "Ubeswap"
	UBESWAP_V3           DexName = "Ubeswap-v3"
	UNISWAP_V2           DexName = "Uniswap-v2"
	UNISWAP_V3           DexName = "Uniswap-v3"
	WIGOSWAP             DexName = "WigoSwap"
	WOMBAT               DexName = "Wombat"
	VELODROME            DexName = "Velodrome"
	VELODROME_V2         DexName = "Velodrome-v2"
	VELOCIMETER          DexName = "Velocimeter"
	ZYBERSWAP            DexName = "Zyberswap"
	ZYBERSWAP_V3         DexName = "Zyberswap-v3"
)

type SlippageMode int

const (
	AGGRESSIVE SlippageMode = iota
	NORMAL
	FLEXIBLE
)

type SlippageConfig struct {
	Slippage float64      `json:"slippage"`
	AutoMode SlippageMode `json:"autoMode"`
}

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
	Prefer []DexName `json:"prefer,omitempty"`
	// ReceiveGasOnDestination receive gas on destination chain
	ReceiveGasOnDestination bool           `json:"receiveGasOnDestination"`
	SlippageConfig          SlippageConfig `json:"slippageConfig"`
}

type SquidRouteType string

const (
	SquidRouteTypeCallBridgeCall SquidRouteType = "CALL_BRIDGE_CALL"
	SquidRouteTypeCallBridge     SquidRouteType = "CALL_BRIDGE"
	SquidRouteTypeBridgeCall     SquidRouteType = "BRIDGE_CALL"
	SquidRouteTypeBridge         SquidRouteType = "BRIDGE"
	SquidRouteTypeEVMOnly        SquidRouteType = "EVM_ONLY"
	SquidRouteTypeCosmosOnly     SquidRouteType = "COSMOS_ONLY"
)

type SquidData struct {
	Target               common.Address `json:"target"`
	Data                 hexutil.Bytes  `json:"data"`
	Value                string         `json:"value"`
	GasLimit             string         `json:"gasLimit"`
	GasPrice             string         `json:"gasPrice"`
	MaxFeePerGas         string         `json:"maxFeePerGas"`
	MaxPriorityFeePerGas string         `json:"maxPriorityFeePerGas"`
}

type RouteResponse struct {
	// @TODO estimate
	// @TodO params
	TransactionRequest SquidData `json:"transactionRequest"`
}
