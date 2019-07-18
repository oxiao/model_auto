package model

import (
	"time"
)

type ZyclMczycl struct {
Model
 MczyclID string `gorm:"type:varchar(36);unique;unique_index;not null" json:"mczycl_id"` // 
 McID string `gorm:"type:varchar(36)" json:"mc_id"` // 
 Mcmc string `gorm:"type:varchar(50)" json:"mcmc"` // 
 Zycllx string `gorm:"type:varchar(50)" json:"zycllx"` // 
 Zycl decimal.Decimal `gorm:"type:decimal(13,3)" json:"zycl"` // 
 Kccl decimal.Decimal `gorm:"type:decimal(13,3)" json:"kccl"` // 
}
