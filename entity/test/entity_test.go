package entity

import (
	"apigo/entity"
	"apigo/service"
	"encoding/json"
	"testing"
)

func TestEntity(t *testing.T) {
	DefUrl := "http://10.0.152.131:8084/open/api/v1"
	prv := `MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQC960PRWoig87lgEhVRPSCd8wNFxT3FHs8DrT9mRjyb0bTlZHOqpR846AZfWlLJn2pbFQCILPlHnKY51HHoMPd0VuaZrBRa62rmVhrlNRsMYV/RsZZWaYloUcnDMyvqDP2Xg6CI4dNXSwl8vJDLzJHiwD+ZWE/AxjZftA0QqdQyuaRxwoabbZiYHtFxynN14KoEhhp6ElWd3S+b/4O7pg99BfqOnZtXNEJiUgLdDjrqGr6HgxZP7djPhQFgx9m4vVvNIfjUxGS35UcvrNJTT/QrLguQ84g+UHi1/Yz7W/G9yyPKvVuTmH3v+i3JBEt/Yk9H/o9J8k8tOLxOOvc4PFufAgMBAAECggEAHiWRP1MyqvHlNCXKsnmUit3/X/zeQEMSs0+156Mwjb9tCpi2b5kEasER+eLZj125wTmFOxiAfWiVTkq1xhi3vwToV+5j0Mbb0jJK1KVoZL+0ORIKfi8Ee8W31D90MhPb0Ug4nGHUbV/g4qcACQmxAqBnuAy3pC8ShICNLl726xIczUsn1MKbzavZNEAejYYs9YMaxyNwg3Qiw9SgCpM3UG85VIV6XuMnj1NypX6pGDv9bu3Sd4SI+OmfaG4GtgsDLE4aP7UMC848qnUnOASOkMmJRcfNWzceFkYCotJ6EUu3GraYdqTRzkz7HCkMCxTYhcanEFWKy/ZLn9VtiheJkQKBgQDiAnE6sxyyKNDJxgI53ELrBx80HgAJy3JcA/u3KN6xkMg0k6iji9X6XP0XDGBLSrnDJcpiJUxiYuVXyhntFS8D2FYIkcNuaPgS5aqzvmBm5PFpqRrCtZiOWqW/t1udGUt3RRt7lyjyUGLfh2oJHHpjjjfqymUB9xLUH+rfRrBtCQKBgQDXHtMtDP8DtjW4VNuSa45g64jmv1YAjD47HydpE76VnrZC5fI259oMzSzopGCA/H6YCkHinl0KHLsE3APNx0T7ID+CkAuGcVJuX0Y38t8AL8RYfciAjn1damBRig7tEmeY38FYN12YSzCa7qTo4gKuBHUxbiOUyNxDuHZmQkjVZwKBgQDU4XwJ0F/9KBjRlWLPYTre6gxoKMHcd/c12MKmGSb5legeLd8wfSyF8ESsCwpAoRgsSlJA8+To30Iq2MBm4gcw1frjg3jTbKgOFKofN/jRsl/6KEB+mlIh9BwfYvQ2G3dL1po2ZYE6DKG07nXgMyTM1U6yJwXRPgpMJ+wxdwIDEQKBgFhxE+ExtqaQAwYF3UAVeDPgoig8Ad+3yN4FsO5Cb9iTp9tZLnvkVoFs2UnMSuC87k8T6IKDGT1PEpSs3+N6SaH1YCcNka90Z694/CWEdKpe+RponEY+TsxZL8BWQky1hGIVnCfom1JBl3obIzGbuf5RVt07quVAr04oSIVCOy+BAoGARAUBCrjprBhn53OkWt6YjmUqMyT4cgIeo+zIjr/n7+TRrcvpFK6/MabSLlJ5AHsU1N95mXjrBv2IIM9TgybtNkK228IOjs5aBZiIDc3DMkHN6tneO+61QaM9oBPW/MR3bgL3X5uai8fy14nNVFOJA7eTriQ1ENTtybxhNyeDh7I=`
	keystr := `83c846b2433348da837b3aeb2e76fea6`

	or := service.NewOpenService(prv, keystr, DefUrl)

	//
	//1.queryAssets
	qar := &entity.QueryAssetsReq{}
	rlt, err := or.Call(1231, qar)
	if err != nil {
		t.Fatal("OpenService err:", err)
	}
	qarp := &entity.QueryAssetsResp{}
	err = json.Unmarshal(rlt, qarp)
	if err != nil {
		t.Fatal("QueryAssetsResp err:", err)
	}
	//2.queryAssetByCode
	qabcr := &entity.QueryAssetByCodeReq{}
	rlt, err = or.Call(1231, qabcr)
	if err != nil {
		t.Fatal("OpenService err:", err)
	}
	qabcrp := &entity.QueryAssetByCodeResp{}
	err = json.Unmarshal(rlt, qabcrp)
	if err != nil {
		t.Fatal("QueryAssetsResp err:", err)
	}
	//3.queryCoins
	qcr := &entity.QueryCoinsReq{}
	rlt, err = or.Call(1231, qcr)
	if err != nil {
		t.Fatal("OpenService err:", err)
	}
	qcrp := &entity.QueryCoinsResp{}
	err = json.Unmarshal(rlt, qcrp)
	if err != nil {
		t.Fatal("QueryAssetsResp err:", err)
	}
	//4.queryCoinByCode
	qcbcr := &entity.QueryCoinByCodeReq{}
	rlt, err = or.Call(1231, qcbcr)
	if err != nil {
		t.Fatal("OpenService err:", err)
	}
	qcbcrp := &entity.QueryCoinByCodeResp{}
	err = json.Unmarshal(rlt, qcbcrp)
	if err != nil {
		t.Fatal("OpenService err:", err)
	}
	//5.queryAddressesByCoinCode
	qabccr := &entity.QueryAddressesByCoinCodeReq{}
	rlt, err = or.Call(1231, qabccr)
	if err != nil {
		t.Fatal("OpenService err:", err)
	}
	qcbccrp := &entity.QueryAddressesByCoinCodeResp{}
	err = json.Unmarshal(rlt, qcbccrp)
	if err != nil {
		t.Fatal("OpenService err:", err)
	}
	//6.queryAddressesInfo
	qair := &entity.QueryAddressesInfoReq{}
	rlt, err = or.Call(1231, qair)
	if err != nil {
		t.Fatal("OpenService err:", err)
	}
	qairp := &entity.QueryAddressesInfoResp{}
	err = json.Unmarshal(rlt, qairp)
	if err != nil {
		t.Fatal("OpenService err:", err)
	}
	//7.checkAddress
	car := &entity.CheckAddressReq{}
	rlt, err = or.Call(1231, car)
	if err != nil {
		t.Fatal("OpenService err:", err)
	}
	carp := &entity.CheckAddressResp{}
	err = json.Unmarshal(rlt, carp)
	if err != nil {
		t.Fatal("OpenService err:", err)
	}
	//8.getDepositAddress
	gdar := &entity.GetDepositAddressReq{}
	rlt, err = or.Call(1231, gdar)
	if err != nil {
		t.Fatal("OpenService err:", err)
	}
	gdarp := &entity.GetDepositAddressResp{}
	err = json.Unmarshal(rlt, gdarp)
	if err != nil {
		t.Fatal("OpenService err:", err)
	}

	//9.withdrawApply
	war := &entity.WithdrawApplyReq{}
	rlt, err = or.Call(1231, war)
	if err != nil {
		t.Fatal("OpenService err:", err)
	}
	warp := &entity.WithdrawApplyResp{}
	err = json.Unmarshal(rlt, warp)
	if err != nil {
		t.Fatal("OpenService err:", err)
	}
	//10.queryByReqOrderId
	qbror := &entity.QueryByReqOrderIdReq{}
	rlt, err = or.Call(1231, qbror)
	if err != nil {
		t.Fatal("OpenService err:", err)
	}
	qbrorp := &entity.QueryByReqOrderIdResp{}
	err = json.Unmarshal(rlt, qbrorp)
	if err != nil {
		t.Fatal("OpenService err:", err)
	}
	//
	//11.transactions
	tr := &entity.TransactionsReq{}
	rlt, err = or.Call(1231, tr)
	if err != nil {
		t.Fatal("OpenService err:", err)
	}
	trp := &entity.TransactionsResp{}
	err = json.Unmarshal(rlt, trp)
	if err != nil {
		t.Fatal("OpenService err:", err)
	}
	//12.transactionById
	tbir := &entity.TransactionByIdReq{}
	rlt, err = or.Call(1231, tbir)
	if err != nil {
		t.Fatal("OpenService err:", err)
	}
	tbirp := &entity.TransactionByIdResp{}
	err = json.Unmarshal(rlt, tbirp)
	if err != nil {
		t.Fatal("OpenService err:", err)
	}
	//13.pendingTransactions
	ptr := &entity.PendingTransactionsReq{}
	rlt, err = or.Call(1231, ptr)
	if err != nil {
		t.Fatal("OpenService err:", err)
	}
	ptrp := &entity.PendingTransactionsResp{}
	err = json.Unmarshal(rlt, ptrp)
	if err != nil {
		t.Fatal("OpenService err:", err)
	}
	//14.pendingTransactionById
	ptbir := &entity.PendingTransactionByIdReq{}
	rlt, err = or.Call(1231, ptbir)
	if err != nil {
		t.Fatal("OpenService err:", err)
	}
	ptbirp := &entity.PendingTransactionByIdResp{}
	err = json.Unmarshal(rlt, ptbirp)
	if err != nil {
		t.Fatal("OpenService err:", err)
	}
	//15.transactionByTxHash
	tbthr := &entity.TransactionByTxHashReq{}
	rlt, err = or.Call(1231, tbthr)
	if err != nil {
		t.Fatal("OpenService err:", err)
	}
	tbthrp := &entity.TransactionByTxHashResp{}
	err = json.Unmarshal(rlt, tbthrp)
	if err != nil {
		t.Fatal("OpenService err:", err)
	}
	//16.blockHeight)
	bhr := &entity.BlockHeightReq{}
	rlt, err = or.Call(1231, bhr)
	if err != nil {
		t.Fatal("OpenService err:", err)
	}
	bhrp := &entity.BlockHeightResp{}
	err = json.Unmarshal(rlt, bhrp)
	if err != nil {
		t.Fatal("OpenService err:", err)
	}
	//17.transactionByBlockHeight
	tbbhr := &entity.TransactionByBlockHeightReq{}
	rlt, err = or.Call(1231, tbbhr)
	if err != nil {
		t.Fatal("OpenService err:", err)
	}
	tbbhrp := &entity.TransactionByBlockHeightResp{}
	err = json.Unmarshal(rlt, tbbhrp)
	if err != nil {
		t.Fatal("OpenService err:", err)
	}
}
