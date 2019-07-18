package model

import (
	"time"
)

type BmMisLink struct {
Model
 SrcMisID string `gorm:"type:varchar(36);unique;unique_index;not null" json:"src_mis_id"` // 
 DstMisID string `gorm:"type:varchar(36);unique;unique_index;not null" json:"dst_mis_id"` // 
 Parms string `gorm:"type:varchar(100)" json:"parms"` // 
}
