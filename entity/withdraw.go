package entity

//withdrawApply
type WithdrawApply struct {
	RequestOrderId string `json:"requestOrderId"`
	CoinCode       string `json:"coinCode"`
	TokenAddress   string `json:"tokenAddress,omitempty"`
	Address        string `json:"address"`
	Memo           string `json:"memo,omitempty"`
	Amount         string `json:"amount"`
	Timestamp      uint64 `json:"timestamp"`
}
type WithdrawApplyReq struct {
	OpenRequest
	Params WithdrawApply `json:"params,omitempty"`
}

func (a *WithdrawApplyReq) Init(id int) {
	a.init()
	a.Id = id
	a.Method = "withdrawApply"
}

func (a *WithdrawApplyReq) Vaildate() error {
	return nil
}

//queryByReqOrderId
type QueryByReqOrderId struct {
	RequestOrderId string `json:"requestOrderId"`
}

type QueryByReqOrderIdReq struct {
	OpenRequest
	Params QueryByReqOrderId `json:"params,omitempty"`
}

func (a *QueryByReqOrderIdReq) Init(id int) {
	a.init()
	a.Id = id
	a.Method = "queryByReqOrderId"
}

func (a *QueryByReqOrderIdReq) Vaildate() error {
	return nil
}
