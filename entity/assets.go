package entity

type QueryAssetsReq struct {
	OpenRequest
	Params interface{} `json:"params,omitempty"`
}

func (a *QueryAssetsReq) Init(id int) {
	a.init()
	a.Id = id
	a.Method = "queryAssets"
}

func (a *QueryAssetsReq) Vaildate() error {
	return nil
}

//
//
//
type QueryAssetByCode struct {
	AssetCode string `json:"assetCode"`
}

type QueryAssetByCodeReq struct {
	OpenRequest
	Params QueryAssetByCode `json:"params,omitempty"`
}

func (a *QueryAssetByCodeReq) Init(id int) {
	a.init()
	a.Id = id
	a.Method = "queryAssetByCode"
}

func (a *QueryAssetByCodeReq) Vaildate() error {
	return nil
}

//
//

type QueryCoinsReq struct {
	OpenRequest
	Params interface{} `json:"params,omitempty"`
}

func (a *QueryCoinsReq) Init(id int) {
	a.init()
	a.Id = id
	a.Method = "queryCoins"
}
func (a *QueryCoinsReq) Vaildate() error {
	return nil
}

//queryCoinByCode
//
type QueryCoinByCode struct {
	CoinCode string `json:"coinCode"`
}

type QueryCoinByCodeReq struct {
	OpenRequest
	Params QueryCoinByCode `json:"params,omitempty"`
}

func (a *QueryCoinByCodeReq) Init(id int) {
	a.init()
	a.Id = id
	a.Method = "queryCoinByCode"
}
func (a *QueryCoinByCodeReq) Vaildate() error {
	return nil
}

//queryCoinByCode
type QueryAddressesByCoinCode struct {
	CoinCode string `json:"coinCode"`
	PageNum  int    `json:"pageNum,omitempty"`
	PageSize int    `json:"pageSize,omitempty"`
}
type QueryAddressesByCoinCodeReq struct {
	OpenRequest
	Params QueryAddressesByCoinCode `json:"params,omitempty"`
}

func (a *QueryAddressesByCoinCodeReq) Init(id int) {
	a.init()
	a.Id = id
	a.Method = "queryAddressesByCoinCode"
}
func (a *QueryAddressesByCoinCodeReq) Vaildate() error {
	return nil
}

//queryCoinByCode
type Address struct {
	Address string `json:"address"`
	Memo    string `json:"memo,omitempty"`
}
type QueryAddressesInfo struct {
	CoinCode    string    `json:"coinCode"`
	AddressList []Address `json:"addressList"`
}
type QueryAddressesInfoReq struct {
	OpenRequest
	Params QueryAddressesInfo `json:"params,omitempty"`
}

func (a *QueryAddressesInfoReq) Init(id int) {
	a.init()
	a.Id = id
	a.Method = "queryAddressesInfo"
}
func (a *QueryAddressesInfoReq) Vaildate() error {
	return nil
}

//checkAddress
type CheckAddress struct {
	CoinCode string `json:"coinCode"`
	Address  string `json:"address"`
	Memo     string `json:"memo,omitempty"`
}
type CheckAddressReq struct {
	OpenRequest
	Params CheckAddress `json:"params,omitempty"`
}

func (a *CheckAddressReq) Init(id int) {
	a.init()
	a.Id = id
	a.Method = "checkAddress"
}
func (a *CheckAddressReq) Vaildate() error {
	return nil
}

//getDepositAddress
type GetDepositAddress struct {
	CoinCode string `json:"coinCode"`
	Limit    uint64 `json:"limit,omitempty"`
}
type GetDepositAddressReq struct {
	OpenRequest
	Params GetDepositAddress `json:"params,omitempty"`
}

func (a *GetDepositAddressReq) Init(id int) {
	a.init()
	a.Id = id
	a.Method = "getDepositAddress"
}
func (a *GetDepositAddressReq) Vaildate() error {
	return nil
}
