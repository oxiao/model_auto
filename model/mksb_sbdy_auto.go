package model

import (
	"time"
)

type MksbSbdy struct {
Model
 SbtzID string `gorm:"type:varchar(36);unique;unique_index;not null" json:"sbtz_id"` // 
 Sbfl string `gorm:"type:varchar(50)" json:"sbfl"` // 
 Sblx string `gorm:"type:varchar(50)" json:"sblx"` // 
 Sbmc string `gorm:"type:varchar(50)" json:"sbmc"` // 
 Sbbsm string `gorm:"type:varchar(50)" json:"sbbsm"` // 
 Ggxh string `gorm:"type:varchar(50)" json:"ggxh"` // 
 Edrl decimal.Decimal `gorm:"type:decimal(13,3)" json:"edrl"` // 
 Jxfs string `gorm:"type:varchar(50)" json:"jxfs"` // 
 Eddy decimal.Decimal `gorm:"type:decimal(13,3)" json:"eddy"` // 
 Eddl decimal.Decimal `gorm:"type:decimal(13,3)" json:"eddl"` // 
 Sbzl string `gorm:"type:varchar(50)" json:"sbzl"` // 
 Fbbz string `gorm:"type:varchar(50)" json:"fbbz"` // 
 Mabz string `gorm:"type:varchar(50)" json:"mabz"` // 
 Sbyt string `gorm:"type:varchar(50)" json:"sbyt"` // 
 Gzrq time.Time `gorm:"type:datetime" json:"gzrq"` // 
 Sbsl decimal.Decimal `gorm:"type:decimal(13,3)" json:"sbsl"` // 
 Dw string `gorm:"type:varchar(50)" json:"dw"` // 
 Azdd string `gorm:"type:varchar(50)" json:"azdd"` // 
 Azry string `gorm:"type:varchar(50)" json:"azry"` // 
 Azsj time.Time `gorm:"type:datetime" json:"azsj"` // 
 Scrq time.Time `gorm:"type:datetime" json:"scrq"` // 
 CcID string `gorm:"type:varchar(50)" json:"cc_id"` // 
 Sccj string `gorm:"type:varchar(50)" json:"sccj"` // 
 Jscs string `gorm:"type:varchar(50)" json:"jscs"` // 
 Glry string `gorm:"type:varchar(200)" json:"glry"` // 管理人员
 Glbm string `gorm:"type:varchar(200)" json:"glbm"` // 管理部門
}
