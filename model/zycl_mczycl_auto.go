package model

import "github.com/shopspring/decimal"

type ZyclMczycl struct {
	MczyclID string          `sql:"index" gorm:"type:varchar(36);primary_key" json:"mczycl_id"` //
	McID     string          `gorm:"type:varchar(36)" json:"mc_id"`                             //
	Mcmc     string          `gorm:"type:varchar(50)" json:"mcmc"`                              //
	Zycllx   string          `gorm:"type:varchar(50)" json:"zycllx"`                            //
	Zycl     decimal.Decimal `gorm:"type:decimal(13,3)" json:"zycl"`                            //
	Kccl     decimal.Decimal `gorm:"type:decimal(13,3)" json:"kccl"`                            //
}
