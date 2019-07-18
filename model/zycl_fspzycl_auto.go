package model

import (
	"time"
)

type ZyclFspzycl struct {
Model
 FspzyclID string `gorm:"type:varchar(36);unique;unique_index;not null" json:"fspzycl_id"` // 
 Sp string `gorm:"type:varchar(50)" json:"sp"` // 
 Ssmc string `gorm:"type:varchar(36)" json:"ssmc"` // 
 Zycllx string `gorm:"type:varchar(50)" json:"zycllx"` // 
 Zycl decimal.Decimal `gorm:"type:decimal(13,3)" json:"zycl"` // 
 Kccl decimal.Decimal `gorm:"type:decimal(13,3)" json:"kccl"` // 
}
