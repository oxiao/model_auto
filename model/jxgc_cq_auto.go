package model

import "github.com/shopspring/decimal"

type JxgcCq struct {
	CqID     string          `sql:"index" gorm:"type:varchar(36);primary_key" json:"cq_id"` //
	McID     string          `sql:"index" gorm:"type:varchar(36);primary_key" json:"mc_id"` //
	JdID     string          `gorm:"type:varchar(36)" json:"jd_id"`                         //
	Cqmc     string          `gorm:"type:varchar(50)" json:"cqmc"`                          //
	Kcsp     string          `gorm:"type:varchar(50)" json:"kcsp"`                          //
	Kcmc     string          `gorm:"type:varchar(50)" json:"kcmc"`                          //
	Hcmxx    string          `gorm:"type:varchar(50)" json:"hcmxx"`                         //
	Ncl      decimal.Decimal `gorm:"type:decimal(13,3)" json:"ncl"`                         //
	Cqcl     decimal.Decimal `gorm:"type:decimal(13,3)" json:"cqcl"`                        //
	Scnl     decimal.Decimal `gorm:"type:decimal(13,3)" json:"scnl"`                        //
	Fwnx     decimal.Decimal `gorm:"type:decimal(13,3)" json:"fwnx"`                        //
	Cqwzfw   string          `gorm:"type:varchar(50)" json:"cqwzfw"`                        //
	Kctj     string          `gorm:"type:varchar(50)" json:"kctj"`                          //
	Xdbz     string          `gorm:"type:varchar(50)" json:"xdbz"`                          //
	Qdhf     string          `gorm:"type:varchar(50)" json:"qdhf"`                          //
	Cc       string          `gorm:"type:varchar(50)" json:"cc"`                            //
	CqMc     string          `gorm:"type:varchar(50)" json:"cq_mc"`                         //
	Cqjcf    string          `gorm:"type:varchar(50)" json:"cqjcf"`                         //
	Cmff     string          `gorm:"type:varchar(50)" json:"cmff"`                          //
	Yssb     string          `gorm:"type:varchar(50)" json:"yssb"`                          //
	Yssb2    string          `gorm:"type:varchar(50)" json:"yssb2"`                         //
	Cqgfl    decimal.Decimal `gorm:"type:decimal(13,3)" json:"cqgfl"`                       //
	Cqhfl    decimal.Decimal `gorm:"type:decimal(13,3)" json:"cqhfl"`                       //
	Tfxl     string          `gorm:"type:varchar(50)" json:"tfxl"`                          //
	Glxt     string          `gorm:"type:varchar(50)" json:"glxt"`                          //
	Ajkjtxxt string          `gorm:"type:varchar(50)" json:"ajkjtxxt"`                      //
	Wscfcs   string          `gorm:"type:varchar(50)" json:"wscfcs"`                        //
	Bzlxcs   string          `gorm:"type:varchar(50)" json:"bzlxcs"`                        //
	Fccs     string          `gorm:"type:varchar(50)" json:"fccs"`                          //
	Fhcs     string          `gorm:"type:varchar(50)" json:"fhcs"`                          //
	Fscs     string          `gorm:"type:varchar(50)" json:"fscs"`                          //
	Cqgd     string          `gorm:"type:varchar(50)" json:"cqgd"`                          //
	Cqps     string          `gorm:"type:varchar(50)" json:"cqps"`                          //
}
