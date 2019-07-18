package model

import (
	"time"
)

type JxgcBhmzMzxt struct {
Model
 MzxtID string `gorm:"type:varchar(36);unique;unique_index;not null" json:"mzxt_id"` // 
 BhmzID string `gorm:"type:varchar(36)" json:"bhmz_id"` // 
 Dh string `gorm:"type:varchar(50)" json:"dh"` // 
 Zbx decimal.Decimal `gorm:"type:decimal(13,3)" json:"zbx"` // 
 Zby decimal.Decimal `gorm:"type:decimal(13,3)" json:"zby"` // 
}
