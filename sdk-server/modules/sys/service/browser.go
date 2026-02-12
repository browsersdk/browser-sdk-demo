package service

import (
	"github.com/baowk/dilu-core/common/consts"
	"github.com/baowk/dilu-core/core/base"
)

type BrowserService struct {
	*base.BaseService
}

var SerBrowser = BrowserService{
	base.NewService(consts.DB_DEF),
}