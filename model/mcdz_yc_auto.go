package model

import "github.com/shopspring/decimal"

type McdzYc struct {
	YcID   string          `sql:"index" gorm:"type:varchar(36);primary_key" json:"yc_id"` //
	Ycmc   string          `gorm:"type:varchar(50)" json:"ycmc"`                          //
	Dcdw   string          `gorm:"type:varchar(50)" json:"dcdw"`                          //
	Ysdcdw string          `gorm:"type:varchar(50)" json:"ysdcdw"`                        //
	Swdcdw string          `gorm:"type:varchar(50)" json:"swdcdw"`                        //
	Nddcdw string          `gorm:"type:varchar(50)" json:"nddcdw"`                        //
	Dznddw string          `gorm:"type:varchar(50)" json:"dznddw"`                        //
	Sfhmdc string          `gorm:"type:varchar(50)" json:"sfhmdc"`                        //
	Sfdc   string          `gorm:"type:varchar(50)" json:"sfdc"`                          //
	Sfdcyx string          `gorm:"type:varchar(50)" json:"sfdcyx"`                        //
	Sfdchd string          `gorm:"type:varchar(50)" json:"sfdchd"`                        //
	Xfdc   string          `gorm:"type:varchar(50)" json:"xfdc"`                          //
	Xfdcyx string          `gorm:"type:varchar(50)" json:"xfdcyx"`                        //
	Xfdchd string          `gorm:"type:varchar(50)" json:"xfdchd"`                        //
	Hmdcdb string          `gorm:"type:varchar(50)" json:"hmdcdb"`                        //
	Hmdchf string          `gorm:"type:varchar(50)" json:"hmdchf"`                        //
	Dcgk   string          `gorm:"type:varchar(50)" json:"dcgk"`                          //
	Dcgj   string          `gorm:"type:varchar(50)" json:"dcgj"`                          //
	Ymcgx  string          `gorm:"type:varchar(50)" json:"ymcgx"`                         // 上， 下
	Ssc    string          `gorm:"type:varchar(50)" json:"ssc"`                           //
	Sscyx  string          `gorm:"type:varchar(50)" json:"sscyx"`                         //
	Sschd  decimal.Decimal `gorm:"type:decimal(13,3)" json:"sschd"`                       //
	Fhc    string          `gorm:"type:varchar(50)" json:"fhc"`                           //
	Fhcyx  string          `gorm:"type:varchar(50)" json:"fhcyx"`                         //
	Fhchd  decimal.Decimal `gorm:"type:decimal(13,3)" json:"fhchd"`                       //
	Rrjc   string          `gorm:"type:varchar(50)" json:"rrjc"`                          //
	Rrcyx  string          `gorm:"type:varchar(50)" json:"rrcyx"`                         //
	Rrchd  decimal.Decimal `gorm:"type:decimal(13,3)" json:"rrchd"`                       //
	Dchd   string          `gorm:"type:varchar(50)" json:"dchd"`                          // 只存平均值
	Swdc   string          `gorm:"type:varchar(50)" json:"swdc"`                          //
	Yxdc   string          `gorm:"type:varchar(50)" json:"yxdc"`                          //
	Yccz   string          `gorm:"type:varchar(50)" json:"yccz"`                          //
}
