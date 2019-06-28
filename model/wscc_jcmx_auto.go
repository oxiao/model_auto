package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type WsccJcmx struct {
	CcjcjlID string          `sql:"index" gorm:"type:varchar(36);primary_key" json:"ccjcjl_id"` //
	WsccxtID string          `gorm:"type:varchar(36)" json:"wsccxt_id"`                         //
	CfbID    string          `gorm:"type:varchar(50)" json:"cfb_id"`                            //
	Cfsj     decimal.Decimal `gorm:"type:decimal(13,3)" json:"cfsj"`                            //
	Dqy      decimal.Decimal `gorm:"type:decimal(13,3)" json:"dqy"`                             //
	Fy       decimal.Decimal `gorm:"type:decimal(13,3)" json:"fy"`                              //
	Yc       decimal.Decimal `gorm:"type:decimal(13,3)" json:"yc"`                              //
	Nd       decimal.Decimal `gorm:"type:decimal(13,3)" json:"nd"`                              //
	Gnwd     decimal.Decimal `gorm:"type:decimal(13,3)" json:"gnwd"`                            //
	Hhl      decimal.Decimal `gorm:"type:decimal(13,3)" json:"hhl"`                             //
	Cl       decimal.Decimal `gorm:"type:decimal(13,3)" json:"cl"`                              //
	Jcsj     time.Time       `gorm:"type:datetime" json:"jcsj"`                                 //
	Jcry     string          `gorm:"type:varchar(50)" json:"jcry"`                              //
	Ccdd     string          `gorm:"type:varchar(50)" json:"ccdd"`                              //
}
