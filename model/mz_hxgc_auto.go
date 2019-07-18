package model

import (
	"time"
)

type MzHxgc struct {
Model
 MdhxgcID string `gorm:"type:varchar(36);unique;unique_index;not null" json:"mdhxgc_id"` // 
 MzID string `gorm:"type:varchar(36)" json:"mz_id"` // 
 Yjz decimal.Decimal `gorm:"type:decimal(13,3)" json:"yjz"` // 
 Wjz decimal.Decimal `gorm:"type:decimal(13,3)" json:"wjz"` // 
 Kwz decimal.Decimal `gorm:"type:decimal(13,3)" json:"kwz"` // 
 Sf decimal.Decimal `gorm:"type:decimal(13,3)" json:"sf"` // 
 Hf decimal.Decimal `gorm:"type:decimal(13,3)" json:"hf"` // 
 Hff decimal.Decimal `gorm:"type:decimal(13,3)" json:"hff"` // 
}
