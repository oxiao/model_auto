package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type JxgcJjgzmSc struct {
	JjgzmID string          `gorm:"type:varchar(36)" json:"jjgzm_id"` //
	JjmMc   string          `gorm:"type:varchar(50)" json:"jjm_mc"`   //
	ScN     string          `gorm:"type:varchar(4)" json:"sc_n"`      //
	ScY     string          `gorm:"type:varchar(2)" json:"sc_y"`      //
	ScX     string          `gorm:"type:varchar(4)" json:"sc_x"`      //
	YsJc    decimal.Decimal `gorm:"type:decimal(18,3)" json:"ys_jc"`  //
	MxJc    decimal.Decimal `gorm:"type:decimal(18,3)" json:"mx_jc"`  //
	BmJc    decimal.Decimal `gorm:"type:decimal(18,3)" json:"bm_jc"`  //
	YxJc    decimal.Decimal `gorm:"type:decimal(18,3)" json:"yx_jc"`  //
	Scs     decimal.Decimal `gorm:"type:decimal(18,3)" json:"scs"`    //
	Sfxy    string          `gorm:"type:varchar(3)" json:"sfxy"`      //
	Scrq    time.Time       `gorm:"type:datetime" json:"scrq"`        //
	Jlsj    time.Time       `gorm:"type:datetime" json:"jlsj"`        //
}
