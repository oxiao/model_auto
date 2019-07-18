package model

import (
	"time"
)

type McGyfx struct {
Model
 GyfxID string `gorm:"type:varchar(36);unique;unique_index;not null" json:"gyfx_id"` // 
 MyID string `gorm:"type:varchar(36)" json:"my_id"` // 
 Sffx string `gorm:"type:varchar(50)" json:"sffx"` // 
 Hffx string `gorm:"type:varchar(50)" json:"hffx"` // 
 Hfffx string `gorm:"type:varchar(50)" json:"hfffx"` // 
 Gdthl decimal.Decimal `gorm:"type:decimal(13,3)" json:"gdthl"` // 
 Tsyeyht decimal.Decimal `gorm:"type:decimal(13,3)" json:"tsyeyht"` // 
 Mdfrl decimal.Decimal `gorm:"type:decimal(13,3)" json:"mdfrl"` // 
 Mzkwz string `gorm:"type:varchar(50)" json:"mzkwz"` // 
 Jxmdhsl decimal.Decimal `gorm:"type:decimal(13,3)" json:"jxmdhsl"` // 
 Jxmd decimal.Decimal `gorm:"type:decimal(13,3)" json:"jxmd"` // 
}
