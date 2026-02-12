package utils

import (
	"dilu/common/config"
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/baowk/dilu-core/common/utils/cryptos"
	"github.com/gin-gonic/gin"
)

func EncodeTeamId(teamId int) (string, error) {
	msg := fmt.Sprintf("%d-%d", teamId, time.Now().Unix())
	en, err := cryptos.AesEncryptCBC([]byte(msg), []byte(config.Ext.AesKey))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(en), nil
}

func GetTeamId(c *gin.Context) int {
	if GetRoleId(c) != 0 {
		return -1
	}
	sTeamId := c.GetHeader("teamId")
	if sTeamId == "" {
		sTeamId = c.GetString("teamId")
	}
	if sTeamId != "" {
		tid, _, err := DecodeTeamKey(sTeamId)
		if err != nil {
			return 0
		}
		return tid
	}
	return 0
}

func DecodeTeamKey(appKey string) (int, int64, error) {
	appKeyBytes, err := hex.DecodeString(appKey)
	if err != nil {
		return 0, 0, err
	}
	de, err := cryptos.AesDecryptCBC(appKeyBytes, []byte(config.Ext.AesKey))
	if err != nil {
		return 0, 0, err
	}
	arr := strings.Split(string(de), "-")
	if len(arr) != 2 {
		return 0, 0, errors.New("appKey error")
	}
	teamId, err := strconv.Atoi(arr[0])
	if err != nil {
		return 0, 0, err
	}
	expire, err := strconv.ParseInt(arr[1], 10, 64)
	if err != nil {
		return 0, 0, err
	}
	return teamId, expire, nil
}

func DecodeAppKey(appKey string) (int, int, int64, error) {
	appKeyBytes, err := hex.DecodeString(appKey)
	if err != nil {
		return 0, 0, 0, err
	}
	de, err := cryptos.AesDecryptCBC(appKeyBytes, []byte(config.Ext.AesKey))
	if err != nil {
		return 0, 0, 0, err
	}
	arr := strings.Split(string(de), "-")
	if len(arr) != 3 {
		return 0, 0, 0, errors.New("appKey error")
	}
	teamId, err := strconv.Atoi(arr[0])
	if err != nil {
		return 0, 0, 0, err
	}

	appId, err := strconv.Atoi(arr[1])
	if err != nil {
		return 0, 0, 0, err
	}

	expire, err := strconv.ParseInt(arr[2], 10, 64)
	if err != nil {
		return 0, 0, 0, err
	}
	return teamId, appId, expire, nil
}

func GetReqTeamId(c *gin.Context, reqTeamId int) int {
	teamId := GetTeamId(c)
	if teamId == -1 {
		if reqTeamId == 0 {
			return teamId
		}
		return reqTeamId
	}
	return teamId
}

func SetAppTeamId(appTeamId int, c *gin.Context) {
	c.Set("appTeamId", appTeamId)
}

func GetAppTeamId(c *gin.Context) int {
	if v, ok := c.Get("appTeamId"); ok {
		if v != nil {
			return v.(int)
		}
	}
	return 0
}

func SetAppId(appId int, c *gin.Context) {
	c.Set("appId", appId)
}

func GetAppId(c *gin.Context) int {
	if v, ok := c.Get("appId"); ok {
		if v != nil {
			return v.(int)
		}
	}
	return 0
}
