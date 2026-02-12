package models

import (
	"dilu/common/model"
	"time"

	"gorm.io/gorm"
)

// 角色
type AdminRole struct {
	Id        int            `json:"id" gorm:"type:int unsigned;primaryKey;autoIncrement;comment:主键"` //主键
	Name      string         `json:"name" gorm:"type:varchar(128);comment:角色名称"`                      //角色名称
	RoleKey   string         `json:"roleKey" gorm:"type:varchar(128);comment:角色代码"`                   //角色代码
	RoleSort  uint           `json:"roleSort" gorm:"type:int unsigned;comment:排序"`                    //排序
	Status    int8           `json:"status" gorm:"type:tinyint;comment:状态"`                           //状态
	Remark    string         `json:"remark" gorm:"type:varchar(255);comment:备注"`                      //备注
	MenuIds   model.Array    `json:"menuIds" gorm:"type:json;comment:菜单ID"`                           //菜单ID
	CreateBy  int            `json:"createBy" gorm:"type:int unsigned;comment:创建者"`                   //创建者
	UpdateBy  int            `json:"updateBy" gorm:"type:int unsigned;comment:更新者"`                   //更新者
	CreatedAt time.Time      `json:"createdAt" gorm:"type:datetime(3);comment:创建时间"`                  //创建时间
	UpdatedAt time.Time      `json:"updatedAt" gorm:"type:datetime(3);comment:最后更新时间"`                //最后更新时间
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index;comment:删除时间"`                                     //删除时间
}

const TBAdminRole = "admin_role"

func (AdminRole) TableName() string {
	return TBAdminRole
}
