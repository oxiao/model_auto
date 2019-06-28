package model

import "github.com/shopspring/decimal"

type JxgcCmgzm struct {
	CmgzmID string          `sql:"index" gorm:"type:varchar(36);primary_key" json:"cmgzm_id"` //
	CqID    string          `gorm:"type:varchar(36)" json:"cq_id"`                            //
	McID    string          `gorm:"type:varchar(36)" json:"mc_id"`                            //
	Mc      string          `gorm:"type:varchar(50)" json:"mc"`                               //
	Sscq    string          `gorm:"type:varchar(50)" json:"sscq"`                             //
	Kcmc    string          `gorm:"type:varchar(50)" json:"kcmc"`                             //
	JsjxGx  string          `gorm:"type:varchar(50)" json:"jsjx_gx"`                          //
	GzmWz   string          `gorm:"type:varchar(50)" json:"gzm_wz"`                           //
	Mcfc    string          `gorm:"type:varchar(50)" json:"mcfc"`                             //
	Ddbyx   string          `gorm:"type:varchar(50)" json:"ddbyx"`                            //
	Dzgz    string          `gorm:"type:varchar(50)" json:"dzgz"`                             //
	Swdz    string          `gorm:"type:varchar(50)" json:"swdz"`                             //
	Dwwh    string          `gorm:"type:varchar(50)" json:"dwwh"`                             //
	Cjdy    string          `gorm:"type:varchar(50)" json:"cjdy"`                             //
	Yhqt    string          `gorm:"type:varchar(50)" json:"yhqt"`                             //
	Cl      decimal.Decimal `gorm:"type:decimal(13,3)" json:"cl"`                             //
	Kccl    decimal.Decimal `gorm:"type:decimal(13,3)" json:"kccl"`                           //
	Scnl    decimal.Decimal `gorm:"type:decimal(13,3)" json:"scnl"`                           //
	Kcq     int             `gorm:"type:int(11)" json:"kcq"`                                  //
	McBzwx  string          `gorm:"type:varchar(50)" json:"mc_bzwx"`                          //
	MZrqx   string          `gorm:"type:varchar(3)" json:"m_zrqx"`                            //
	CmmKd   decimal.Decimal `gorm:"type:decimal(13,3)" json:"cmm_kd"`                         //
	CmmGd   decimal.Decimal `gorm:"type:decimal(13,3)" json:"cmm_gd"`                         //
	CmmXc   decimal.Decimal `gorm:"type:decimal(13,3)" json:"cmm_xc"`                         //
	GzmMj   decimal.Decimal `gorm:"type:decimal(13,3)" json:"gzm_mj"`                         //
	SfBf    string          `gorm:"type:varchar(3)" json:"sf_bf"`                             //
	SfFt    string          `gorm:"type:varchar(3)" json:"sf_ft"`                             //
	McQj    decimal.Decimal `gorm:"type:decimal(13,3)" json:"mc_qj"`                          //
	MRz     decimal.Decimal `gorm:"type:decimal(13,3)" json:"m_rz"`                           //
	MYd     decimal.Decimal `gorm:"type:decimal(13,3)" json:"m_yd"`                           //
	Gmgd    decimal.Decimal `gorm:"type:decimal(13,3)" json:"gmgd"`                           //
	XhJc    decimal.Decimal `gorm:"type:decimal(13,3)" json:"xh_jc"`                          //
	XhCl    decimal.Decimal `gorm:"type:decimal(13,3)" json:"xh_cl"`                          //
	XhWcl   decimal.Decimal `gorm:"type:decimal(13,3)" json:"xh_wcl"`                         //
	Rcq     decimal.Decimal `gorm:"type:decimal(13,3)" json:"rcq"`                            //
	HcgXl   decimal.Decimal `gorm:"type:decimal(13,3)" json:"hcg_xl"`                         //
	KmXh    decimal.Decimal `gorm:"type:decimal(13,3)" json:"km_xh"`                          //
	JcXh    decimal.Decimal `gorm:"type:decimal(13,3)" json:"jc_xh"`                          //
	RJd     decimal.Decimal `gorm:"type:decimal(13,3)" json:"r_jd"`                           //
	YJd     decimal.Decimal `gorm:"type:decimal(13,3)" json:"y_jd"`                           //
	RCl     decimal.Decimal `gorm:"type:decimal(13,3)" json:"r_cl"`                           //
	YCl     decimal.Decimal `gorm:"type:decimal(13,3)" json:"y_cl"`                           //
	Cmff    string          `gorm:"type:varchar(50)" json:"cmff"`                             //
	CmZb    string          `gorm:"type:varchar(50)" json:"cm_zb"`                            //
	Jfxl    string          `gorm:"type:varchar(50)" json:"jfxl"`                             //
	Ffxl    string          `gorm:"type:varchar(50)" json:"ffxl"`                             //
	Fl2     decimal.Decimal `gorm:"type:decimal(13,3)" json:"fl2"`                            //
	WsfzYl  decimal.Decimal `gorm:"type:decimal(13,3)" json:"wsfz_yl"`                        //
	WsfzHl  decimal.Decimal `gorm:"type:decimal(13,3)" json:"wsfz_hl"`                        //
	WsfzYcl decimal.Decimal `gorm:"type:decimal(13,3)" json:"wsfz_ycl"`                       //
	Fc      string          `gorm:"type:varchar(50)" json:"fc"`                               //
	Fmh     string          `gorm:"type:varchar(50)" json:"fmh"`                              //
	FsRysl  decimal.Decimal `gorm:"type:decimal(13,3)" json:"fs_rysl"`                        //
	FsZdysl decimal.Decimal `gorm:"type:decimal(13,3)" json:"fs_zdysl"`                       //
	FsFzcs  string          `gorm:"type:varchar(50)" json:"fs_fzcs"`                          //
	Aqjc    string          `gorm:"type:varchar(50)" json:"aqjc"`                             //
	Sgzz    string          `gorm:"type:varchar(50)" json:"sgzz"`                             //
	Ysxt    string          `gorm:"type:varchar(50)" json:"ysxt"`                             //
	Psxt    string          `gorm:"type:varchar(50)" json:"psxt"`                             //
	Gsxt    string          `gorm:"type:varchar(50)" json:"gsxt"`                             //
	Gdxt    string          `gorm:"type:varchar(50)" json:"gdxt"`                             //
	Zhfs    string          `gorm:"type:varchar(50)" json:"zhfs"`                             //
	Kyjk    string          `gorm:"type:varchar(50)" json:"kyjk"`                             //
	Mdcl    string          `gorm:"type:varchar(50)" json:"mdcl"`                             //
	Bzlx    string          `gorm:"type:varchar(50)" json:"bzlx"`                             //
	DwQd    string          `gorm:"type:varchar(50)" json:"dw_qd"`                            //
	DwFxd   string          `gorm:"type:varchar(50)" json:"dw_fxd"`                           //
	DwZl    string          `gorm:"type:varchar(3)" json:"dw_zl"`                             //
	DwJl    decimal.Decimal `gorm:"type:decimal(13,3)" json:"dw_jl"`                          //
	DwQsd   string          `gorm:"type:varchar(50)" json:"dw_qsd"`                           //
	DwDxd   decimal.Decimal `gorm:"type:decimal(13,3)" json:"dw_dxd"`                         //
}
