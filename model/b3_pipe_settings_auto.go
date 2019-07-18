package model

import (
	"time"
)

type B3PipeSetting struct {
Model
 ID int `gorm:"type:bigint(20) unsigned;unique;unique_index;not null" json:"id"` // 
 CreatedAt time.Time `gorm:"type:timestamp" json:"created_at"` // 
 UpdatedAt time.Time `gorm:"type:timestamp" json:"updated_at"` // 
 DeletedAt time.Time `gorm:"type:timestamp" json:"deleted_at"` // 
 Category string `gorm:"type:varchar(32)" json:"category"` // 
 Name string `gorm:"type:varchar(64)" json:"name"` // 
 Value string `gorm:"type:text" json:"value"` // 
 BlogID int `gorm:"type:bigint(20) unsigned" json:"blog_id"` // 
}
