package model

import (
	"time"
)

type MkKczgz struct {
Model
 KczgzID string `gorm:"type:varchar(36);unique;unique_index;not null" json:"kczgz_id"` // 
 MkjbxxID string `gorm:"type:varchar(36)" json:"mkjbxx_id"` // 
 KzXm string `gorm:"type:varchar(50)" json:"kz_xm"` // 
 KzzgZh string `gorm:"type:varchar(50)" json:"kzzg_zh"` // 
 KzzgYxq time.Time `gorm:"type:datetime" json:"kzzg_yxq"` // 
 AqzgZh string `gorm:"type:varchar(50)" json:"aqzg_zh"` // 
 AqzgYxq string `gorm:"type:varchar(50)" json:"aqzg_yxq"` // 
}
