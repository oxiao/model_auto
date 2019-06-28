package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type JxgcKcgk struct {
	KcgkID    string          `sql:"index" gorm:"type:varchar(36);primary_key" json:"kcgk_id"` //
	Jsgm      string          `gorm:"type:varchar(50)" json:"jsgm"`                            //
	Jksj      time.Time       `gorm:"type:datetime" json:"jksj"`                               //
	Gygcmj    decimal.Decimal `gorm:"type:decimal(13,3)" json:"gygcmj"`                        //
	Jtbjd     string          `gorm:"type:varchar(50)" json:"jtbjd"`                           //
	Jtmj      decimal.Decimal `gorm:"type:decimal(13,3)" json:"jtmj"`                          //
	Jhscnx    decimal.Decimal `gorm:"type:decimal(13,3)" json:"jhscnx"`                        //
	Fwnx      int             `gorm:"type:int(11)" json:"fwnx"`                                //
	Sjscnl    decimal.Decimal `gorm:"type:decimal(13,3)" json:"sjscnl"`                        //
	Hdscnl    decimal.Decimal `gorm:"type:decimal(13,3)" json:"hdscnl"`                        //
	Zgncl     decimal.Decimal `gorm:"type:decimal(13,3)" json:"zgncl"`                         //
	Zdncl     decimal.Decimal `gorm:"type:decimal(13,3)" json:"zdncl"`                         //
	Pjncl     decimal.Decimal `gorm:"type:decimal(13,3)" json:"pjncl"`                         //
	Jx        string          `gorm:"type:varchar(50)" json:"jx"`                              //
	KtfsJt    string          `gorm:"type:varchar(50)" json:"ktfs_jt"`                         //
	KtfsKcsp  string          `gorm:"type:varchar(3)" json:"ktfs_kcsp"`                        //
	KtfsKczb  string          `gorm:"type:varchar(3)" json:"ktfs_kczb"`                        //
	KtfsDxbz  string          `gorm:"type:varchar(3)" json:"ktfs_dxbz"`                        //
	Kcfs      string          `gorm:"type:varchar(50)" json:"kcfs"`                            //
	Cmff      string          `gorm:"type:varchar(50)" json:"cmff"`                            //
	CmffLb    string          `gorm:"type:varchar(3)" json:"cmff_lb"`                          //
	Tsfs      string          `gorm:"type:varchar(3)" json:"tsfs"`                             //
	Jcbcb     decimal.Decimal `gorm:"type:decimal(13,3)" json:"jcbcb"`                         //
	Zkbcb     decimal.Decimal `gorm:"type:decimal(13,3)" json:"zkbcb"`                         //
	Hcl       decimal.Decimal `gorm:"type:decimal(13,3)" json:"hcl"`                           //
	Jdgd      decimal.Decimal `gorm:"type:decimal(13,3)" json:"jdgd"`                          //
	Kcqd      decimal.Decimal `gorm:"type:decimal(13,3)" json:"kcqd"`                          //
	Pzkcmccs  decimal.Decimal `gorm:"type:decimal(13,3)" json:"pzkcmccs"`                      //
	PzkcmcID  decimal.Decimal `gorm:"type:decimal(13,3)" json:"pzkcmc_id"`                     //
	Pzkcmczhd decimal.Decimal `gorm:"type:decimal(13,3)" json:"pzkcmczhd"`                     //
	XcmcID    decimal.Decimal `gorm:"type:decimal(13,3)" json:"xcmc_id"`                       //
	Xcmcqj    decimal.Decimal `gorm:"type:decimal(13,3)" json:"xcmcqj"`                        //
	Xcmczhd   decimal.Decimal `gorm:"type:decimal(13,3)" json:"xcmczhd"`                       //
	Xcmcdcgd  decimal.Decimal `gorm:"type:decimal(13,3)" json:"xcmcdcgd"`                      //
	Ckqmj     decimal.Decimal `gorm:"type:decimal(13,3)" json:"ckqmj"`                         //
	Dykcspsd  decimal.Decimal `gorm:"type:decimal(13,3)" json:"dykcspsd"`                      //
	DekcspSd  decimal.Decimal `gorm:"type:decimal(13,3)" json:"dekcsp_sd"`                     //
	Jzgd      decimal.Decimal `gorm:"type:decimal(13,3)" json:"jzgd"`                          //
	Kjtf      string          `gorm:"type:varchar(50)" json:"kjtf"`                            //
	Tfff      string          `gorm:"type:varchar(50)" json:"tfff"`                            //
	Tffs      string          `gorm:"type:varchar(50)" json:"tffs"`                            //
	Jfj       string          `gorm:"type:varchar(50)" json:"jfj"`                             //
	Hfj       string          `gorm:"type:varchar(50)" json:"hfj"`                             //
	Cqtf      string          `gorm:"type:varchar(50)" json:"cqtf"`                            //
	Jjtf      string          `gorm:"type:varchar(50)" json:"jjtf"`                            //
}
