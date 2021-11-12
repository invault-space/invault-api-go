package entity

type IOpenEntity interface {
	Init(int)
	Vaildate() error
}

type OpenRequest struct {
	Id      int         `json:"id"`
	JsonRpc string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params,omitempty"`
}

func (a *OpenRequest) init() {
	a.JsonRpc = "2.0"
}

// func NewRequest(a *OpenRequest) IOpenEntity {
// 	return a
// }
// func (a *OpenRequest) Init() {
// 	a.Id = 0
// 	a.JsonRpc = "2.0"
// }

// func (a *OpenRequest) Vaildate() error {
// 	v := a.Params.(IOpenEntity)
// 	return v.Vaildate()
// }

// func (a *OpenRequest) String() string {
// 	b, _ := json.Marshal(a)
// 	return string(b)
// }

///
///
///
type errResult struct {
	code    int    `json:"Code"`
	Message string `json:"message"`
}
type OpenResult struct {
	Id      int       `json:"id"`
	JsonRpc string    `json:"jsonrpc"`
	Method  string    `json:"method"`
	Error   errResult `json:"error"`
	// Reuslt interface{} `json:"reuslt,omitempty"`
}
