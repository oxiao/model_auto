package model

import (
	"time"
)

type JxgcBhmz struct {
Model
 BhmzID string `gorm:"type:varchar(36);unique;unique_index;not null" json:"bhmz_id"` // 
 McID string `gorm:"type:varchar(36)" json:"mc_id"` // 
 MzID string `gorm:"type:varchar(50)" json:"mz_id"` // 
 Mzxz string `gorm:"type:varchar(50)" json:"mzxz"` // 
 Cw string `gorm:"type:varchar(50)" json:"cw"` // 
 GzmMc string `gorm:"type:varchar(50)" json:"gzm_mc"` // 
 Zxc decimal.Decimal `gorm:"type:decimal(13,3)" json:"zxc"` // 
 Xc decimal.Decimal `gorm:"type:decimal(13,3)" json:"xc"` // 
 Hd decimal.Decimal `gorm:"type:decimal(13,3)" json:"hd"` // 
 Mj decimal.Decimal `gorm:"type:decimal(13,3)" json:"mj"` // 
 Qj decimal.Decimal `gorm:"type:decimal(13,3)" json:"qj"` // 
 Cl decimal.Decimal `gorm:"type:decimal(13,3)" json:"cl"` // 
 Spjg string `gorm:"type:varchar(50)" json:"spjg"` // 
 Pzl decimal.Decimal `gorm:"type:decimal(13,3)" json:"pzl"` // 
 Pzsj time.Time `gorm:"type:datetime" json:"pzsj"` // 
 Mzlssj time.Time `gorm:"type:datetime" json:"mzlssj"` // 
 Pzwh string `gorm:"type:varchar(50)" json:"pzwh"` // 
 Mztxl decimal.Decimal `gorm:"type:decimal(13,3)" json:"mztxl"` // 
 CcFcl decimal.Decimal `gorm:"type:decimal(13,3)" json:"cc_fcl"` // 
 Zb string `gorm:"type:varchar(50)" json:"zb"` // 
}
