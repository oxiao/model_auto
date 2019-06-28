package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type WsccHjjc struct {
	HjjcsjID string          `sql:"index" gorm:"type:varchar(36);primary_key" json:"hjjcsj_id"` //
	CfbID    string          `gorm:"type:varchar(50)" json:"cfb_id"`                            //
	Jcsj     time.Time       `gorm:"type:datetime" json:"jcsj"`                                 //
	Nd       decimal.Decimal `gorm:"type:decimal(13,3)" json:"nd"`                              //
	Wd       decimal.Decimal `gorm:"type:decimal(13,3)" json:"wd"`                              //
	Yl       decimal.Decimal `gorm:"type:decimal(13,3)" json:"yl"`                              //
	Hhll     decimal.Decimal `gorm:"type:decimal(13,3)" json:"hhll"`                            //
	Cll      decimal.Decimal `gorm:"type:decimal(13,3)" json:"cll"`                             //
}
