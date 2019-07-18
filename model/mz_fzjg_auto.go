package model

import (
	"time"
)

type MzFzjg struct {
Model
 MdfzjgID string `gorm:"type:varchar(36);unique;unique_index;not null" json:"mdfzjg_id"` // 
 MzID string `gorm:"type:varchar(36)" json:"mz_id"` // 
 Fxh decimal.Decimal `gorm:"type:decimal(13,3)" json:"fxh"` // 
 Fxd decimal.Decimal `gorm:"type:decimal(13,3)" json:"fxd"` // 
 Tyzw string `gorm:"type:varchar(50)" json:"tyzw"` // 
 Tyzwjj decimal.Decimal `gorm:"type:decimal(13,3)" json:"tyzwjj"` // 
 Syddm string `gorm:"type:varchar(50)" json:"syddm"` // 
}
