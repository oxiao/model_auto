package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type SwdzKjt struct {
	KjtsxxID string          `sql:"index" gorm:"type:varchar(36);primary_key" json:"kjtsxx_id"` //
	Tssj     time.Time       `gorm:"type:datetime" json:"tssj"`                                 //
	Tsdd     string          `gorm:"type:varchar(50)" json:"tsdd"`                              //
	Tslx     string          `gorm:"type:varchar(50)" json:"tslx"`                              //
	Tstd     string          `gorm:"type:varchar(50)" json:"tstd"`                              //
	Tssy     string          `gorm:"type:varchar(50)" json:"tssy"`                              //
	Zdysl    decimal.Decimal `gorm:"type:decimal(13,3)" json:"zdysl"`                           //
	Yslfw    string          `gorm:"type:varchar(50)" json:"yslfw"`                             //
	Tsms     string          `gorm:"type:varchar(50)" json:"tsms"`                              //
}
