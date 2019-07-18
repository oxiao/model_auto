package model

import (
	"time"
)

type BmFile struct {
Model
 FileID string `gorm:"type:varchar(36);unique;unique_index;not null" json:"file_id"` // 
 Name string `gorm:"type:varchar(256)" json:"name"` // 
 Path string `gorm:"type:varchar(512)" json:"path"` // 
 Type string `gorm:"type:varchar(36)" json:"type"` // 
 Size int `gorm:"type:int(11)" json:"size"` // 
 Dist string `gorm:"type:varchar(36)" json:"dist"` // 
 CreateDate time.Time `gorm:"type:datetime" json:"create_date"` // 
 Creator string `gorm:"type:varchar(36)" json:"creator"` // 
 State string `gorm:"type:varchar(3)" json:"state"` // 
 StateTime time.Time `gorm:"type:datetime" json:"state_time"` // 标志更改时的时间
}
