package dto

import (
	"dilu/modules/sys/models"

	"github.com/baowk/dilu-core/core/base"
)

type AdminDeptGetPageReq struct {
	base.ReqPage `query:"-"`
	SortOrder    string `json:"-" query:"type:order;column:id"`
	Status       int8   `json:"status" query:"column:status"` //状态 1正常 2关闭

}

func (AdminDeptGetPageReq) TableName() string {
	return models.TBAdminDept
}

// 部门
type AdminDeptDto struct {
	Id        int    `json:"id"`        //主键
	ParentId  int    `json:"parentId"`  //父id
	DeptPath  string `json:"deptPath"`  //部门路径
	Name      string `json:"name"`      //部门名
	Type      int8   `json:"type"`      //类型 1分公司 2部门
	Principal string `json:"principal"` //部门领导
	Phone     string `json:"phone"`     //手机号
	Email     string `json:"email"`     //邮箱
	Sort      int8   `json:"sort"`      //排序
	Status    int8   `json:"status"`    //状态 1正常 2关闭
	Remark    string `json:"remark"`    //备注
}
