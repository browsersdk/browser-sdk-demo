package models

// SysMenuApiRule
type SysMenuApiRule struct {
	SysMenuId uint `json:"sysMenuId" gorm:"type:int unsigned;primaryKey;comment:菜单id"` //菜单id
	SysApiId  uint `json:"sysApiId"  gorm:"type:int unsigned;primaryKey;comment:接口id"` //接口id
}

const TBSysMenuApiRule = "sys_menu_api_rule"

func (SysMenuApiRule) TableName() string {
	return TBSysMenuApiRule
}

func NewSysMenuApiRule() *SysMenuApiRule {
	return &SysMenuApiRule{}
}
