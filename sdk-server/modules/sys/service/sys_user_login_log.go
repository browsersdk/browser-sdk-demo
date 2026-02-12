package service

import (
	"github.com/baowk/dilu-core/core/base"
)

type SysUserLoginLogService struct {
	*base.BaseService
}

var SerSysUserLoginLog = SysUserLoginLogService{
	base.NewService("sys"),
}

