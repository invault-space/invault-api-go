package entity

import "encoding/json"

//
///
///
type TransactionsResult struct {
	CustodyOrderId    string `json:"custodyOrderId"`
	CoinCode          string `json:"coinCode"`
	CoinName          string `json:"coinName"`
	AssetCode         string `json:"assetCode"`
	ToAddress         string `json:"toAddress"`
	AddressSourceType string `json:"addressSourceType"`
	Memo              string `json:"memo"`
	OrderType         int    `json:"orderType"`
	OrderStatus       int    `json:"orderStatus"`
	OrderTimestamp    string `json:"orderTimestamp"`
	EndTime           int64  `json:"endTime"`
	Amount            string `json:"amount"`
	OrderAmount       string `json:"orderAmount"`
	FeeCoin           string `json:"feeCoin"`
	OrderFee          string `json:"orderFee"`
	TransactionHash   string `json:"transactionHash"`
	BlockHeight       string `json:"blockHeight"`
	BlockHash         string `json:"blockHash"`
	ChainCoin         string `json:"chainCoin"`
	TokenAddress      string `json:"tokenAddress"`
	TimeStamp         int64  `json:"timeStamp"`
	TransactionStatus int    `json:"transactionStatus"`
	TransactionIndex  string `json:"transactionIndex"`
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
	CustodyOrderId      string `json:"custodyOrderId"`
	CoinCode            string `json:"coinCode"`
	CoinName            string `json:"coinName"`
	AssetCode           string `json:"assetCode"`
	ToAddress           string `json:"toAddress"`
	AddressSourceType   string `json:"addressSourceType"`
	Memo                string `json:"memo"`
	OrderType           int    `json:"orderType"`
	OrderStatus         int    `json:"orderStatus"`
	OrderTimestamp      string `json:"orderTimestamp"`
	EndTime             int64  `json:"endTime"`
	Amount              string `json:"amount"`
	TransactionHash     string `json:"transactionHash"`
	BlockHeight         string `json:"blockHeight"`
	BlockHash           string `json:"blockHash"`
	ChainCoin           string `json:"chainCoin"`
	TokenAddress        string `json:"tokenAddress"`
	TimeStamp           int64  `json:"timeStamp"`
	TransactionStatus   int    `json:"transactionStatus"`
	TransactionIndex    string `json:"transactionIndex"`
	ConfirmingThreshold int    `json:"confirmingThreshold"`
	ConfirmedNum        int    `json:"confirmedNum"`
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
	CustodyOrderId    string `json:"custodyOrderId"`
	CoinCode          string `json:"coinCode"`
	CoinName          string `json:"coinName"`
	AssetCode         string `json:"assetCode"`
	ToAddress         string `json:"toAddress"`
	AddressSourceType string `json:"addressSourceType"`
	Memo              string `json:"memo"`
	OrderType         int    `json:"orderType"`
	OrderStatus       int    `json:"orderStatus"`
	OrderTimestamp    string `json:"orderTimestamp"`
	Amount            string `json:"amount"`
	TransactionHash   string `json:"transactionHash"`
	BlockHeight       string `json:"blockHeight"`
	BlockHash         string `json:"blockHash"`
	ChainCoin         string `json:"chainCoin"`
	TokenAddress      string `json:"tokenAddress"`
	TimeStamp         int64  `json:"timeStamp"`
	TransactionStatus int    `json:"transactionStatus"`
	TransactionIndex  string `json:"transactionIndex"`
	Confirmations     int64  `json:"confirmations"`
	ConfirmedNum      int64  `json:"confirmedNum"`
	EndTime           string `json:"endTime"`
	Decimals          int    `json:"decimals"`
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
