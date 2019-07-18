package model

import (
	"time"
)

type JxgcJd struct {
Model
 JdID string `gorm:"type:varchar(36);unique;unique_index;not null" json:"jd_id"` // 
 McID string `gorm:"type:varchar(36)" json:"mc_id"` // 
 SpID string `gorm:"type:varchar(36)" json:"sp_id"` // 
 Mc string `gorm:"type:varchar(50)" json:"mc"` // 
 Sp decimal.Decimal `gorm:"type:decimal(13,3)" json:"sp"` // 
 Xc decimal.Decimal `gorm:"type:decimal(13,3)" json:"xc"` // 
 ZxCd decimal.Decimal `gorm:"type:decimal(13,3)" json:"zx_cd"` // 
}
