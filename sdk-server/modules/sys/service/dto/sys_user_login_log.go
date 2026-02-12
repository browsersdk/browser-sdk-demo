package dto

import (
	"dilu/modules/sys/models"

	"github.com/baowk/dilu-core/core/base"
)

type SysUserLoginLogGetPageReq struct {
	base.ReqPage `query:"-"`
	SortOrder    string `json:"-" query:"type:order;column:id"`
	Uid          int    `json:"uid" query:"column:uid"` //用户id
	Ip           string `json:"ip" query:"column:ip"`   //登录ip
}

func (SysUserLoginLogGetPageReq) TableName() string {
	return models.TBSysUserLoginLog
}

// SysUserLoginLog
type SysUserLoginLogDto struct {
	Id        int    `json:"id"`        //主键
	Uid       uint   `json:"uid"`       //用户id
	Method    int8   `json:"method"`    //登录方式
	Ip        string `json:"ip"`        //登录ip
	Region    string `json:"region"`    //地域
	ClientId  string `json:"clientId"`  //客户端
	ClientVer string `json:"clientVer"` //操作系统
	Os        string `json:"os"`        //操作系统
	OsVer     string `json:"osVer"`     //操作系统版本
}
