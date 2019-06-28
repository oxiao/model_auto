package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type JxgcJjgzmSgrz struct {
	JjgzmID string          `gorm:"type:varchar(36)" json:"jjgzm_id"` //
	RzBh    string          `gorm:"type:varchar(50)" json:"rz_bh"`    //
	JjRq    time.Time       `gorm:"type:datetime" json:"jj_rq"`       //
	JjmMc   string          `gorm:"type:varchar(50)" json:"jjm_mc"`   //
	Jcs     decimal.Decimal `gorm:"type:decimal(18,3)" json:"jcs"`    //
	Ljjcs   decimal.Decimal `gorm:"type:decimal(18,3)" json:"ljjcs"`  //
	Cqrs    int             `gorm:"type:int(11)" json:"cqrs"`         //
	Dbgx    int             `gorm:"type:int(11)" json:"dbgx"`         //
	Xbgx    string          `gorm:"type:varchar(50)" json:"xbgx"`     //
	Sfys    string          `gorm:"type:varchar(3)" json:"sfys"`      //
	Ysjc    decimal.Decimal `gorm:"type:decimal(18,3)" json:"ysjc"`   //
	Jlr     string          `gorm:"type:varchar(50)" json:"jlr"`      //
	Ysr     string          `gorm:"type:varchar(36)" json:"ysr"`      //
	Sfst    string          `gorm:"type:varchar(3)" json:"sfst"`      //
	DRecord time.Time       `gorm:"type:datetime" json:"d_record"`    //
	Yssj    time.Time       `gorm:"type:datetime" json:"yssj"`        //
	Bz      string          `gorm:"type:varchar(500)" json:"bz"`      //
}
