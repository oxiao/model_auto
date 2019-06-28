package model

import (
	"time"
)

type JxgzJhctx struct {
	JhctxID string    `sql:"index" gorm:"type:varchar(36);primary_key" json:"jhctx_id"` //
	CmgzmID string    `gorm:"type:varchar(36)" json:"cmgzm_id"`                         //
	McID    string    `gorm:"type:varchar(36)" json:"mc_id"`                            //
	Mc      string    `gorm:"type:varchar(150)" json:"mc"`                              //
	JhctSj  time.Time `gorm:"type:datetime" json:"jhct_sj"`                             //
	JlSj    time.Time `gorm:"type:datetime" json:"jl_sj"`                               //
}
