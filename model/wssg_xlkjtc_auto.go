package model

import "github.com/shopspring/decimal"

type WssgXlkjtc struct {
	XlkjtcxxID string          `sql:"index" gorm:"type:varchar(36);primary_key" json:"xlkjtcxx_id"` //
	Tcsg       string          `gorm:"type:varchar(50)" json:"tcsg"`                                //
	Dlxx       string          `gorm:"type:varchar(50)" json:"dlxx"`                                //
	Wsyl       decimal.Decimal `gorm:"type:decimal(13,3)" json:"wsyl"`                              //
	Wshl       decimal.Decimal `gorm:"type:decimal(13,3)" json:"wshl"`                              //
	Wsycl      decimal.Decimal `gorm:"type:decimal(13,3)" json:"wsycl"`                             //
}
