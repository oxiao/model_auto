package model

import (
	"time"
)

type BmRoleMenu struct {
Model
 RoleMenuID string `gorm:"type:varchar(36);unique;unique_index;not null" json:"role_menu_id"` // 角色菜单编号
 RoleID string `gorm:"type:varchar(36)" json:"role_id"` // 角色编号
 MisID string `gorm:"type:varchar(64)" json:"mis_id"` // 系统编号
 Menu string `gorm:"type:text" json:"menu"` // 菜单
 Map string `gorm:"type:text" json:"map"` // 地图
 PlugIn string `gorm:"type:text" json:"plug_in"` // 插件
 Desktop string `gorm:"type:text" json:"desktop"` // 桌面
}
