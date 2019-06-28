package model

import "github.com/shopspring/decimal"

type ZyclMzzycl struct {
	MzzyclID string          `sql:"index" gorm:"type:varchar(36);primary_key" json:"mzzycl_id"` //
	Ssmc     string          `gorm:"type:varchar(36)" json:"ssmc"`                              //
	Mj       decimal.Decimal `gorm:"type:decimal(13,3)" json:"mj"`                              //
	Bmd      decimal.Decimal `gorm:"type:decimal(13,3)" json:"bmd"`                             //
	Gygc     decimal.Decimal `gorm:"type:decimal(13,3)" json:"gygc"`                            //
	Bjmz     decimal.Decimal `gorm:"type:decimal(13,3)" json:"bjmz"`                            //
	Dcmz     decimal.Decimal `gorm:"type:decimal(13,3)" json:"dcmz"`                            //
	Czmz     decimal.Decimal `gorm:"type:decimal(13,3)" json:"czmz"`                            //
	Qt       decimal.Decimal `gorm:"type:decimal(13,3)" json:"qt"`                              //
}
