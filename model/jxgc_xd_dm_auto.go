package model

import "github.com/shopspring/decimal"

type JxgcXdDm struct {
	XddmID   string          `sql:"index" gorm:"type:varchar(36);primary_key" json:"xddm_id"` //
	XdxxID   string          `gorm:"type:varchar(36)" json:"xdxx_id"`                         //
	Dmxz     string          `gorm:"type:varchar(50)" json:"dmxz"`                            //
	Jkd      decimal.Decimal `gorm:"type:decimal(13,3)" json:"jkd"`                           //
	Jgd      decimal.Decimal `gorm:"type:decimal(13,3)" json:"jgd"`                           //
	Jmj      decimal.Decimal `gorm:"type:decimal(13,3)" json:"jmj"`                           //
	Zjlx     string          `gorm:"type:varchar(50)" json:"zjlx"`                            //
	Zhcl     string          `gorm:"type:varchar(50)" json:"zhcl"`                            //
	Zhjg     string          `gorm:"type:varchar(50)" json:"zhjg"`                            //
	Zhxgjc   string          `gorm:"type:varchar(50)" json:"zhxgjc"`                          //
	Ggcc     decimal.Decimal `gorm:"type:decimal(13,3)" json:"ggcc"`                          //
	Ggxh     string          `gorm:"type:varchar(50)" json:"ggxh"`                            //
	Gzgg     string          `gorm:"type:varchar(50)" json:"gzgg"`                            //
	Dcgd     decimal.Decimal `gorm:"type:decimal(13,3)" json:"dcgd"`                          //
	Sjjjdmcc decimal.Decimal `gorm:"type:decimal(13,3)" json:"sjjjdmcc"`                      //
	Yxzdz    decimal.Decimal `gorm:"type:decimal(13,3)" json:"yxzdz"`                         //
	Jsjjdmcc decimal.Decimal `gorm:"type:decimal(13,3)" json:"jsjjdmcc"`                      //
	Qj       string          `gorm:"type:varchar(50)" json:"qj"`                              //
	Sgbz     string          `gorm:"type:varchar(50)" json:"sgbz"`                            //
	Sgqz     string          `gorm:"type:varchar(50)" json:"sgqz"`                            //
	Sgpd     decimal.Decimal `gorm:"type:decimal(13,3)" json:"sgpd"`                          //
	Sgls     string          `gorm:"type:varchar(50)" json:"sgls"`                            //
	Sgdm     string          `gorm:"type:varchar(50)" json:"sgdm"`                            //
	Sggb     string          `gorm:"type:varchar(50)" json:"sggb"`                            //
	Gdbz     string          `gorm:"type:varchar(50)" json:"gdbz"`                            //
	Dlbz     string          `gorm:"type:varchar(50)" json:"dlbz"`                            //
}
