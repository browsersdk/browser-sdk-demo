package dto

import (
	"dilu/modules/sys/models"
	"time"

	"github.com/baowk/dilu-core/core/base"
)

type AdminRoleGetPageReq struct {
	base.ReqPage `query:"-"`
	SortOrder    string `json:"-" query:"type:order;column:id"`
	Status       int8   `json:"status" query:"column:status"` //状态

}

func (AdminRoleGetPageReq) TableName() string {
	return models.TBAdminRole
}

// 角色
type AdminRoleDto struct {
	Id       int    `json:"id"`       //主键
	Name     string `json:"name"`     //角色名称
	RoleKey  string `json:"roleKey"`  //角色代码
	RoleSort uint   `json:"roleSort"` //排序
	Status   int8   `json:"status"`   //状态
	Remark   string `json:"remark"`   //备注
	MenuIds  []int  `json:"menuIds"`  // 菜单ids
}

type AdminRoleDtoResp struct {
	Id        int       `json:"id" `       //主键
	Name      string    `json:"name"`      //角色名称
	RoleKey   string    `json:"roleKey"`   //角色代码
	RoleSort  int       `json:"roleSort"`  //排序
	Status    int       `json:"status"`    //状态
	Remark    string    `json:"remark" `   //备注
	MenuIds   []int     `json:"menuIds"`   //菜单id
	CreateBy  int       `json:"createBy" ` //创建者
	UpdateBy  int       `json:"updateBy" ` //更新者
	CreatedAt time.Time `json:"createdAt"` //创建时间
	UpdatedAt time.Time `json:"updatedAt"` //最后更新时间
}
