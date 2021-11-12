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

type OpenService struct {
	Url  string
	Head string
	Sign string

	PrvKey string
	KeyStr string
}

func NewOpenService(prvKey, keyStr, url string) OpenService {
	return OpenService{
		PrvKey: prvKey,
		KeyStr: keyStr,
		Url:    url,
	}
}

func (or *OpenService) Call(id int, a entity.IOpenEntity) ([]byte, error) {
	a.Init(id)
	err := a.Vaildate()
	if err != nil {
		return nil, err
	}
	b, _ := json.Marshal(a)

	sign, err := crypto.RsaSign(or.PrvKey, string(b))
	if err != nil {
		return nil, err
	}
	result, err := or.HttpPost([]byte(string(b)), sign)

	return result, err
}

func (or *OpenService) HttpPost(body []byte, sign string) ([]byte, error) {

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
