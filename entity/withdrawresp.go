package entity

import "encoding/json"

//
///
///
type WithdrawApplyResult struct {
	Success        bool   `json:"success"`
	Message        string `json:"message"`
	RequestOrderId string `json:"requestOrderId"`
	CoinCode       string `json:"coinCode"`
	Address        string `json:"address"`
	Memo           string `json:"memo"`
	Amount         string `json:"amount"`
	Timestamp      int64  `json:"timestamp"`
}
type WithdrawApplyResp struct {
	OpenResult
	Reuslt WithdrawApplyResult `json:"result,omitempty"`
}

func (a *WithdrawApplyResp) ParseResult(j []byte) (*WithdrawApplyResp, error) {
	err := json.Unmarshal(j, a)
	return a, err
}

//
///
///
type QueryByReqOrderIdResult struct {
	CustodyOrderId    string `json:"custodyOrderId"`
	RequestOrderId    string `json:"requestOrderId"`
	OrderType         int    `json:"orderType"`
	OrderStatus       int    `json:"orderStatus"`
	Description       string `json:"description"`
	OrderTimeStamp    string `json:"orderTimeStamp"`
	CoinCode          string `json:"coinCode"`
	AssetCode         string `json:"assetCode"`
	FromAddress       string `json:"fromAddress"`
	ToAddress         string `json:"toAddress"`
	Amount            string `json:"amount"`
	Memo              string `json:"memo"`
	TransactionFee    string `json:"transactionFee"`
	FeeCoin           string `json:"feeCoin"`
	OrderFee          string `json:"orderFee"`
	TransactionHash   string `json:"transactionHash"`
	TransactionStatus int    `json:"transactionStatus"`
	TransactionType   string `json:"transactionType"`
	BlockHeight       int64  `json:"blockHeight"`
	BlockHash         string `json:"blockHash"`
	ChainCoin         string `json:"chainCoin"`
	TokenAddress      string `json:"tokenAddress"`
	TransactionIndex  string `json:"transactionIndex"`
}
type QueryByReqOrderIdResp struct {
	OpenResult
	Reuslt QueryByReqOrderIdResult `json:"result,omitempty"`
}

func (a *QueryByReqOrderIdResp) ParseResult(j []byte) (*QueryByReqOrderIdResp, error) {
	err := json.Unmarshal(j, a)
	return a, err
}
