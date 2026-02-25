package service

import (
	"context"
	"dilu/common/config"
	"fmt"

	"github.com/baowk/dilu-core/common/consts"
	"github.com/baowk/dilu-core/core/base"
	brosdk "github.com/browsersdk/brosdk-server-go"
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
	}

	resp, err := sdk.GetUserSig(context.Background(), &req)
	if err != nil {
		return err
	}
	*data = *resp

	return nil
}
