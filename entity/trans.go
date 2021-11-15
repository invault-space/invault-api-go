package entity

type Transactions struct {
	CoinCode  string `json:"coinCode,omitempty"`
	OrderType string `json:"orderType"`
	Address   string `json:"address,omitempty"`
	BeginTime uint64 `json:"beginTime,omitempty"`
	EndTime   uint64 `json:"endTime,omitempty"`
}

type TransactionsReq struct {
	OpenRequest
	Params Transactions `json:"params,omitempty"`
}

func (a *TransactionsReq) Init(id int) {
	a.init()
	a.Id = id
	a.Method = "transactions"
}

func (a *TransactionsReq) Vaildate() error {
	return nil
}

//
type TransactionById struct {
	OrderType      string `json:"orderType"`
	CustodyOrderId string `json:"custodyOrderId"`
}

type TransactionByIdReq struct {
	OpenRequest
	Params TransactionById `json:"params,omitempty"`
}

func (a *TransactionByIdReq) Init(id int) {
	a.init()
	a.Id = id
	a.Method = "transactionById"
}

func (a *TransactionByIdReq) Vaildate() error {
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

type PendingTransactionsReq struct {
	OpenRequest
	Params PendingTransactions `json:"params,omitempty"`
}

func (a *PendingTransactionsReq) Init(id int) {
	a.init()
	a.Id = id
	a.Method = "pendingTransactions"
}

func (a *PendingTransactionsReq) Vaildate() error {
	return nil
}

//
type PendingTransactionById struct {
	OrderType      string `json:"orderType"`
	CustodyOrderId string `json:"custodyOrderId"`
}

type PendingTransactionByIdReq struct {
	OpenRequest
	Params PendingTransactionById `json:"params,omitempty"`
}

func (a *PendingTransactionByIdReq) Init(id int) {
	a.init()
	a.Id = id
	a.Method = "pendingTransactionById"
}

func (a *PendingTransactionByIdReq) Vaildate() error {
	return nil
}

//
type TransactionByTxHash struct {
	CoinCode string `json:"coinCode"`
	TxHash   string `json:"txHash"`
}

type TransactionByTxHashReq struct {
	OpenRequest
	Params TransactionByTxHash `json:"params,omitempty"`
}

func (a *TransactionByTxHashReq) Init(id int) {
	a.init()
	a.Id = id
	a.Method = "transactionByTxHash"
}

func (a *TransactionByTxHashReq) Vaildate() error {
	return nil
}

//
type BlockHeightReq struct {
	OpenRequest
	Params interface{} `json:"params,omitempty"`
}

func (a *BlockHeightReq) Init(id int) {
	a.init()
	a.Id = id
	a.Method = "blockHeight"
}

func (a *BlockHeightReq) Vaildate() error {
	return nil
}

//
type TransactionByBlockHeight struct {
	CoinCode    string `json:"coinCode"`
	BlockHeight uint64 `json:"blockHeight"`
}
type TransactionByBlockHeightReq struct {
	OpenRequest
	Params TransactionByBlockHeight `json:"params,omitempty"`
}

func (a *TransactionByBlockHeightReq) Init(id int) {
	a.init()
	a.Id = id
	a.Method = "transactionByBlockHeight"
}

func (a *TransactionByBlockHeightReq) Vaildate() error {
	return nil
}
