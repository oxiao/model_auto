package model

import (
	"time"
)

type JxgcSp struct {
Model
 SpID string `gorm:"type:varchar(36);unique;unique_index;not null" json:"sp_id"` // 
 McID string `gorm:"type:varchar(36)" json:"mc_id"` // 
 Mc string `gorm:"type:varchar(50)" json:"mc"` // 
 Bg decimal.Decimal `gorm:"type:decimal(13,3)" json:"bg"` // 
 Mc1 string `gorm:"type:varchar(50)" json:"mc1"` // 
 Ktfs string `gorm:"type:varchar(50)" json:"ktfs"` // 
 Fwnx decimal.Decimal `gorm:"type:decimal(13,3)" json:"fwnx"` // 
 Spcl decimal.Decimal `gorm:"type:decimal(13,3)" json:"spcl"` // 
 Spncl decimal.Decimal `gorm:"type:decimal(13,3)" json:"spncl"` // 
}
