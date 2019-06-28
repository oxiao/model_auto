package model

import "github.com/shopspring/decimal"

type WssgBz struct {
	WsbzxxID string          `sql:"index" gorm:"type:varchar(36);primary_key" json:"wsbzxx_id"` //
	Wsjjzb   decimal.Decimal `gorm:"type:decimal(13,3)" json:"wsjjzb"`                          //
	Yqndjc   string          `gorm:"type:varchar(50)" json:"yqndjc"`                            //
	Tfztjc   string          `gorm:"type:varchar(50)" json:"tfztjc"`                            //
	Hypjzb   decimal.Decimal `gorm:"type:decimal(13,3)" json:"hypjzb"`                          //
	Mhzb     string          `gorm:"type:varchar(50)" json:"mhzb"`                              //
	Sbyhzb   decimal.Decimal `gorm:"type:decimal(13,3)" json:"sbyhzb"`                          //
	Wsbzsg   string          `gorm:"type:varchar(50)" json:"wsbzsg"`                            //
}
