package service

func If(condition bool, falseVal string, trueVal ...string) string {
	if condition {
		return falseVal
	}
	return trueVal[0]
}

///
///
///
type OpenReuslt struct {
	Id      int         `json:"id"`
	JsonRpc string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Error   interface{} `json:"error,omitempty"`
}
