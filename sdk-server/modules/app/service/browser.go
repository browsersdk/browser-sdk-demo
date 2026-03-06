package service

import (
	"context"
	"dilu/common/config"
	"dilu/modules/sys/models"
	"dilu/modules/sys/service"
	"dilu/modules/sys/service/dto"
	"fmt"

	"github.com/baowk/dilu-core/common/consts"
	"github.com/baowk/dilu-core/core/base"
	brosdk "github.com/browsersdk/brosdk-server-go"
	"github.com/jinzhu/copier"
	"gorm.io/datatypes"
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

func (s *BrowserService) Create(uid int, req *dto.BrowserDto) (data *models.Browser, err error) {
	if req.EnvName == "" {
		req.EnvName = fmt.Sprintf("用户%d的浏览器环境", uid)
	}

	data = &models.Browser{}
	copier.Copy(&data, req)

	data.UserId = uid
	data.Status = 1

	data.EnvName = req.EnvName

	if err = service.SerBrowser.Create(&data); err != nil {
		return
	}

	sdk, err := s.getBroSdk()
	if err != nil {
		return
	}

	req.Data.EnvName = req.EnvName
	req.Data.CustomerId = fmt.Sprintf("%d", uid)

	resp, err := sdk.EnvCreate(context.Background(), req.Data)
	if err != nil {
		return
	}

	data.EnvId = resp.EnvId
	data.Data = datatypes.NewJSONType(resp)
	data.Status = 3
	if err = service.SerBrowser.UpdateById(&data); err != nil {
		return
	}

	return
}

func (s *BrowserService) Update(uid int, req *dto.BrowserDto) (browser *models.Browser, err error) {
	browser = &models.Browser{}
	if err := service.SerBrowser.DB().Where("id = ?", req.Id).Find(browser).Error; err != nil {
		return nil, err
	}
	if req.EnvName != "" {
		browser.EnvName = req.EnvName
	}

	if browser.UserId != uid {
		return nil, fmt.Errorf("无权限")
	}

	sdk, err := s.getBroSdk()
	if err != nil {
		return nil, err
	}
	data := req.Data
	if data != nil {
		req.Data.EnvName = browser.EnvName
		env, err := sdk.EnvUpdate(context.Background(), data)
		if err != nil {
			return nil, err
		}
		browser.Data = datatypes.NewJSONType(env)
	}
	if err := service.SerBrowser.UpdateById(&browser); err != nil {
		return nil, err
	}

	return browser, nil
}
