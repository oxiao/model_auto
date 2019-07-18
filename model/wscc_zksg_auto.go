package model

import (
	"time"
)

type WsccZksg struct {
Model
 ZksgID string `gorm:"type:varchar(36);unique;unique_index;not null" json:"zksg_id"` // 
 ZkID string `gorm:"type:varchar(50)" json:"zk_id"` // 
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
 Dlxx string `gorm:"type:varchar(50)" json:"dlxx"` // 
 DlxxSd decimal.Decimal `gorm:"type:decimal(13,3)" json:"dlxx_sd"` // 
 Sgsj time.Time `gorm:"type:datetime" json:"sgsj"` // 
 Bc string `gorm:"type:varchar(50)" json:"bc"` // 
 Szry string `gorm:"type:varchar(50)" json:"szry"` // 
 Fkcd decimal.Decimal `gorm:"type:decimal(18,3)" json:"fkcd"` // 
 Fkcl string `gorm:"type:varchar(50)" json:"fkcl"` // 
 Fksj time.Time `gorm:"type:datetime" json:"fksj"` // 
}
