package model

import (
	"time"
)

type MksbXh struct {
Model
 XhID string `gorm:"type:varchar(36);unique;unique_index;not null" json:"xh_id"` // 
 Xhmc string `gorm:"type:varchar(50)" json:"xhmc"` // 
 Zyfl string `gorm:"type:varchar(36)" json:"zyfl"` // 
 Sblx string `gorm:"type:varchar(36)" json:"sblx"` // 
 Sm string `gorm:"type:varchar(512)" json:"sm"` // 
 Zt string `gorm:"type:varchar(3)" json:"zt"` // 
}
