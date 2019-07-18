package model

import (
	"time"
)

type BmMi struct {
Model
 MisID string `gorm:"type:varchar(36);unique;unique_index;not null" json:"mis_id"` // 
 Name string `gorm:"type:varchar(64)" json:"name"` // 
 Config string `gorm:"type:text" json:"config"` // {logo,title,skin,parm...}
 Menu string `gorm:"type:text" json:"menu"` // 菜单
 Map string `gorm:"type:text" json:"map"` // 地图
 PlugIn string `gorm:"type:text" json:"plug_in"` // 插件
 Desktop string `gorm:"type:text" json:"desktop"` // 桌面
}
