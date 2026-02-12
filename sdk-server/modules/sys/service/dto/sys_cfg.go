package dto

import (
	"dilu/modules/sys/models"

	"github.com/baowk/dilu-core/core/base"
)

type SysCfgGetPageReq struct {
	base.ReqPage `query:"-"`
	SortOrder    string `json:"-" query:"type:order;column:id"`
	TeamId       int    `json:"teamId" query:"type:eq;column:team_id"`
	Status       int    `json:"status" query:""` //Status
	Type         string `json:"type" query:"type:eq;column:type"`
	Key          string `json:"key" query:"type:eq;column:key"`
}

func (SysCfgGetPageReq) TableName() string {
	return models.TBSysCfg
}

// 配置
type SysCfgDto struct {
	Id     int    `json:"id"`     //主键
	TeamId int    `json:"teamId"` // 团队id -1为平台
	Type   string `json:"type"`   //类型
	Key    string `json:"key"`    //类型下key
	Val    string `json:"val"`    //值
	Status int    `json:"status"` //Status
	Remark string `json:"remark"` //备注

}
