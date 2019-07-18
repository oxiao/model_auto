package model

import (
	"time"
)

type WsfcYswshl struct {
Model
 YswshlscID string `gorm:"type:varchar(36);unique;unique_index;not null" json:"yswshlsc_id"` // 
 McID string `gorm:"type:varchar(36)" json:"mc_id"` // 
 Dd string `gorm:"type:varchar(50)" json:"dd"` // 
 Bg decimal.Decimal `gorm:"type:decimal(13,3)" json:"bg"` // 
 Mc string `gorm:"type:varchar(50)" json:"mc"` // 
 Ms decimal.Decimal `gorm:"type:decimal(13,3)" json:"ms"` // 
 Hl decimal.Decimal `gorm:"type:decimal(13,3)" json:"hl"` // 
 Yl decimal.Decimal `gorm:"type:decimal(13,3)" json:"yl"` // 
 Cdr string `gorm:"type:varchar(50)" json:"cdr"` // 
 Tcjd string `gorm:"type:varchar(50)" json:"tcjd"` // 
 Zb decimal.Decimal `gorm:"type:decimal(13,3)" json:"zb"` // 
 Sj time.Time `gorm:"type:datetime" json:"sj"` // 
 Cddw string `gorm:"type:varchar(50)" json:"cddw"` // 
}
