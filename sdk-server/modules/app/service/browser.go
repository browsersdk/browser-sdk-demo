package service

import (
	"context"
	"dilu/common/config"
	"dilu/common/utils"
	"dilu/modules/sys/models"
	"dilu/modules/sys/service"
	"encoding/json"
	"fmt"

	"github.com/baowk/dilu-core/common/consts"
	"github.com/baowk/dilu-core/core/base"
	brosdk "github.com/browsersdk/brosdk-server-go"
	"github.com/jinzhu/copier"
)

type BrowserService struct {
	*base.BaseService
}

var SerAppBrowser = BrowserService{
	base.NewService(consts.DB_DEF),
}

var sdk *brosdk.Client

func (s *BrowserService) getBroSdk() (*brosdk.Client, error) {
	if config.Ext.BroSdk.ApiKey == "" {
		return nil, fmt.Errorf("请配置 bro sdk api key")
	}

	var err error
	if sdk == nil {
		if config.Ext.BroSdk.Endpoint == "" {
			sdk, err = brosdk.NewClient(config.Ext.BroSdk.ApiKey)
		} else {
			sdk, err = brosdk.NewClient(config.Ext.BroSdk.ApiKey, brosdk.WithEndpoint(config.Ext.BroSdk.Endpoint))
		}
	}
	return sdk, err
}

func (s *BrowserService) GetUserSig(uid int, data *brosdk.UserSigData) error {
	sdk, err := s.getBroSdk()
	if err != nil {
		return err
	}
	req := brosdk.GetUserSigRequest{
		CustomerId: fmt.Sprintf("%d", uid),
		Duration:   86400,
	}

	resp, err := sdk.GetUserSig(context.Background(), &req)
	if err != nil {
		return err
	}
	*data = *resp

	return nil
}

func (s *BrowserService) Create(uid int, req *brosdk.EnvInfo) error {
	var data models.Browser
	copier.Copy(&data, req)

	data.UserId = uid
	data.EnvId = utils.GenUint()
	data.Status = 1

	if req.EnvName == "" {
		req.EnvName = fmt.Sprintf("用户%d的浏览器环境", uid)
	}
	data.Name = req.EnvName
	if req.KernelVersion == "" {
		req.KernelVersion = "134"
	}
	if req.Kernel == "" {
		req.Kernel = "Chrome"
	}
	if req.System == "" {
		req.System = "Windows 11"
	}

	if req.PublicIp == "" {
		req.PublicIp = "123.123.123.123"
	}

	if err := service.SerBrowser.Create(&data); err != nil {
		return err
	}

	sdk, err := s.getBroSdk()
	if err != nil {
		return err
	}

	req.CustomerId = fmt.Sprintf("%d", uid)

	resp, err := sdk.EnvCreate(context.Background(), req)
	if err != nil {
		return err
	}

	edata, err := json.Marshal(resp)
	if err != nil {
		return err
	}
	data.Data = string(edata)
	data.Status = 3
	if err := service.SerBrowser.UpdateById(&data); err != nil {
		return err
	}

	return nil
}

func (s *BrowserService) Update(uid int, req *brosdk.EnvInfo) error {
	var browser models.Browser
	if err := service.SerBrowser.Get(req.EnvId, &browser); err != nil {
		return err
	}

	if browser.UserId != uid {
		return fmt.Errorf("无权限")
	}

	var curEnv brosdk.EnvInfo
	if err := json.Unmarshal([]byte(browser.Data), &curEnv); err != nil {
		return err
	}
	sdk, err := s.getBroSdk()
	if err != nil {
		return err
	}
	env, err := sdk.EnvUpdate(context.Background(), req)
	if err != nil {
		return err
	}
	edata, err := json.Marshal(env)
	if err != nil {
		return err
	}
	browser.Data = string(edata)
	if err := service.SerBrowser.UpdateById(&browser); err != nil {
		return err
	}

	return nil
}
