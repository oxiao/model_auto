package model

import (
	"time"
)

type WsccZksj struct {
Model
 ZkID string `gorm:"type:varchar(36);unique;unique_index;not null" json:"zk_id"` // 
 Zkmc string `gorm:"type:varchar(50)" json:"zkmc"` // 
 Dd string `gorm:"type:varchar(50)" json:"dd"` // 
 Bzdjl decimal.Decimal `gorm:"type:decimal(13,3)" json:"bzdjl"` // 
 Qj decimal.Decimal `gorm:"type:decimal(13,3)" json:"qj"` // 
 Fwj string `gorm:"type:varchar(50)" json:"fwj"` // 
 Jdbgd decimal.Decimal `gorm:"type:decimal(13,3)" json:"jdbgd"` // 
 Jzxjl decimal.Decimal `gorm:"type:decimal(13,3)" json:"jzxjl"` // 
 Zb string `gorm:"type:varchar(50)" json:"zb"` // 
 Yb string `gorm:"type:varchar(50)" json:"yb"` // 
 Ks string `gorm:"type:varchar(50)" json:"ks"` // 
 Jmsd decimal.Decimal `gorm:"type:decimal(13,3)" json:"jmsd"` // 
 Jysd decimal.Decimal `gorm:"type:decimal(13,3)" json:"jysd"` // 
 Fkcd decimal.Decimal `gorm:"type:decimal(18,3)" json:"fkcd"` // 
 Fkcl string `gorm:"type:varchar(50)" json:"fkcl"` // 
}
