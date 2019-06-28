package model

import "github.com/shopspring/decimal"

type JxgcBhmzMzxt struct {
	MzxtID string          `sql:"index" gorm:"type:varchar(36);primary_key" json:"mzxt_id"` //
	BhmzID string          `gorm:"type:varchar(36)" json:"bhmz_id"`                         //
	Dh     string          `gorm:"type:varchar(50)" json:"dh"`                              //
	Zbx    decimal.Decimal `gorm:"type:decimal(13,3)" json:"zbx"`                           //
	Zby    decimal.Decimal `gorm:"type:decimal(13,3)" json:"zby"`                           //
}
