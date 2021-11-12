package crypto

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"strings"
)

var RSA_KEY_SIZE int = 2048

func GenRsaKey() (string, string, error) {

	privateKey, err := rsa.GenerateKey(rand.Reader, RSA_KEY_SIZE)
	if err != nil {
		fmt.Println(err)
		return "", "", err
	}
	derStream, _ := x509.MarshalPKCS8PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	prvkey := pem.EncodeToMemory(block)
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		fmt.Println(err)
		return "", "", err
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	pubkey := pem.EncodeToMemory(block)

	// 得到公钥
	pubStr := base64.StdEncoding.EncodeToString(pubkey)
	prvStr := base64.StdEncoding.EncodeToString(prvkey)

	return pubStr, prvStr, nil
}

func RsaSign(prvKey string, data string) (string, error) {
	p, err := base64.StdEncoding.DecodeString(prvKey)
	if err != nil {
		fmt.Println("1 err:", err)
		return "", err
	}

	// h := sha1.New()
	// h.Write(d)
	// hashed := h.Sum(nil)

	hashed := sha256.Sum256([]byte(data))
	// hash := sha1.Sum() ..Sum256(d)

	privateKey, err := x509.ParsePKCS8PrivateKey(p)
	if err != nil {
		fmt.Println("ParsePKCS8PrivateKey err:", err)
		return "", err
	}

	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey.(*rsa.PrivateKey), crypto.SHA256, hashed[:])
	if err != nil {
		fmt.Printf("Error from signing: %s\n", err)
		return "", err
	}
	out := strings.ToUpper(hex.EncodeToString(signature))

	return out, nil
}

func RsaVerify(data, sign, pubk string) error {
	pk, err := base64.StdEncoding.DecodeString(pubk)
	if err != nil {
		return err
	}
	pubKey, err := x509.ParsePKIXPublicKey(pk)
	if err != nil {
		return err
	}

	hashed := sha256.Sum256([]byte(data))

	signByte, err := hex.DecodeString(sign)
	if err != nil {
		return err
	}
	err = rsa.VerifyPKCS1v15(pubKey.(*rsa.PublicKey), crypto.SHA256, hashed[:], signByte)
	if err != nil {
		return err
	}
	return nil
}
