package dto

import (
	"github.com/baowk/dilu-core/core/base"
	"github.com/browsersdk/brosdk-server-go"
	"gorm.io/datatypes"
)

type BrowserGetPageReq struct {
	base.ReqPage `query:"-"`
	SortOrder    string `json:"-" query:"type:order;column:id"`
	EnvName      string `json:"envName" form:"envName"` //名称
	EnvId        string `json:"envId" form:"envId"`     //环境ID
	UserId       int    `json:"userId" form:"userId"`   //用户ID
}

// Browser
type BrowserDto struct {
	Id      int                                `json:"id"`      //主键
	EnvName string                             `json:"envName"` //名称
	EnvId   string                             `json:"envId"`   //环境ID
	UserId  int                                `json:"userId"`  //用户ID
	Data    datatypes.JSONType[brosdk.EnvInfo] `json:"data"`    //数据
	Status  int8                               `json:"status"`  //状态 1 停止 3 启动
}
