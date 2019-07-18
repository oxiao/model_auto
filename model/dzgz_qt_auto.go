package model

import (
	"time"
)

type DzgzQt struct {
Model
 McID string `gorm:"type:varchar(36)" json:"mc_id"` // 
 QtID string `gorm:"type:varchar(36);unique;unique_index;not null" json:"qt_id"` // 
 Mc string `gorm:"type:varchar(50)" json:"mc"` // 
 Lx string `gorm:"type:varchar(36)" json:"lx"` // 
 Xz string `gorm:"type:varchar(36)" json:"xz"` // 类型：岩浆岩，断裂，火成岩侵入，降落柱，节理，褶曲
 Fgfl string `gorm:"type:varchar(50)" json:"fgfl"` // 
 Ly string `gorm:"type:varchar(36)" json:"ly"` // 
 Ms string `gorm:"type:varchar(200)" json:"ms"` // 
}
