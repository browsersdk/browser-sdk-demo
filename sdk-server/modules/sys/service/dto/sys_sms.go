package dto

import (
	"dilu/modules/sys/models"

	"github.com/baowk/dilu-core/core/base"
)

type SysSmsGetPageReq struct {
	base.ReqPage `query:"-"`
	SortOrder    string `json:"-" query:"type:order;column:id"`
	Status       int8   `json:"status" query:"column:status"` //状态UseStatus int8 `json:"useStatus" query:"column:use_status"` //使用状态

}

func (SysSmsGetPageReq) TableName() string {
	return models.TBSysSms
}

// 短信
type SysSmsDto struct {
	Id        int    `json:"id"`        //主键
	Phone     string `json:"phone"`     //手机号
	Code      string `json:"code"`      //验证码
	Type      string `json:"type"`      //类型
	Status    int8   `json:"status"`    //状态
	UseStatus int8   `json:"useStatus"` //使用状态
}
