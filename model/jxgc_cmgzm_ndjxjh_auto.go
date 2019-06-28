package model

import "github.com/shopspring/decimal"

type JxgcCmgzmNdjxjh struct {
	CmgzmID string          `gorm:"type:varchar(36)" json:"cmgzm_id"` //
	GzmMc   string          `gorm:"type:varchar(50)" json:"gzm_mc"`   //
	Hcgy    string          `gorm:"type:varchar(200)" json:"hcgy"`    //
	Cg      decimal.Decimal `gorm:"type:decimal(18,3)" json:"cg"`     //
	Rz      decimal.Decimal `gorm:"type:decimal(18,3)" json:"rz"`     //
	Qj      decimal.Decimal `gorm:"type:decimal(18,3)" json:"qj"`     //
	CmCd    decimal.Decimal `gorm:"type:decimal(18,3)" json:"cm_cd"`  //
	ZxCd    decimal.Decimal `gorm:"type:decimal(18,3)" json:"zx_cd"`  //
	Sycl    decimal.Decimal `gorm:"type:decimal(18,3)" json:"sycl"`   //
	Jynf    string          `gorm:"type:varchar(4)" json:"jynf"`      //
	Ycl1    decimal.Decimal `gorm:"type:decimal(18,3)" json:"ycl_1"`  //
	Ycl2    decimal.Decimal `gorm:"type:decimal(18,3)" json:"ycl_2"`  //
	Ycl3    decimal.Decimal `gorm:"type:decimal(18,3)" json:"ycl_3"`  //
	Ycl4    decimal.Decimal `gorm:"type:decimal(18,3)" json:"ycl_4"`  //
	Ycl5    decimal.Decimal `gorm:"type:decimal(18,3)" json:"ycl_5"`  //
	Ycl6    decimal.Decimal `gorm:"type:decimal(18,3)" json:"ycl_6"`  //
	Ycl7    decimal.Decimal `gorm:"type:decimal(18,3)" json:"ycl_7"`  //
	Ycl8    decimal.Decimal `gorm:"type:decimal(18,3)" json:"ycl_8"`  //
	Ycl9    decimal.Decimal `gorm:"type:decimal(18,3)" json:"ycl_9"`  //
	Ycl10   decimal.Decimal `gorm:"type:decimal(18,3)" json:"ycl_10"` //
	Ycl11   decimal.Decimal `gorm:"type:decimal(18,3)" json:"ycl_11"` //
	Ycl12   decimal.Decimal `gorm:"type:decimal(18,3)" json:"ycl_12"` //
}
