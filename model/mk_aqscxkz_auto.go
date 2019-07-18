package model

import (
	"time"
)

type MkAqscxkz struct {
Model
 AqscxkzID string `gorm:"type:varchar(36);unique;unique_index;not null" json:"aqscxkz_id"` // 
 MkjbxxID string `gorm:"type:varchar(36)" json:"mkjbxx_id"` // 
 Zh string `gorm:"type:varchar(50)" json:"zh"` // 
 Yxq decimal.Decimal `gorm:"type:decimal(13,3)" json:"yxq"` // 
 Pzscnl string `gorm:"type:varchar(50)" json:"pzscnl"` // 
 Pzkcmc string `gorm:"type:varchar(50)" json:"pzkcmc"` // 
 Bzjg string `gorm:"type:varchar(50)" json:"bzjg"` // 
 Bzsj time.Time `gorm:"type:datetime" json:"bzsj"` // 
 Zt string `gorm:"type:varchar(50)" json:"zt"` // 
}
