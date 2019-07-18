package model

import (
	"time"
)

type ZyclMzzycl struct {
Model
 MzzyclID string `gorm:"type:varchar(36);unique;unique_index;not null" json:"mzzycl_id"` // 
 Ssmc string `gorm:"type:varchar(36)" json:"ssmc"` // 
 Mj decimal.Decimal `gorm:"type:decimal(13,3)" json:"mj"` // 
 Bmd decimal.Decimal `gorm:"type:decimal(13,3)" json:"bmd"` // 
 Gygc decimal.Decimal `gorm:"type:decimal(13,3)" json:"gygc"` // 
 Bjmz decimal.Decimal `gorm:"type:decimal(13,3)" json:"bjmz"` // 
 Dcmz decimal.Decimal `gorm:"type:decimal(13,3)" json:"dcmz"` // 
 Czmz decimal.Decimal `gorm:"type:decimal(13,3)" json:"czmz"` // 
 Qt decimal.Decimal `gorm:"type:decimal(13,3)" json:"qt"` // 
}
