package model

import (
	"time"
)

type WssgBz struct {
Model
 WsbzxxID string `gorm:"type:varchar(36);unique;unique_index;not null" json:"wsbzxx_id"` // 
 Wsjjzb decimal.Decimal `gorm:"type:decimal(13,3)" json:"wsjjzb"` // 
 Yqndjc string `gorm:"type:varchar(50)" json:"yqndjc"` // 
 Tfztjc string `gorm:"type:varchar(50)" json:"tfztjc"` // 
 Hypjzb decimal.Decimal `gorm:"type:decimal(13,3)" json:"hypjzb"` // 
 Mhzb string `gorm:"type:varchar(50)" json:"mhzb"` // 
 Sbyhzb decimal.Decimal `gorm:"type:decimal(13,3)" json:"sbyhzb"` // 
 Wsbzsg string `gorm:"type:varchar(50)" json:"wsbzsg"` // 
}
