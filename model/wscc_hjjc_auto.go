package model

import (
	"time"
)

type WsccHjjc struct {
Model
 HjjcsjID string `gorm:"type:varchar(36);unique;unique_index;not null" json:"hjjcsj_id"` // 
 CfbID string `gorm:"type:varchar(50)" json:"cfb_id"` // 
 Jcsj time.Time `gorm:"type:datetime" json:"jcsj"` // 
 Nd decimal.Decimal `gorm:"type:decimal(13,3)" json:"nd"` // 
 Wd decimal.Decimal `gorm:"type:decimal(13,3)" json:"wd"` // 
 Yl decimal.Decimal `gorm:"type:decimal(13,3)" json:"yl"` // 
 Hhll decimal.Decimal `gorm:"type:decimal(13,3)" json:"hhll"` // 
 Cll decimal.Decimal `gorm:"type:decimal(13,3)" json:"cll"` // 
}
