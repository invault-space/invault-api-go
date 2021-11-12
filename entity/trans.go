package entity

type Transactions struct {
	CoinCode  string `json:"coinCode,omitempty"`
	OrderType string `json:"orderType"`
	Address   string `json:"address,omitempty"`
	BeginTime uint64 `json:"beginTime,omitempty"`
	EndTime   uint64 `json:"endTime,omitempty"`
}

func (a Transactions) Method() string {
	return "transactions"
}
func (a Transactions) Vaildate() error {
	return nil
}

//
type TransactionById struct {
	OrderType      string `json:"orderType"`
	CustodyOrderId string `json:"custodyOrderId"`
}

func (a TransactionById) Method() string {
	return "transactionById"
}
func (a TransactionById) Vaildate() error {
	return nil
}

//
type PendingTransactions struct {
	CoinCode  string `json:"coinCode,omitempty"`
	OrderType string `json:"orderType"`
	Address   string `json:"address,omitempty"`
	BeginTime uint64 `json:"beginTime,omitempty"`
	EndTime   uint64 `json:"endTime,omitempty"`
}

func (a PendingTransactions) Method() string {
	return "pendingTransactions"
}
func (a PendingTransactions) Vaildate() error {
	return nil
}

//
type PendingTransactionById struct {
	OrderType      string `json:"orderType"`
	CustodyOrderId string `json:"custodyOrderId"`
}

func (a PendingTransactionById) Method() string {
	return "pendingTransactionById"
}
func (a PendingTransactionById) Vaildate() error {
	return nil
}

//
type TransactionByTxHash struct {
	CoinCode string `json:"coinCode"`
	TxHash   string `json:"txHash"`
}

func (a TransactionByTxHash) Method() string {
	return "transactionByTxHash"
}
func (a TransactionByTxHash) Vaildate() error {
	return nil
}

//
type BlockHeight struct {
}

func (a BlockHeight) Method() string {
	return "blockHeight"
}
func (a BlockHeight) Vaildate() error {
	return nil
}

//
type TransactionByBlockHeight struct {
	CoinCode    string `json:"coinCode"`
	BlockHeight uint64 `json:"blockHeight"`
}

func (a TransactionByBlockHeight) Method() string {
	return "transactionByBlockHeight"
}
func (a TransactionByBlockHeight) Vaildate() error {
	return nil
}
