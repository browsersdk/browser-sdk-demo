package models

import (
	"time"

	"gorm.io/gorm"
)

// 部门
type AdminDept struct {
	Id        int            `json:"id" gorm:"type:int unsigned;primaryKey;autoIncrement;comment:主键"` //主键
	ParentId  int            `json:"parentId" gorm:"type:int unsigned;comment:父id"`                   //父id
	DeptPath  string         `json:"deptPath" gorm:"type:varchar(255);comment:部门路径"`                  //部门路径
	Name      string         `json:"name" gorm:"type:varchar(128);comment:部门名"`                       //部门名
	Type      int8           `json:"type" gorm:"type:tinyint;comment:类型 1分公司 2部门"`                    //类型 1分公司 2部门
	Principal string         `json:"principal" gorm:"type:varchar(128);comment:部门领导"`                 //部门领导
	Phone     string         `json:"phone" gorm:"type:varchar(11);comment:手机号"`                       //手机号
	Email     string         `json:"email" gorm:"type:varchar(128);comment:邮箱"`                       //邮箱
	Sort      int8           `json:"sort" gorm:"type:tinyint;comment:排序"`                             //排序
	Status    int8           `json:"status" gorm:"type:tinyint;comment:状态 1正常 2关闭"`                   //状态 1正常 2关闭
	Remark    string         `json:"remark" gorm:"type:varchar(255);comment:备注"`                      //备注
	CreateBy  int            `json:"createBy" gorm:"type:int unsigned;comment:创建者"`                   //创建者
	UpdateBy  int            `json:"updateBy" gorm:"type:int unsigned;comment:更新者"`                   //更新者
	CreatedAt time.Time      `json:"createdAt" gorm:"type:datetime(3);comment:创建时间"`                  //创建时间
	UpdatedAt time.Time      `json:"updatedAt" gorm:"type:datetime(3);comment:最后更新时间"`                //最后更新时间
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index;comment:删除时间"`                                     //删除时间
}

const TBAdminDept = "admin_dept"

func (AdminDept) TableName() string {
	return TBAdminDept
}
