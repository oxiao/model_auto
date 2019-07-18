package model

import (
	"time"
)

type MkScxkz struct {
Model
 ScxkzID string `gorm:"type:varchar(36);unique;unique_index;not null" json:"scxkz_id"` // 
 MkjbxxID string `gorm:"type:varchar(36)" json:"mkjbxx_id"` // 
 Zh string `gorm:"type:varchar(50)" json:"zh"` // 
 Yxq string `gorm:"type:varchar(50)" json:"yxq"` // 
 Cspfwh string `gorm:"type:varchar(50)" json:"cspfwh"` // 
 Ztgcyswh string `gorm:"type:varchar(50)" json:"ztgcyswh"` // 
 Xkscnl string `gorm:"type:varchar(50)" json:"xkscnl"` // 
 Xkkcmc string `gorm:"type:varchar(50)" json:"xkkcmc"` // 
 Bzjg string `gorm:"type:varchar(50)" json:"bzjg"` // 
 Bzsj time.Time `gorm:"type:datetime" json:"bzsj"` // 
}
