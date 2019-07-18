package model

import (
	"time"
)

type JxgcZkzt struct {
Model
 ZkztID string `gorm:"type:varchar(36);unique;unique_index;not null" json:"zkzt_id"` // 
 Zklx string `gorm:"type:varchar(50)" json:"zklx"` // 
 Ztlx string `gorm:"type:varchar(50)" json:"ztlx"` // 
 Zkgc string `gorm:"type:varchar(50)" json:"zkgc"` // 
 Zksj string `gorm:"type:varchar(50)" json:"zksj"` // 
 Zksg string `gorm:"type:varchar(50)" json:"zksg"` // 
 Zjsb string `gorm:"type:varchar(50)" json:"zjsb"` // 
 Ztry string `gorm:"type:varchar(50)" json:"ztry"` // 
 Aqyh string `gorm:"type:varchar(50)" json:"aqyh"` // 
 Zcsj string `gorm:"type:varchar(50)" json:"zcsj"` // 
 Zcsg string `gorm:"type:varchar(50)" json:"zcsg"` // 
}
