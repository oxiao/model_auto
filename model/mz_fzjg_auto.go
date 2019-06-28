package model

import "github.com/shopspring/decimal"

type MzFzjg struct {
	MdfzjgID string          `sql:"index" gorm:"type:varchar(36);primary_key" json:"mdfzjg_id"` //
	MzID     string          `gorm:"type:varchar(36)" json:"mz_id"`                             //
	Fxh      decimal.Decimal `gorm:"type:decimal(13,3)" json:"fxh"`                             //
	Fxd      decimal.Decimal `gorm:"type:decimal(13,3)" json:"fxd"`                             //
	Tyzw     string          `gorm:"type:varchar(50)" json:"tyzw"`                              //
	Tyzwjj   decimal.Decimal `gorm:"type:decimal(13,3)" json:"tyzwjj"`                          //
	Syddm    string          `gorm:"type:varchar(50)" json:"syddm"`                             //
}
