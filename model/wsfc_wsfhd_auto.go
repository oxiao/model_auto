package model

import "github.com/shopspring/decimal"

type WsfcWsfhd struct {
	WsfhdID string          `sql:"index" gorm:"type:varchar(36);primary_key" json:"wsfhd_id"` //
	CqID    string          `gorm:"type:varchar(36)" json:"cq_id"`                            //
	McID    string          `gorm:"type:varchar(36)" json:"mc_id"`                            //
	Xjbg    decimal.Decimal `gorm:"type:decimal(13,3)" json:"xjbg"`                           //
	Ms      decimal.Decimal `gorm:"type:decimal(13,3)" json:"ms"`                             //
	Wshl    decimal.Decimal `gorm:"type:decimal(13,3)" json:"wshl"`                           //
	Wsyl    decimal.Decimal `gorm:"type:decimal(13,3)" json:"wsyl"`                           //
}
