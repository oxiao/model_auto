package model

import (
	"time"
)

type ZyclFcqzycl struct {
Model
 FcqzyclID string `gorm:"type:varchar(36);unique;unique_index;not null" json:"fcqzycl_id"` // 
 Cqmc string `gorm:"type:varchar(50)" json:"cqmc"` // 
 Ssmc string `gorm:"type:varchar(36)" json:"ssmc"` // 
 Zycllx string `gorm:"type:varchar(50)" json:"zycllx"` // 
 Zycl decimal.Decimal `gorm:"type:decimal(13,3)" json:"zycl"` // 
 Kccl decimal.Decimal `gorm:"type:decimal(13,3)" json:"kccl"` // 
}
