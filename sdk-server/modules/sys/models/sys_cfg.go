package models

import (
	"time"
)

// 配置
type SysCfg struct {
	Id        int       `json:"id" gorm:"type:int unsigned;primaryKey;autoIncrement;comment:主键"` //主键
	TeamId    int       `json:"teamId" gorm:"type:int unsigned;comment:团队id"`                    // 团队id -1为平台
	Type      string    `json:"type" gorm:"type:varchar(64);comment:Type"`                       //类型
	Key       string    `json:"key" gorm:"type:varchar(128);comment:key"`                        //类型下key
	Val       string    `json:"val" gorm:"type:text;comment:值"`                                  //值
	Status    int       `json:"status" gorm:"type:tinyint;comment:Status"`                       //Status
	Remark    string    `json:"remark" gorm:"type:varchar(128);comment:Remark"`                  //Remark
	UpdateBy  int       `json:"updateBy" gorm:"type:int unsigned;comment:更新者"`                   //更新者
	UpdatedAt time.Time `json:"updatedAt" gorm:"type:datetime(3);comment:最后更新时间"`                //最后更新时间
	//Name      string    `json:"name" gorm:"type:varchar(128);comment:名字"`                        //名字
}

const TBSysCfg = "sys_cfg"

func (SysCfg) TableName() string {
	return TBSysCfg
}

func NewSysCfg() *SysCfg {
	return &SysCfg{}
}
