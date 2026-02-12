package dto

import (
	"github.com/baowk/dilu-core/core/base"
)

type BrowserGetPageReq struct {
	base.ReqPage `query:"-"`
	Name         string `json:"name" form:"name"`     //名称
	EnvId        uint64 `json:"envId" form:"envId"`   //环境ID
	UserId       int    `json:"userId" form:"userId"` //用户ID
}

// Browser
type BrowserDto struct {
	Id     int    `json:"id"`     //主键
	Name   string `json:"name"`   //名称
	EnvId  uint64 `json:"envId"`  //环境ID
	UserId int    `json:"userId"` //用户ID
	Data   string `json:"data"`   //数据
	Status int8   `json:"status"` //状态 1 停止 3 启动
}
