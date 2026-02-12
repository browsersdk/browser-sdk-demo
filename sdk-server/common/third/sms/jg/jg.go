package jg

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"log"

	"dilu/common/https"
)

const (
	api_url = "https://api.verification.jpush.cn/v1/web/loginTokenVerify"
)

func GetMobile(token string, outId string) (string, error) {
	apiSk := "Basic " + base64.StdEncoding.EncodeToString([]byte("34d855993e16ce0432d933be:f07a757821fda43abbef89b5"))
	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": apiSk,
	}

	params := map[string]interface{}{
		"loginToken": token,
	}
	if outId != "" {
		params["exID"] = outId
	}

	data, err := json.Marshal(params)
	if err != nil {
		return "", err
	}

	client := https.NewByH(headers)
	res, err := client.Post(api_url, data)
	if err != nil {
		return "", err
	}

	fmt.Println("res:", string(res))

	var _res Res
	err = json.Unmarshal(res, &_res)
	if err != nil {
		return "", err
	}
	if _res.Code == 8000 {
		return decode(_res.Phone)
	} else {
		return "", fmt.Errorf("get mobile error: %s,code:%d", _res.Content, _res.Code)
	}
}

func decode(encrypted string) (string, error) {
	PREFIX := "-----BEGIN RSA PRIVATE KEY-----"
	SUFFIX := "-----END RSA PRIVATE KEY-----"
	prikey := `-----BEGIN PRIVATE KEY-----
MIICdQIBADANBgkqhkiG9w0BAQEFAASCAl8wggJbAgEAAoGBAKUcPRemJcRtQwO7
XMAhabxOT/a7DNyRcztWGrXcyvzChQoVRrT8w6NhupXxvDqUg2uPDk4/3FLEDSkd
d0L1oHgv6uFgYCixZc4qrli0utbzuakPVVpHdFWxwlNv6sAOFpjrfNvDzcMRRD1O
065t2D2d3XAYU6sTHDfo2l237teNAgMBAAECgYBlK/6LZTWzPThZKw/UcyT1TA9X
oppo9X9klohbc+W2KAOZgBwJfvDqRlIs2yl5w9Mbr1cWv67j0Fo4HWQc1aHzxxlZ
XzbY5SQ0i0Qgux8/Zlymt5xWE3OyyxZFCLkQS31SuDmpqkT0YOYfXBL3YykyS1V3
/zRs8wQRtIR5mkeBQQJBANm7C4pQoJ4+gD9JCTQj4wLgBLEWXDGLRix0Y955N6Vt
qZrrff6VqwxvtdnilXs+w72zTSCueZm1Splv3qJYcPUCQQDCIYBvfQmZNX8peC4X
HK43LUXXNTpxmDAnoEWvhO5qSBMRTTugleBEAF+c6Opl3YL5uqhzqIhd1/XwH8Hc
tE05AkAX/XD/E3UvGmndDzoRYabgqTg7HuCXoOXhfg2G9Mo50wRuCmZ+h5UqMKDT
2hBoXPkKTNhBspe5ZO1MTLQ2JRKFAkAxbHATckT4UYtfVm59idq8x3Tpdm67ruBL
pl03c2NzgYgBNqWAm8et6F5vR0ktx/hpdeEfGQAmzC2cBanIc8rZAkBrR4fKW6gL
jXmyk/2YHh8RZKUH6k6IF5ujsAtDNhiyB2g1qqsKpazA3kJ9/x4YY0h6+33A5L5r
6SgeomjORkfU
-----END PRIVATE KEY-----`

	encryptedB, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		log.Println("invalid encrypted")
		return "", err
	}
	key := PREFIX + "\n" + prikey + "\n" + SUFFIX
	result, err := RsaDecrypt(encryptedB, []byte(key))
	if err != nil {
		log.Println("err:", err)
		return "", err
	}
	log.Println("result:", string(result))
	return string(result), nil
}

// 私钥解密
func RsaDecrypt(encrypted, prikey []byte) ([]byte, error) {
	var data []byte
	block, _ := pem.Decode(prikey)
	if block == nil {
		return data, errors.New("private key error")
	}
	rsaKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return data, err
	}
	key, ok := rsaKey.(*rsa.PrivateKey)
	if !ok {
		return data, errors.New("invalid private key")
	}
	data, err = rsa.DecryptPKCS1v15(rand.Reader, key, encrypted)
	return data, err
}

type Res struct {
	Id      int64  `json:"id"`
	Code    int    `json:"code"`
	Content string `json:"content"`
	ExtId   string `json:"extId"`
	Phone   string `json:"phone"`
}
