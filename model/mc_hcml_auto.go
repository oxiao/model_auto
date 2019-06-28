package model

import "github.com/shopspring/decimal"

type McHcml struct {
	CqID    string          `gorm:"type:varchar(36)" json:"cq_id"`     //
	McID    string          `gorm:"type:varchar(36)" json:"mc_id"`     //
	CmgzmID string          `gorm:"type:varchar(36)" json:"cmgzm_id"`  //
	ZxKccd  decimal.Decimal `gorm:"type:decimal(18,3)" json:"zx_kccd"` //
	QxKccd  decimal.Decimal `gorm:"type:decimal(18,3)" json:"qx_kccd"` //
	Cg      decimal.Decimal `gorm:"type:decimal(18,3)" json:"cg"`      //
	Hcl     decimal.Decimal `gorm:"type:decimal(18,3)" json:"hcl"`     //
	Smd     decimal.Decimal `gorm:"type:decimal(18,3)" json:"smd"`     //
}
