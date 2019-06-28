package model

import "github.com/shopspring/decimal"

type ZyclLnkcwczb struct {
	LnkcwczbID string          `sql:"index" gorm:"type:varchar(36);primary_key" json:"lnkcwczb_id"` //
	Nf         decimal.Decimal `gorm:"type:decimal(13,3)" json:"nf"`                                //
	Ymcl       decimal.Decimal `gorm:"type:decimal(13,3)" json:"ymcl"`                              //
	Fcm        decimal.Decimal `gorm:"type:decimal(13,3)" json:"fcm"`                               //
	Lwm        decimal.Decimal `gorm:"type:decimal(13,3)" json:"lwm"`                               //
	Zjc        decimal.Decimal `gorm:"type:decimal(13,3)" json:"zjc"`                               //
	Ktjc       decimal.Decimal `gorm:"type:decimal(13,3)" json:"ktjc"`                              //
}
