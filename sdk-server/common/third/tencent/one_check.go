package tencent

import (
	"crypto/sha256"
	"dilu/common/config"
	"dilu/common/https"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"

	"strconv"
	"time"
)

type OneClickMobileLogin struct {
	client *https.HTTPClient
}

type Res struct {
	Result int    `json:"result"`
	Errmsg string `json:"errmsg"`
	Mobile string `json:"mobile"`
}

func New() *OneClickMobileLogin {
	oa := &OneClickMobileLogin{
		client: &(https.HTTPClient{
			BaseURL: "https://yun.tim.qq.com",
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}),
	}
	return oa
}

func (m *OneClickMobileLogin) Login(carrier, token, appType string) (string, error) {
	appid := config.Ext.OneClickLogin.AppId
	appKey := config.Ext.OneClickLogin.AppKey
	timeNow := strconv.Itoa(int(time.Now().Unix()))
	random := timeNow[:9]
	sig := m.GetSHA256Sig(appKey, random, timeNow)
	body := map[string]string{
		"time":    timeNow,
		"carrier": carrier,
		"token":   token,
		"sig":     sig,
	}
	data, err := json.Marshal(body)
	if err != nil {
		return "", err
	}
	response, err := m.client.Post("/v5/rapidauth/validate?sdkappid="+appid+"&random="+random+"", data)
	if err != nil {
		return "", err
	}

	var res Res
	if err = json.Unmarshal(response, &res); err != nil {
		return "", err
	}
	if res.Result == 0 {
		return res.Mobile, nil
	} else {
		return "", errors.New(res.Errmsg)
	}
}

func (m *OneClickMobileLogin) GetSHA256Sig(appkey, random, time string) string {
	sigStr := fmt.Sprintf("appkey=%s&random=%s&time=%s", appkey, random, time)
	hash := sha256.New()
	hash.Write([]byte(sigStr))
	hashInBytes := hash.Sum(nil)
	sig := hex.EncodeToString(hashInBytes)
	return sig
}
