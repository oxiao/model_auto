package model

import (
	"time"
)

type B3PipeNavigation struct {
Model
 ID int `gorm:"type:bigint(20) unsigned;unique;unique_index;not null" json:"id"` // 
 CreatedAt time.Time `gorm:"type:timestamp" json:"created_at"` // 
 UpdatedAt time.Time `gorm:"type:timestamp" json:"updated_at"` // 
 DeletedAt time.Time `gorm:"type:timestamp" json:"deleted_at"` // 
 Title string `gorm:"type:varchar(128)" json:"title"` // 
 Url string `gorm:"type:varchar(255)" json:"url"` // 
 IconUrl string `gorm:"type:varchar(255)" json:"icon_url"` // 
 OpenMethod string `gorm:"type:varchar(32)" json:"open_method"` // 
 Number int `gorm:"type:int(11)" json:"number"` // 
 BlogID int `gorm:"type:bigint(20) unsigned" json:"blog_id"` // 
}
