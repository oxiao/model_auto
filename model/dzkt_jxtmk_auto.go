package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type DzktJxtmk struct {
	JbzlID string          `sql:"index" gorm:"type:varchar(36);primary_key" json:"jbzl_id"` //
	McID   string          `gorm:"type:varchar(36)" json:"mc_id"`                           //
	Zkwz   string          `gorm:"type:varchar(50)" json:"zkwz"`                            //
	Jbzdjl decimal.Decimal `gorm:"type:decimal(13,3)" json:"jbzdjl"`                        //
	Ktmd   string          `gorm:"type:varchar(50)" json:"ktmd"`                            //
	Fx     string          `gorm:"type:varchar(50)" json:"fx"`                              //
	Jd     decimal.Decimal `gorm:"type:decimal(13,3)" json:"jd"`                            //
	Kgrq   time.Time       `gorm:"type:datetime" json:"kgrq"`                               //
	Jgrq   time.Time       `gorm:"type:datetime" json:"jgrq"`                               //
	Zksd   decimal.Decimal `gorm:"type:decimal(13,3)" json:"zksd"`                          //
	Sgdw   time.Time       `gorm:"type:datetime" json:"sgdw"`                               //
	Mcmc   string          `gorm:"type:varchar(50)" json:"mcmc"`                            //
	Dbbg   decimal.Decimal `gorm:"type:decimal(13,3)" json:"dbbg"`                          //
	Mh     decimal.Decimal `gorm:"type:decimal(13,3)" json:"mh"`                            //
	Jzhd   decimal.Decimal `gorm:"type:decimal(13,3)" json:"jzhd"`                          //
	TsCw   string          `gorm:"type:varchar(50)" json:"ts_cw"`                           //
	TsCh   decimal.Decimal `gorm:"type:decimal(13,3)" json:"ts_ch"`                         //
	TsKs   decimal.Decimal `gorm:"type:decimal(13,3)" json:"ts_ks"`                         //
	TsLjhd decimal.Decimal `gorm:"type:decimal(13,3)" json:"ts_ljhd"`                       //
	TsYsl  decimal.Decimal `gorm:"type:decimal(13,3)" json:"ts_ysl"`                        //
	TsSy   decimal.Decimal `gorm:"type:decimal(13,3)" json:"ts_sy"`                         //
	TsLscd decimal.Decimal `gorm:"type:decimal(13,3)" json:"ts_lscd"`                       //
	Zkzz   string          `gorm:"type:varchar(50)" json:"zkzz"`                            //
	Zkjg   string          `gorm:"type:varchar(50)" json:"zkjg"`                            //
	Zkfk   string          `gorm:"type:varchar(50)" json:"zkfk"`                            //
	Ysgc   string          `gorm:"type:varchar(50)" json:"ysgc"`                            //
	Pfs    string          `gorm:"type:varchar(50)" json:"pfs"`                             //
}
