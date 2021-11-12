package entity

type QueryAssets struct {
}

func (a QueryAssets) Method() string {
	return "queryAssets"
}
func (a QueryAssets) Vaildate() error {
	return nil
}

//
type QueryAssetByCode struct {
	AssetCode string `json:"assetCode"`
}

func (a QueryAssetByCode) Method() string {
	return "queryAssetByCode"
}
func (a QueryAssetByCode) Vaildate() error {
	return nil
}

//
type QueryCoins struct {
}

func (a QueryCoins) Method() string {
	return "queryCoins"
}
func (a QueryCoins) Vaildate() error {
	return nil
}

//queryCoinByCode
type QueryCoinByCode struct {
	CoinCode string `json:"coinCode"`
}

func (a QueryCoinByCode) Method() string {
	return "queryCoinByCode"
}
func (a QueryCoinByCode) Vaildate() error {
	return nil
}

//queryCoinByCode
type QueryAddressesByCoinCode struct {
	CoinCode string `json:"coinCode"`
	PageNum  int    `json:"pageNum,omitempty"`
	PageSize int    `json:"pageSize,omitempty"`
}

func (a QueryAddressesByCoinCode) Method() string {
	return "queryAddressesByCoinCode"
}
func (a QueryAddressesByCoinCode) Vaildate() error {
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

func (a QueryAddressesInfo) Method() string {
	return "queryAddressesInfo"
}
func (a QueryAddressesInfo) Vaildate() error {
	return nil
}

//checkAddress
type CheckAddress struct {
	CoinCode string `json:"coinCode"`
	Address  string `json:"address"`
	Memo     string `json:"memo,omitempty"`
}

func (a CheckAddress) Method() string {
	return "checkAddress"
}
func (a CheckAddress) Vaildate() error {
	return nil
}

//getDepositAddress
type GetDepositAddress struct {
	CoinCode string `json:"coinCode"`
	Limit    uint64 `json:"limit,omitempty"`
}

func (a GetDepositAddress) Method() string {
	return "getDepositAddress"
}
func (a GetDepositAddress) Vaildate() error {
	return nil
}
