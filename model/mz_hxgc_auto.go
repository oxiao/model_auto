package model

import "github.com/shopspring/decimal"

type MzHxgc struct {
	MdhxgcID string          `sql:"index" gorm:"type:varchar(36);primary_key" json:"mdhxgc_id"` //
	MzID     string          `gorm:"type:varchar(36)" json:"mz_id"`                             //
	Yjz      decimal.Decimal `gorm:"type:decimal(13,3)" json:"yjz"`                             //
	Wjz      decimal.Decimal `gorm:"type:decimal(13,3)" json:"wjz"`                             //
	Kwz      decimal.Decimal `gorm:"type:decimal(13,3)" json:"kwz"`                             //
	Sf       decimal.Decimal `gorm:"type:decimal(13,3)" json:"sf"`                              //
	Hf       decimal.Decimal `gorm:"type:decimal(13,3)" json:"hf"`                              //
	Hff      decimal.Decimal `gorm:"type:decimal(13,3)" json:"hff"`                             //
}
