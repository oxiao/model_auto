package model

import (
	"time"
)

type B3PipeArchive struct {
Model
 ID int `gorm:"type:bigint(20) unsigned;unique;unique_index;not null" json:"id"` // 
 CreatedAt time.Time `gorm:"type:timestamp" json:"created_at"` // 
 UpdatedAt time.Time `gorm:"type:timestamp" json:"updated_at"` // 
 DeletedAt time.Time `gorm:"type:timestamp" json:"deleted_at"` // 
 Year string `gorm:"type:varchar(4)" json:"year"` // 
 Month string `gorm:"type:varchar(2)" json:"month"` // 
 ArticleCount int `gorm:"type:int(11)" json:"article_count"` // 
 BlogID int `gorm:"type:bigint(20) unsigned" json:"blog_id"` // 
}
