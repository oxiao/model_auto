package model

import (
	"time"
)

type BmEmployeeMenu struct {
Model
 EmployeeMenuID string `gorm:"type:varchar(36);unique;unique_index;not null" json:"employee_menu_id"` // 人员菜单编号
 EmployeeID string `gorm:"type:varchar(36)" json:"employee_id"` // 员工编号
 RoleID string `gorm:"type:varchar(36)" json:"role_id"` // 角色编号
 MisID string `gorm:"type:varchar(64)" json:"mis_id"` // 系统编号
 Menu string `gorm:"type:text" json:"menu"` // 菜單
 Map string `gorm:"type:text" json:"map"` // 地图
 PlugIn string `gorm:"type:text" json:"plug_in"` // 插件
 Desktop string `gorm:"type:text" json:"desktop"` // 桌面
}
