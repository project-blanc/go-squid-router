package v1

import (
	"fmt"
	"net/url"
	"strconv"
)

type CallType int

const (
	CallTypeDefault CallType = iota
	CallTypeFullTokenBalance
	CallTypeFullNativeBalance
	CollectTokenBalance
)

type CustomContractCall struct {
	CallType     CallType
	Target       string
	Value        string
	CallData     string
	EstimatedGas string
	Payload      struct {
		TokenAddress  string
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
	Prefer []string
	// ReceiveGasOnDestination receive gas on destination chain
	ReceiveGasOnDestination bool
}

func (p RouteRequestParameters) ToQuery() url.Values {
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
		query.Add(fmt.Sprintf("prefer[%d]", i), prefer)
	}

	query.Add("receiveGasOnDestination", strconv.FormatBool(p.ReceiveGasOnDestination))

	return query
}

type TransactionRequest struct {
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
