package dto

import (
	"dilu/modules/sys/models"

	"github.com/baowk/dilu-core/core/base"
)

type UploadLogGetPageReq struct {
	base.ReqPage `query:"-"`
	SortOrder    string `json:"-" query:"type:order;column:id"`
	Status       int8   `json:"status" query:"column:status"` //1使用中 -1 删除SyncStatus int8 `json:"syncStatus" query:"column:sync_status"` //同步状态

}

func (UploadLogGetPageReq) TableName() string {
	return models.TBUploadLog
}

// UploadLog
type UploadLogDto struct {
	Id         int    `json:"id"`         //主键
	Url        string `json:"url"`        //资源地址
	Provider   string `json:"provider"`   //提供方
	Status     int8   `json:"status"`     //1使用中 -1 删除
	SyncStatus int8   `json:"syncStatus"` //同步状态
}
