package service

import (
	"apigo/crypto"
	"apigo/entity"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type OpenRequest struct {
	Url  string
	Head string
	Sign string

	PrvKey string
	KeyStr string
}

func NewOpenRequest(prvKey, keyStr, url string) OpenRequest {
	return OpenRequest{
		PrvKey: prvKey,
		KeyStr: keyStr,
		Url:    url,
	}
}

func (or *OpenRequest) Call(a *entity.Abstract) (*OpenReuslt, error) {
	err := a.Vaildate()
	if err != nil {
		return nil, err
	}
	sign, err := crypto.RsaSign(or.PrvKey, a.String())
	if err != nil {
		return nil, err
	}
	result, err := or.HttpPost([]byte(a.String()), sign)

	ret := OpenReuslt{}
	err = json.Unmarshal(result, &ret)
	return &ret, err
}

func (or *OpenRequest) HttpPost(body []byte, sign string) ([]byte, error) {

	req, _ := http.NewRequest("POST", or.Url, bytes.NewBuffer(body))

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("sign", sign)
	req.Header.Set("sign", sign)
	req.Header.Set("keyStr", or.KeyStr)
	req.Header.Set("timeStamp", strconv.FormatInt(time.Now().Unix(), 10))

	res, err := (&http.Client{}).Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return content, nil
}
