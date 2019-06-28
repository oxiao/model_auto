package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type McMcbzxsy struct {
	McbzxsyID string          `sql:"index" gorm:"type:varchar(36);primary_key" json:"mcbzxsy_id"` //
	McID      string          `gorm:"type:varchar(36)" json:"mc_id"`                              //
	Mc        string          `gorm:"type:varchar(50)" json:"mc"`                                 //
	Hycd      decimal.Decimal `gorm:"type:decimal(13,3)" json:"hycd"`                             //
	Hyl       decimal.Decimal `gorm:"type:decimal(13,3)" json:"hyl"`                              //
	Yfyl      decimal.Decimal `gorm:"type:decimal(13,3)" json:"yfyl"`                             //
	Krw       string          `gorm:"type:varchar(50)" json:"krw"`                                //
	Brw       string          `gorm:"type:varchar(50)" json:"brw"`                                //
	Ywbzwx    string          `gorm:"type:varchar(50)" json:"ywbzwx"`                             //
	Sysj      time.Time       `gorm:"type:datetime" json:"sysj"`                                  //
	Sydw      string          `gorm:"type:varchar(50)" json:"sydw"`                               //
}
