package model

import "github.com/shopspring/decimal"

type ZyclMtcl struct {
	MtclID      string          `sql:"index" gorm:"type:varchar(36);primary_key" json:"mtcl_id"` //
	Lnkcwczb    decimal.Decimal `gorm:"type:decimal(13,3)" json:"lnkcwczb"`                      //
	DzCl        decimal.Decimal `gorm:"type:decimal(13,3)" json:"dz_cl"`                         //
	GyCl        decimal.Decimal `gorm:"type:decimal(13,3)" json:"gy_cl"`                         //
	SjCl        decimal.Decimal `gorm:"type:decimal(13,3)" json:"sj_cl"`                         //
	SjkcCl      decimal.Decimal `gorm:"type:decimal(13,3)" json:"sjkc_cl"`                       //
	QjCl        decimal.Decimal `gorm:"type:decimal(13,3)" json:"qj_cl"`                         //
	McCl        decimal.Decimal `gorm:"type:decimal(13,3)" json:"mc_cl"`                         //
	SpCl        decimal.Decimal `gorm:"type:decimal(13,3)" json:"sp_cl"`                         //
	CqCl        decimal.Decimal `gorm:"type:decimal(13,3)" json:"cq_cl"`                         //
	MzCl        decimal.Decimal `gorm:"type:decimal(13,3)" json:"mz_cl"`                         //
	KcCl        decimal.Decimal `gorm:"type:decimal(13,3)" json:"kc_cl"`                         //
	SjNcl       decimal.Decimal `gorm:"type:decimal(13,3)" json:"sj_ncl"`                        //
	ClByxs      decimal.Decimal `gorm:"type:decimal(13,3)" json:"cl_byxs"`                       //
	TmJjKccl    decimal.Decimal `gorm:"type:decimal(13,3)" json:"tm_jj_kccl"`                    //
	TmJjJccl    decimal.Decimal `gorm:"type:decimal(13,3)" json:"tm_jj_jccl"`                    //
	TmJjYkccl   decimal.Decimal `gorm:"type:decimal(13,3)" json:"tm_jj_ykccl"`                   //
	TmJjYkcjccl decimal.Decimal `gorm:"type:decimal(13,3)" json:"tm_jj_ykcjccl"`                 //
	KzJjYkccl   decimal.Decimal `gorm:"type:decimal(13,3)" json:"kz_jj_ykccl"`                   //
	KzJjYkcjccl decimal.Decimal `gorm:"type:decimal(13,3)" json:"kz_jj_ykcjccl"`                 //
	TmBjJccl    decimal.Decimal `gorm:"type:decimal(13,3)" json:"tm_bj_jccl"`                    //
	TmBjKccl    decimal.Decimal `gorm:"type:decimal(13,3)" json:"tm_bj_kccl"`                    //
	KzBjJccl    decimal.Decimal `gorm:"type:decimal(13,3)" json:"kz_bj_jccl"`                    //
	TmCbjJccl   decimal.Decimal `gorm:"type:decimal(13,3)" json:"tm_cbj_jccl"`                   //
	TmCbjKccl   decimal.Decimal `gorm:"type:decimal(13,3)" json:"tm_cbj_kccl"`                   //
	KzCbjJccl   decimal.Decimal `gorm:"type:decimal(13,3)" json:"kz_cbj_jccl"`                   //
	TmNyCl      decimal.Decimal `gorm:"type:decimal(13,3)" json:"tm_ny_cl"`                      //
	KzNyCl      decimal.Decimal `gorm:"type:decimal(13,3)" json:"kz_ny_cl"`                      //
	TcNyCl      decimal.Decimal `gorm:"type:decimal(13,3)" json:"tc_ny_cl"`                      //
	YcNyCl      decimal.Decimal `gorm:"type:decimal(13,3)" json:"yc_ny_cl"`                      //
}
