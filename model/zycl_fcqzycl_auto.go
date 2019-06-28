package model

import "github.com/shopspring/decimal"

type ZyclFcqzycl struct {
	FcqzyclID string          `sql:"index" gorm:"type:varchar(36);primary_key" json:"fcqzycl_id"` //
	Cqmc      string          `gorm:"type:varchar(50)" json:"cqmc"`                               //
	Ssmc      string          `gorm:"type:varchar(36)" json:"ssmc"`                               //
	Zycllx    string          `gorm:"type:varchar(50)" json:"zycllx"`                             //
	Zycl      decimal.Decimal `gorm:"type:decimal(13,3)" json:"zycl"`                             //
	Kccl      decimal.Decimal `gorm:"type:decimal(13,3)" json:"kccl"`                             //
}
