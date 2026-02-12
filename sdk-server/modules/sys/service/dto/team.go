package dto

import (
	"time"

	"github.com/baowk/dilu-core/core/base"
)

type SysTeamGetPageReq struct {
	base.ReqPage `query:"-"`
	Status       int    `json:"status" form:"status"` //状态
	Name         string `json:"name" form:"name"`     // 团队名
}

// 团队
type SysTeamDto struct {
	Id     int    `json:"id"`     //主键
	Name   string `json:"name"`   //团队名
	Owner  int    `json:"owner"`  //团队拥有者
	Status int    `json:"status"` //状态
}

type SysTeamData struct {
	Id        int       `json:"id"`        //主键
	Name      string    `json:"name"`      //团队名
	Owner     int       `json:"owner"`     //团队拥有者
	Status    int       `json:"status"`    //状态
	AppSk     string    `json:"appSk"`     //应用密钥
	UpdateBy  int       `json:"updateBy"`  // 更新人
	CreatedAt time.Time `json:"createdAt"` // 创建时间
	UpdatedAt time.Time `json:"updatedAt"` // 更新时间
}
