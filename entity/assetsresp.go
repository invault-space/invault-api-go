package entity

import "encoding/json"

//
///
///
type Asset struct {
	AssetCode    string `json:"assetCode"`
	FullLogoUrl  string `json:"fullLogoUrl"`
	AssetName    string `json:"assetName"`
	CanDeposit   bool   `json:"canDeposit"`
	CanWithdraw  bool   `json:"canWithdraw"`
	AmountUsable string `json:"amountUsable"`
	AmountFrozen string `json:"amountFrozen"`
	AmountAll    string `json:"amountAll"`
}
type QueryAssetsResult struct {
	AssetList []Asset `json:"assetList"`
	Status    int     `json:"status"`
}
type QueryAssetsResp struct {
	OpenResult
	Reuslt QueryAssetsResult `json:"result,omitempty"`
}

func (a *QueryAssetsResp) ParseResult(j []byte) (*QueryAssetsResp, error) {
	err := json.Unmarshal(j, a)
	return a, err
}

///
///
///

type Coin struct {
	Network         string `json:"network"`
	ChainCoin       string `json:"chainCoin"`
	DisplayDecimals int    `json:"displayDecimals"`
	CoinAlias       string `json:"coinAlias"`
	CoinCode        string `json:"coinCode"`
	AmountUsable    string `json:"amountUsable"`
	AmountFrozen    string `json:"amountFrozen"`
	AmountAll       string `json:"amountAll"`
	CanDeposit      bool   `json:"canDeposit"`
	CanWithdraw     bool   `json:"canWithdraw"`
}
type QueryAssetByCodeResult struct {
	AssetCode    string `json:"assetCode"`
	FullLogoUrl  string `json:"fullLogoUrl"`
	AssetName    string `json:"assetName"`
	CanDeposit   bool   `json:"canDeposit"`
	CanWithdraw  bool   `json:"canWithdraw"`
	AmountUsable string `json:"amountUsable"`
	AmountFrozen string `json:"amountFrozen"`
	AmountAll    string `json:"amountAll"`
	CoinList     []Coin `json:"coin"`
}
type QueryAssetByCodeResp struct {
	OpenResult
	Reuslt QueryAssetByCodeResult `json:"result,omitempty"`
}

func (a *QueryAssetByCodeResp) ParseResult(j []byte) (*QueryAssetByCodeResp, error) {

	err := json.Unmarshal(j, a)
	return a, err
}

///
///
///
type QueryCoinsResult struct {
	AssetCode       string `json:"assetCode"`
	FullLogoUrl     string `json:"fullLogoUrl"`
	CoinCode        string `json:"coinCode"`
	CoinName        string `json:"coinName"`
	CoinAlias       string `json:"coinAlias"`
	NetworkName     string `json:"networkName"`
	Decimals        int64  `json:"decimals"`
	DisplayDecimals string `json:"displayDecimals"`
	CanDeposit      bool   `json:"canDeposit"`
	CanWithdraw     bool   `json:"canWithdraw"`
	Network         string `json:"network"`
	ChainCoin       string `json:"chainCoin"`
	FeeCoin         string `json:"feeCoin"`
	FeeBusiness     string `json:"feeBusiness"`
	FeeTransaction  string `json:"feeTransaction"`
	WithdrawMin     string `json:"withdrawMin"`
	DepositMin      string `json:"depositMin"`
	WithdrawMax     string `json:"withdrawMax"`
	TokenAddress    string `json:"tokenAddress"`
	IsToken         bool   `json:"isToken"`
	IsMemo          bool   `json:"isMemo"`
	MemoRegex       string `json:"memoRegex"`
	ChainType       string `json:"chainType"`
	Confirmations   int64  `json:"confirmations"`
	AmountUsable    string `json:"amountUsable"`
	AmountFrozen    string `json:"amountFrozen"`
	AmountAll       string `json:"amountAll"`
	AddressRegex    string `json:"addressRegex"`
	Label           string `json:"label"`
}
type QueryCoinsResp struct {
	OpenResult
	Reuslt []QueryCoinsResult `json:"result,omitempty"`
}

func (a *QueryCoinsResp) ParseResult(j []byte) (*QueryCoinsResp, error) {

	err := json.Unmarshal(j, a)
	return a, err
}

///
///
///
type QueryCoinByCodeResp struct {
	OpenResult
	Reuslt QueryCoinsResult `json:"result,omitempty"`
}

func (a *QueryCoinByCodeResp) ParseResult(j []byte) (*QueryCoinByCodeResp, error) {

	err := json.Unmarshal(j, a)
	return a, err
}

///
///
///
type AddressList struct {
	Address string `json:"address"`
	Memo    string `json:"memo"`
}
type QueryAddressesByCoinCodeResult struct {
	CoinCode    string        `json:"coinCode"`
	AddressList []AddressList `json:"addressList"`
	Size        int           `json:"size"`
	Total       int           `json:"total"`
	PageNum     int           `json:"pageNum"`
	PageSize    int           `json:"pageSize"`
}
type QueryAddressesByCoinCodeResp struct {
	OpenResult
	Reuslt QueryAddressesByCoinCodeResult `json:"result,omitempty"`
}

func (a *QueryAddressesByCoinCodeResp) ParseResult(j []byte) (*QueryAddressesByCoinCodeResp, error) {

	err := json.Unmarshal(j, a)
	return a, err
}

///
///
///
type QueryAddressesInfoResult struct {
	CoinCode    string    `json:"coinCode"`
	AddressList []Address `json:"addressList"`
}
type QueryAddressesInfoResp struct {
	OpenResult
	Reuslt QueryAddressesInfoResult `json:"result,omitempty"`
}

func (a *QueryAddressesInfoResp) ParseResult(j []byte) (*QueryAddressesInfoResp, error) {

	err := json.Unmarshal(j, a)
	return a, err
}

///
///
///
type CheckAddressResult struct {
	CoinCode string `json:"coinCode"`
	Address  string `json:"address"`
	Memo     string `json:"memo"`
	Status   int    `json:"status"`
}
type CheckAddressResp struct {
	OpenResult
	Reuslt CheckAddressResult `json:"result,omitempty"`
}

func (a *CheckAddressResp) ParseResult(j []byte) (*CheckAddressResp, error) {
	err := json.Unmarshal(j, a)
	return a, err
}

///
///
///
type GetDepositAddressResp struct {
	OpenResult
	Reuslt QueryAddressesInfoResult `json:"result,omitempty"`
}

func (a *GetDepositAddressResp) ParseResult(j []byte) (*GetDepositAddressResp, error) {
	err := json.Unmarshal(j, a)
	return a, err
}
