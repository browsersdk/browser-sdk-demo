package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// 用户
type Admin struct {
	Id         int       `json:"id" gorm:"type:int unsigned;primaryKey;autoIncrement;comment:主键"` //主键
	Username   string    `json:"username" gorm:"type:varchar(32);comment:用户名"`                    //用户名
	Phone      string    `json:"phone" gorm:"type:varchar(11);comment:手机号"`                       //手机号
	Email      string    `json:"email" gorm:"type:varchar(128);comment:邮箱"`                       //邮箱
	Password   string    `json:"password" gorm:"type:varchar(128);comment:密码"`                    //密码
	Nickname   string    `json:"nickname" gorm:"type:varchar(128);comment:昵称"`                    //昵称
	Name       string    `json:"name" gorm:"type:varchar(64);comment:姓名"`                         //姓名
	Avatar     string    `json:"avatar" gorm:"type:varchar(255);comment:头像"`                      //头像
	Bio        string    `json:"bio" gorm:"type:varchar(255);comment:签名"`                         //签名
	Birthday   time.Time `json:"birthday" gorm:"type:date;default:(-);comment:生日 格式 yyyy-MM-dd"`  //生日 格式 yyyy-MM-dd
	Gender     int8      `json:"gender" gorm:"type:tinyint(1);comment:性别 1男 2女 3未知"`              //性别 1男 2女 3未知
	RoleId     int       `json:"roleId" gorm:"type:int;comment:角色id"`                             //角色id
	DeptId     int       `json:"deptId" gorm:"type:int;comment:部门id"`                             //部门id
	DeptPath   string    `json:"deptPath" gorm:"type:varchar(255);comment:部门路径"`                  //部门路径
	Remark     string    `json:"remark" gorm:"type:varchar(255);comment:备注"`                      //备注
	LockTime   time.Time `json:"lockTime" gorm:"type:datetime(3);default:(-);comment:锁定结束时间"`     //锁定结束时间
	Status     int8      `json:"status" gorm:"type:tinyint;comment:状态 1正常 "`                      //状态 1正常
	ClientId   string    `json:"clientId" gorm:"type:varchar(64);comment:客户端id"`                  //客户端id
	RegIp      string    `json:"regIp" gorm:"type:varchar(42);comment:注册ip"`                      //注册ip
	IpLocation string    `json:"ipLocation" gorm:"type:varchar(100);comment:IpLocation"`          //
	UpdateBy   int       `json:"updateBy" gorm:"type:int unsigned;comment:更新者"`                   //更新者
	CreatedAt  time.Time `json:"createdAt" gorm:"type:datetime(4);comment:创建时间"`                  //创建时间
	UpdatedAt  time.Time `json:"updatedAt" gorm:"type:datetime(4);comment:最后更新时间"`                //最后更新时间
}

const TBAdmin = "admin"

func (Admin) TableName() string {
	return TBAdmin
}

// 加密
func (e *Admin) Encrypt() (err error) {
	if e.Password == "" {
		return
	}

	var hash []byte
	if hash, err = bcrypt.GenerateFromPassword([]byte(e.Password), bcrypt.DefaultCost); err != nil {
		return
	} else {
		e.Password = string(hash)
		return
	}
}

func (e *Admin) BeforeCreate(_ *gorm.DB) error {
	return e.Encrypt()
}

func (e *Admin) GenPwd(pwd string) (enPwd string, err error) {
	var hash []byte
	if hash, err = bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost); err != nil {
		return
	} else {
		enPwd = string(hash)
	}
	return
}

func (e *Admin) CompPwd(srcPwd string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(e.Password), []byte(srcPwd)); err != nil {
		return false
	}
	return true
}
