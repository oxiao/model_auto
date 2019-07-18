package model

import (
	"time"
)

type MkJbxx struct {
Model
 XxID string `gorm:"type:varchar(36);unique;unique_index;not null" json:"xx_id"` // 
 Kjmc string `gorm:"type:varchar(50)" json:"kjmc"` // 
 Zymc string `gorm:"type:varchar(50)" json:"zymc"` // 
 Ssdq string `gorm:"type:varchar(36)" json:"ssdq"` // 所属地区、省、市
 Xxdz string `gorm:"type:varchar(150)" json:"xxdz"` // 
 Zyfzr string `gorm:"type:varchar(20)" json:"zyfzr"` // 
 Sj string `gorm:"type:varchar(20)" json:"sj"` // 
 Dh string `gorm:"type:varchar(20)" json:"dh"` // 
 Qyxz string `gorm:"type:varchar(36)" json:"qyxz"` // 
 Sjsj time.Time `gorm:"type:datetime" json:"sjsj"` // 
 Tcsj time.Time `gorm:"type:datetime" json:"tcsj"` // 
 Lsgx string `gorm:"type:varchar(10)" json:"lsgx"` // 
 Kjzt string `gorm:"type:varchar(36)" json:"kjzt"` // 
 Zjjkzb string `gorm:"type:varchar(20)" json:"zjjkzb"` // 
 Ckxkz string `gorm:"type:varchar(20)" json:"ckxkz"` // 
}
