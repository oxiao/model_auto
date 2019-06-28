package model

import "github.com/shopspring/decimal"

type ZyclFspzycl struct {
	FspzyclID string          `sql:"index" gorm:"type:varchar(36);primary_key" json:"fspzycl_id"` //
	Sp        string          `gorm:"type:varchar(50)" json:"sp"`                                 //
	Ssmc      string          `gorm:"type:varchar(36)" json:"ssmc"`                               //
	Zycllx    string          `gorm:"type:varchar(50)" json:"zycllx"`                             //
	Zycl      decimal.Decimal `gorm:"type:decimal(13,3)" json:"zycl"`                             //
	Kccl      decimal.Decimal `gorm:"type:decimal(13,3)" json:"kccl"`                             //
}
