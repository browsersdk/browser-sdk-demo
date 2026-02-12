package dto

import (
	"dilu/modules/sys/models"
	"time"

	"github.com/baowk/dilu-core/core/base"
)

type AdminGetPageReq struct {
	base.ReqPage `query:"-"`
	SortOrder    string `json:"-" query:"type:order;column:id"`
	Status       int8   `json:"status" query:"column:status"`  //状态 1正常
	DeptId       int    `json:"deptId" query:"column:dept_id"` //部门id
	Name         string `json:"name" query:"column:name"`      // 姓名
	Phone        string `json:"phone" query:"column:phone"`    //手机号
}

func (AdminGetPageReq) TableName() string {
	return models.TBAdmin
}

// 用户
type AdminDto struct {
	Id         int       `json:"id"`         //主键
	Username   string    `json:"username"`   //用户名
	Phone      string    `json:"phone"`      //手机号
	Email      string    `json:"email"`      //邮箱
	Password   string    `json:"password"`   //密码
	Nickname   string    `json:"nickname"`   //昵称
	Name       string    `json:"name"`       //姓名
	Avatar     string    `json:"avatar"`     //头像
	Bio        string    `json:"bio"`        //签名
	Birthday   time.Time `json:"birthday"`   //生日 格式 yyyy-MM-dd
	Gender     int8      `json:"gender"`     //性别 1男 2女 3未知
	RoleId     int       `json:"roleId"`     //角色id
	DeptId     int       `json:"deptId"`     //部门id
	Remark     string    `json:"remark"`     //备注
	LockTime   time.Time `json:"lockTime"`   //锁定结束时间
	Status     int8      `json:"status"`     //状态 1正常
	ClientId   string    `json:"clientId"`   //客户端id
	RegIp      string    `json:"regIp"`      //注册ip
	IpLocation string    `json:"ipLocation"` //
}
