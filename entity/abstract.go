package entity

import "encoding/json"

type IOpenEntity interface {
	Method() string
	Vaildate() error
}

type Abstract struct {
	Id      int         `json:"id"`
	JsonRpc string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params,omitempty"`
}

func NewRequest(params IOpenEntity) *Abstract {
	a := Abstract{}
	a.Init()

	a.Method = params.Method()
	a.Params = params
	return &a
}
func (a *Abstract) Init() {
	a.Id = 0
	a.JsonRpc = "2.0"
	a.Params = struct{}{}
}

func (a *Abstract) Vaildate() error {
	v := a.Params.(IOpenEntity)
	return v.Vaildate()
}

func (a *Abstract) String() string {
	b, _ := json.Marshal(a)
	return string(b)
}
