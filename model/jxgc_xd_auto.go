package model

import "github.com/shopspring/decimal"

type JxgcXd struct {
	XdxxID string          `sql:"index" gorm:"type:varchar(36);primary_key" json:"xdxx_id"` //
	Bm     string          `gorm:"type:varchar(36)" json:"bm"`                              //
	Mc     string          `gorm:"type:varchar(50)" json:"mc"`                              //
	Lx     string          `gorm:"type:varchar(3)" json:"lx"`                               //
	Yt     string          `gorm:"type:varchar(3)" json:"yt"`                               //
	Fl     decimal.Decimal `gorm:"type:decimal(18,3)" json:"fl"`                            //
	Fs     decimal.Decimal `gorm:"type:decimal(18,3)" json:"fs"`                            //
}
