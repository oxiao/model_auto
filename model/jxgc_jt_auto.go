package model

import "github.com/shopspring/decimal"

type JxgcJt struct {
	JtxxID   string          `sql:"index" gorm:"type:varchar(36);primary_key" json:"jtxx_id"` //
	Jtmc     string          `gorm:"type:varchar(50)" json:"jtmc"`                            //
	Ktfs     string          `gorm:"type:varchar(3)" json:"ktfs"`                             //
	Jtlx     string          `gorm:"type:varchar(3)" json:"jtlx"`                             //
	Jkbg     decimal.Decimal `gorm:"type:decimal(13,3)" json:"jkbg"`                          //
	Jtwz     string          `gorm:"type:varchar(50)" json:"jtwz"`                            //
	Jkyt     string          `gorm:"type:varchar(50)" json:"jkyt"`                            //
	Jtlb     string          `gorm:"type:varchar(50)" json:"jtlb"`                            //
	JtDmXz   string          `gorm:"type:varchar(3)" json:"jt_dm_xz"`                         //
	JddmMj   decimal.Decimal `gorm:"type:decimal(18,3)" json:"jddm_mj"`                       //
	JtDmZhfs string          `gorm:"type:varchar(3)" json:"jt_dm_zhfs"`                       //
	JtdmJk   decimal.Decimal `gorm:"type:decimal(18,3)" json:"jtdm_jk"`                       //
	JtSd     decimal.Decimal `gorm:"type:decimal(18,3)" json:"jt_sd"`                         //
	Jxcd     decimal.Decimal `gorm:"type:decimal(18,3)" json:"jxcd"`                          //
	Jdcd     decimal.Decimal `gorm:"type:decimal(18,3)" json:"jdcd"`                          //
	Tszb     string          `gorm:"type:varchar(500)" json:"tszb"`                           //
	Btsgff   string          `gorm:"type:varchar(50)" json:"btsgff"`                          //
	JtZhfs   string          `gorm:"type:varchar(3)" json:"jt_zhfs"`                          //
	JtTsnl   string          `gorm:"type:varchar(50)" json:"jt_tsnl"`                         //
}
