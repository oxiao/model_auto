package model

import (
	"time"
)

type B3PipeCategorie struct {
Model
 ID int `gorm:"type:bigint(20) unsigned;unique;unique_index;not null" json:"id"` // 
 CreatedAt time.Time `gorm:"type:timestamp" json:"created_at"` // 
 UpdatedAt time.Time `gorm:"type:timestamp" json:"updated_at"` // 
 DeletedAt time.Time `gorm:"type:timestamp" json:"deleted_at"` // 
 Title string `gorm:"type:varchar(128)" json:"title"` // 
 Path string `gorm:"type:varchar(255)" json:"path"` // 
 Description string `gorm:"type:varchar(255)" json:"description"` // 
 MetaKeywords string `gorm:"type:varchar(255)" json:"meta_keywords"` // 
 MetaDescription string `gorm:"type:text" json:"meta_description"` // 
 Tags string `gorm:"type:text" json:"tags"` // 
 Number int `gorm:"type:int(11)" json:"number"` // 
 BlogID int `gorm:"type:bigint(20) unsigned" json:"blog_id"` // 
}
