package dto

import (
	"dilu/modules/sys/models"

	"github.com/baowk/dilu-core/core/base"
)

type SysEmailGetPageReq struct {
	base.ReqPage `query:"-"`
	SortOrder    string `json:"-" query:"type:order;column:id"`
	Status       int8   `json:"status" query:"column:status"` //状态UseStatus int8 `json:"useStatus" query:"column:use_status"` //使用状态

}

func (SysEmailGetPageReq) TableName() string {
	return models.TBSysEmail
}

// 邮件
type SysEmailDto struct {
	Id        int    `json:"id"`        //主键
	Email     string `json:"email"`     //邮箱地址
	Code      string `json:"code"`      //验证码
	Type      string `json:"type"`      //类型
	Status    int8   `json:"status"`    //状态
	UseStatus int8   `json:"useStatus"` //使用状态
}
