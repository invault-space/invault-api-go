package entity

import "encoding/json"

//
///
///
type TransactionsResult struct {
	CustodyOrderId    string      `json:"custodyOrderId"`
	CoinCode          string      `json:"coinCode"`
	CoinName          string      `json:"coinName"`
	AssetCode         string      `json:"assetCode"`
	ToAddress         string      `json:"toAddress"`
	AddressSourceType string      `json:"addressSourceType"`
	Memo              string      `json:"memo"`
	OrderType         json.Number `json:"orderType"`
	OrderStatus       json.Number `json:"orderStatus"`
	OrderTimestamp    string      `json:"orderTimestamp"`
	EndTime           json.Number `json:"endTime"`
	Amount            string      `json:"amount"`
	OrderAmount       string      `json:"orderAmount"`
	FeeCoin           string      `json:"feeCoin"`
	OrderFee          string      `json:"orderFee"`
	TransactionHash   string      `json:"transactionHash"`
	BlockHeight       string      `json:"blockHeight"`
	BlockHash         string      `json:"blockHash"`
	ChainCoin         string      `json:"chainCoin"`
	TokenAddress      string      `json:"tokenAddress"`
	TimeStamp         json.Number `json:"timeStamp"`
	TransactionStatus json.Number `json:"transactionStatus"`
	TransactionIndex  string      `json:"transactionIndex"`
}
type TransactionsResp struct {
	OpenResult
	Reuslt []TransactionsResult `json:"result,omitempty"`
}

func (a *TransactionsResp) ParseResult(j []byte) (*TransactionsResp, error) {
	err := json.Unmarshal(j, a)
	return a, err
}

//
///
///

type TransactionByIdResp struct {
	OpenResult
	Reuslt TransactionsResult `json:"result,omitempty"`
}

func (a *TransactionByIdResp) ParseResult(j []byte) (*TransactionByIdResp, error) {

	err := json.Unmarshal(j, a)
	return a, err
}

//
///
///
type PendingTransactionsReuslt struct {
	CustodyOrderId      string      `json:"custodyOrderId"`
	CoinCode            string      `json:"coinCode"`
	CoinName            string      `json:"coinName"`
	AssetCode           string      `json:"assetCode"`
	ToAddress           string      `json:"toAddress"`
	AddressSourceType   string      `json:"addressSourceType"`
	Memo                string      `json:"memo"`
	OrderType           json.Number `json:"orderType"`
	OrderStatus         json.Number `json:"orderStatus"`
	OrderTimestamp      string      `json:"orderTimestamp"`
	EndTime             json.Number `json:"endTime"`
	Amount              string      `json:"amount"`
	TransactionHash     string      `json:"transactionHash"`
	BlockHeight         string      `json:"blockHeight"`
	BlockHash           string      `json:"blockHash"`
	ChainCoin           string      `json:"chainCoin"`
	TokenAddress        string      `json:"tokenAddress"`
	TimeStamp           json.Number `json:"timeStamp"`
	TransactionStatus   json.Number `json:"transactionStatus"`
	TransactionIndex    string      `json:"transactionIndex"`
	ConfirmingThreshold json.Number `json:"confirmingThreshold"`
	ConfirmedNum        json.Number `json:"confirmedNum"`
}
type PendingTransactionsResp struct {
	OpenResult
	Reuslt []PendingTransactionsReuslt `json:"result,omitempty"`
}

func (a *PendingTransactionsResp) ParseResult(j []byte) (*PendingTransactionsResp, error) {

	err := json.Unmarshal(j, a)
	return a, err
}

//
///
///
type PendingTransactionByIdResult struct {
	CustodyOrderId    string      `json:"custodyOrderId"`
	CoinCode          string      `json:"coinCode"`
	CoinName          string      `json:"coinName"`
	AssetCode         string      `json:"assetCode"`
	ToAddress         string      `json:"toAddress"`
	AddressSourceType string      `json:"addressSourceType"`
	Memo              string      `json:"memo"`
	OrderType         json.Number `json:"orderType"`
	OrderStatus       json.Number `json:"orderStatus"`
	OrderTimestamp    string      `json:"orderTimestamp"`
	Amount            string      `json:"amount"`
	TransactionHash   string      `json:"transactionHash"`
	BlockHeight       string      `json:"blockHeight"`
	BlockHash         string      `json:"blockHash"`
	ChainCoin         string      `json:"chainCoin"`
	TokenAddress      string      `json:"tokenAddress"`
	TimeStamp         json.Number `json:"timeStamp"`
	TransactionStatus json.Number `json:"transactionStatus"`
	TransactionIndex  string      `json:"transactionIndex"`
	Confirmations     json.Number `json:"confirmations"`
	ConfirmedNum      json.Number `json:"confirmedNum"`
	EndTime           string      `json:"endTime"`
	Decimals          json.Number `json:"decimals"`
}
type PendingTransactionByIdResp struct {
	OpenResult
	Reuslt PendingTransactionByIdResult `json:"result,omitempty"`
}

func (a *PendingTransactionByIdResp) ParseResult(j []byte) (*PendingTransactionByIdResp, error) {

	err := json.Unmarshal(j, a)
	return a, err
}

//
///
///
type TransactionByTxHashResp struct {
	OpenResult
	Reuslt TransactionsResult `json:"result,omitempty"`
}

func (a *TransactionByTxHashResp) ParseResult(j []byte) (*TransactionByTxHashResp, error) {

	err := json.Unmarshal(j, a)
	return a, err
}

//
///
///
type BlockHeightResult struct {
	CoinCode    string `json:"coinCode"`
	BlockHeight string `json:"blockHeight"`
	ChainCoin   string `json:"chainCoin"`
}
type BlockHeightResp struct {
	OpenResult
	Reuslt []BlockHeightResult `json:"result,omitempty"`
}

func (a *BlockHeightResp) ParseResult(j []byte) (*BlockHeightResp, error) {

	err := json.Unmarshal(j, a)
	return a, err
}

//
///
///
type TransactionByBlockHeightResp struct {
	OpenResult
	Reuslt TransactionsResult `json:"result,omitempty"`
}

func (a *TransactionByBlockHeightResp) ParseResult(j []byte) (*TransactionByBlockHeightResp, error) {

	err := json.Unmarshal(j, a)
	return a, err
}
