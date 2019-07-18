package model

import (
	"time"
)

type WsfcWsfhd struct {
Model
 WsfhdID string `gorm:"type:varchar(36);unique;unique_index;not null" json:"wsfhd_id"` // 
 CqID string `gorm:"type:varchar(36)" json:"cq_id"` // 
 McID string `gorm:"type:varchar(36)" json:"mc_id"` // 
 Xjbg decimal.Decimal `gorm:"type:decimal(13,3)" json:"xjbg"` // 
 Ms decimal.Decimal `gorm:"type:decimal(13,3)" json:"ms"` // 
 Wshl decimal.Decimal `gorm:"type:decimal(13,3)" json:"wshl"` // 
 Wsyl decimal.Decimal `gorm:"type:decimal(13,3)" json:"wsyl"` // 
}
