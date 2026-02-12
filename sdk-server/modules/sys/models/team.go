package models

import (
	"dilu/common/config"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/baowk/dilu-core/common/utils/cryptos"
)

// 团队
type SysTeam struct {
	Id              int       `json:"id" gorm:"type:int;primaryKey;autoIncrement;comment:主键"`  // 主键
	Name            string    `json:"name" gorm:"type:varchar(32);comment:团队名"`                // 团队名
	Owner           int       `json:"owner" gorm:"type:int unsigned;comment:团队拥有者"`            // 团队拥有者
	Status          int       `json:"status" gorm:"type:tinyint;comment:状态"`                   // 状态 2 正常 -1 关闭
	AppKeyCreatedAt time.Time `json:"appKeyCreatedAt" gorm:"type:datetime;comment:appKey创建时间"` // appKey创建时间
	UpdateBy        int       `json:"updateBy" gorm:"type:int;comment:更新人"`                    // 更新人
	CreatedAt       time.Time `json:"createdAt" gorm:"type:datetime;comment:创建时间"`             // 创建时间
	UpdatedAt       time.Time `json:"updatedAt" gorm:"type:datetime;comment:更新时间"`             // 更新时间
}

func (SysTeam) TableName() string {
	return "team"
}

func NewSysTeam() *SysTeam {
	return &SysTeam{}
}

func (e *SysTeam) GetAppKey() string {
	// if e.AppKeyCreatedAt.IsZero() {
	// 	return ""
	// }
	en, err := cryptos.AesEncryptCBC([]byte(fmt.Sprintf("%d-%d", e.Id, e.AppKeyCreatedAt.Unix())), []byte(config.Ext.AesKey))
	if err != nil {
		return ""
	}
	return hex.EncodeToString(en)
}
