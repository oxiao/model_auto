package model

import "github.com/shopspring/decimal"

type MzGysy struct {
	MdgyxsyID  string          `sql:"index" gorm:"type:varchar(36);primary_key" json:"mdgyxsy_id"` //
	MzID       string          `gorm:"type:varchar(36)" json:"mz_id"`                              //
	Jzczscd    string          `gorm:"type:varchar(50)" json:"jzczscd"`                            //
	Ljzs       decimal.Decimal `gorm:"type:decimal(13,3)" json:"ljzs"`                             //
	Njzs       decimal.Decimal `gorm:"type:decimal(13,3)" json:"njzs"`                             //
	Jzzgzs     decimal.Decimal `gorm:"type:decimal(13,3)" json:"jzzgzs"`                           //
	Zgpzxs     decimal.Decimal `gorm:"type:decimal(13,3)" json:"zgpzxs"`                           //
	Aypzdsy    string          `gorm:"type:varchar(50)" json:"aypzdsy"`                            //
	Jssxcd     string          `gorm:"type:varchar(50)" json:"jssxcd"`                             //
	Pzylsy     string          `gorm:"type:varchar(50)" json:"pzylsy"`                             //
	Gjglsy     string          `gorm:"type:varchar(50)" json:"gjglsy"`                             //
	Lpglsy     string          `gorm:"type:varchar(50)" json:"lpglsy"`                             //
	Ebqkjlljsy string          `gorm:"type:varchar(50)" json:"ebqkjlljsy"`                         //
	Ksqd       decimal.Decimal `gorm:"type:decimal(13,3)" json:"ksqd"`                             //
	Rwdx       string          `gorm:"type:varchar(50)" json:"rwdx"`                               //
	Mdeyhtfyx  string          `gorm:"type:varchar(50)" json:"mdeyhtfyx"`                          //
	Jzx        string          `gorm:"type:varchar(50)" json:"jzx"`                                //
	Kmx        string          `gorm:"type:varchar(50)" json:"kmx"`                                //
	Mhrrx      string          `gorm:"type:varchar(50)" json:"mhrrx"`                              //
	Mhzd       decimal.Decimal `gorm:"type:decimal(13,3)" json:"mhzd"`                             //
	Dmhcdmdtx  string          `gorm:"type:varchar(50)" json:"dmhcdmdtx"`                          //
	Hskmxzs    decimal.Decimal `gorm:"type:decimal(13,3)" json:"hskmxzs"`                          //
	Sfsy       string          `gorm:"type:varchar(50)" json:"sfsy"`                               //
	Mfsfsy     string          `gorm:"type:varchar(50)" json:"mfsfsy"`                             //
	Fcsy       string          `gorm:"type:varchar(50)" json:"fcsy"`                               //
	Mnfcsy     string          `gorm:"type:varchar(50)" json:"mnfcsy"`                             //
	Fxsy       string          `gorm:"type:varchar(50)" json:"fxsy"`                               //
	Rssy       string          `gorm:"type:varchar(50)" json:"rssy"`                               //
}
