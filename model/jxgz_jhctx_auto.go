package model

import (
	"time"
)

type JxgzJhctx struct {
Model
 JhctxID string `gorm:"type:varchar(36);unique;unique_index;not null" json:"jhctx_id"` // 
 CmgzmID string `gorm:"type:varchar(36)" json:"cmgzm_id"` // 
 McID string `gorm:"type:varchar(36)" json:"mc_id"` // 
 Mc string `gorm:"type:varchar(150)" json:"mc"` // 
 JhctSj time.Time `gorm:"type:datetime" json:"jhct_sj"` // 
 JlSj time.Time `gorm:"type:datetime" json:"jl_sj"` // 
}
