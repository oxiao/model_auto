package model

import "github.com/shopspring/decimal"

type JxgcCmgzmFcwx struct {
	GzmBm  string          `sql:"index" gorm:"type:varchar(36);primary_key" json:"gzm_bm"` //
	McID   string          `gorm:"type:varchar(36)" json:"mc_id"`                          //
	GzmMc  string          `gorm:"type:varchar(100)" json:"gzm_mc"`                        //
	FcFsd  string          `gorm:"type:varchar(50)" json:"fc_fsd"`                         //
	FcXfx  string          `gorm:"type:varchar(50)" json:"fc_xfx"`                         //
	FcXifx string          `gorm:"type:varchar(50)" json:"fc_xifx"`                        //
	FcSrx  string          `gorm:"type:varchar(50)" json:"fc_srx"`                         // 亲水性矿尘,疏水性矿尘
	FcDhx  string          `gorm:"type:varchar(50)" json:"fc_dhx"`                         // 正电荷,负电荷
	FcHxx  string          `gorm:"type:varchar(50)" json:"fc_hxx"`                         // 矽尘,非矽尘
	FcBzx  string          `gorm:"type:varchar(50)" json:"fc_bzx"`                         //
	Sio2   decimal.Decimal `gorm:"type:decimal(18,3)" json:"sio2"`                         //
	Kcld   decimal.Decimal `gorm:"type:decimal(18,3)" json:"kcld"`                         // （um）
	Bbmj   decimal.Decimal `gorm:"type:decimal(18,3)" json:"bbmj"`                         // （m2/kg或cm2/g）
	Sfhl   decimal.Decimal `gorm:"type:decimal(18,3)" json:"sfhl"`                         //
	Hff    decimal.Decimal `gorm:"type:decimal(18,3)" json:"hff"`                          //
	Frl    decimal.Decimal `gorm:"type:decimal(18,3)" json:"frl"`                          //
	Zmd    decimal.Decimal `gorm:"type:decimal(18,3)" json:"zmd"`                          //
	Kxl    decimal.Decimal `gorm:"type:decimal(18,3)" json:"kxl"`                          //
	McKzx  decimal.Decimal `gorm:"type:decimal(18,3)" json:"mc_kzx"`                       // 可注,不可注
	McXsx  decimal.Decimal `gorm:"type:decimal(18,3)" json:"mc_xsx"`                       //
	McJgxs decimal.Decimal `gorm:"type:decimal(18,3)" json:"mc_jgxs"`                      //
	Jcj    decimal.Decimal `gorm:"type:decimal(18,3)" json:"jcj"`                          //
	Jmd    decimal.Decimal `gorm:"type:decimal(18,3)" json:"jmd"`                          //
	C      decimal.Decimal `gorm:"type:decimal(18,3)" json:"c"`                            //
	H      decimal.Decimal `gorm:"type:decimal(18,3)" json:"h"`                            //
	O      decimal.Decimal `gorm:"type:decimal(18,3)" json:"o"`                            //
	N      decimal.Decimal `gorm:"type:decimal(18,3)" json:"n"`                            //
	S      decimal.Decimal `gorm:"type:decimal(18,3)" json:"s"`                            //
	Hfhl   decimal.Decimal `gorm:"type:decimal(18,3)" json:"hfhl"`                         //
}
