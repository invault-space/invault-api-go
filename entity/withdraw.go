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

func (a WithdrawApply) Method() string {
	return "withdrawApply"
}
func (a WithdrawApply) Vaildate() error {
	return nil
}

//queryByReqOrderId
type QueryByReqOrderId struct {
	RequestOrderId string `json:"requestOrderId"`
}

func (a QueryByReqOrderId) Method() string {
	return "queryByReqOrderId"
}
func (a QueryByReqOrderId) Vaildate() error {
	return nil
}
