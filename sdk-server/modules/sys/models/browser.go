package models

import (
	"time"

	"gorm.io/gorm"
)

// Browser
type Browser struct {
	Id        int            `json:"id" gorm:"type:int unsigned;primaryKey;autoIncrement;comment:主键"` //主键
	Name      string         `json:"name" gorm:"type:varchar(255);comment:名称"`                        //名称
	EnvId     uint64         `json:"envId" gorm:"type:bigint;comment:环境ID"`                           //环境ID
	UserId    int            `json:"userId" gorm:"type:int ;comment:用户ID"`                            //用户ID
	Data      string         `json:"data" gorm:"type:text;comment:数据"`                                //数据
	CreatedAt time.Time      `json:"createdAt" gorm:"type:datetime(3);comment:创建时间"`                  //创建时间
	UpdatedAt time.Time      `json:"updatedAt" gorm:"type:datetime(3);comment:最后更新时间"`                //最后更新时间
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index;comment:删除时间"`                                     //删除时间
	Status    int8           `json:"status" gorm:"type:tinyint(1);comment:状态"`                        //状态 1 停止 3启动
}

func (Browser) TableName() string {
	return "browser"
}

func NewBrowser() *Browser {
	return &Browser{}
}
