package model

import (
	"time"
)

type DzgzDc struct {
Model
 DcID string `gorm:"type:varchar(36);unique;unique_index;not null" json:"dc_id"` // 
 McID string `gorm:"type:varchar(36)" json:"mc_id"` // 
 Fgfl string `gorm:"type:varchar(50)" json:"fgfl"` // 
 Mc string `gorm:"type:varchar(50)" json:"mc"` // 
 Lc decimal.Decimal `gorm:"type:decimal(13,3)" json:"lc"` // 
 Zx string `gorm:"type:varchar(50)" json:"zx"` // 
 Cd decimal.Decimal `gorm:"type:decimal(13,3)" json:"cd"` // 
 Qj string `gorm:"type:varchar(50)" json:"qj"` // 
 Dcxz string `gorm:"type:varchar(50)" json:"dcxz"` // 
 Lxxz string `gorm:"type:varchar(50)" json:"lxxz"` // 
 Ly string `gorm:"type:varchar(50)" json:"ly"` // 
}
