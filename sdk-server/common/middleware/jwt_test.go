package middleware

import (
	"dilu/common/utils"
	"testing"
	"time"
)

func TestGen(t *testing.T) {
	key := "12e96bbf8e03035893ea92b7cf12c000"
	uid := 1
	issuer := "issuer"
	subject := "subject"

	expiresAt := time.Now().Add(time.Hour * 24)

	token, err := utils.GenerateAppToken(uid, issuer, subject, key, expiresAt)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(token)

	var customClaims utils.AppClaims
	if err = utils.ParseApp(token, &customClaims, key); err != nil {
		t.Fatal(err)
	}
	t.Log(customClaims)

	refreshT, err := utils.GenerateAppRefreshToken(uid, issuer, subject, key, expiresAt)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(refreshT)

	var ccR utils.AppClaims
	if err = utils.ParseAppRefresh(refreshT, &ccR, key); err != nil {
		t.Fatal(err)
	}
	t.Log(ccR)

	roleId := 1
	adminUid := 1

	tokena, err := utils.GenerateAdminToken(adminUid, roleId, issuer, subject, key, expiresAt)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tokena)

	var cca utils.AdminClaims
	if err = utils.ParseAdmin(tokena, &cca, key); err != nil {
		t.Fatal(err)
	}
	t.Log(cca)

	refreshaT, err := utils.GenerateAdminRefreshToken(adminUid, roleId, issuer, subject, key, expiresAt)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(refreshaT)

	var ccaR utils.AdminClaims
	if err = utils.ParseAdminRefresh(refreshaT, &ccaR, key); err != nil {
		t.Fatal(err)
	}
	t.Log(ccaR)
}

func TestParse(t *testing.T) {
	var cClaims utils.AppClaims
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjEsImV4cCI6MTc0MTM0MzU1MH0.FKguK6W6trEveL5UZc_3Fbs2oNDzsWmbyHA5xplf2h4"
	key := "12e96bbf8e03035893ea92b7cf12c000"
	err := utils.ParseApp(token, &cClaims, key)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(cClaims)
}

func TestParseAdmin(t *testing.T) {
	var cClaims utils.AdminClaims
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6Ijc3YTkxZTQ2LWQ3NGYtNDNjZi1iYzg1LWQ3MGU1OGFhZTlmYSJ9.Ymp_WDowcZLjYMtuYfo0e-1_6d0svJiXFor2F7hW_Rs"
	key := "12e96bbf8e03035893ea92b7cf12c000"
	err := utils.ParseAdmin(token, &cClaims, key)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(cClaims)
}

func TestParseReAdmin(t *testing.T) {
	var cClaims utils.AdminClaims
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6Ijc3YTkxZTQ2LWQ3NGYtNDNjZi1iYzg1LWQ3MGU1OGFhZTlmYSJ9.Ymp_WDowcZLjYMtuYfo0e-1_6d0svJiXFor2F7hW_Rs"
	key := "12e96bbf8e03035893ea92b7cf12c000"
	err := utils.ParseAdminRefresh(token, &cClaims, key)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(cClaims)
}
