package model

import (
	"time"
)

type JxgcXdbz struct {
Model
 QyID string `gorm:"type:varchar(36);unique;unique_index;not null" json:"qy_id"` // 
 CmgzmID string `gorm:"type:varchar(36)" json:"cmgzm_id"` // 
 Lb string `gorm:"type:varchar(3)" json:"lb"` // 切眼,风巷,机巷
 Bh string `gorm:"type:varchar(50)" json:"bh"` // 
 Mc string `gorm:"type:varchar(50)" json:"mc"` // 
 Tjx string `gorm:"type:varchar(50)" json:"tjx"` // 
}
