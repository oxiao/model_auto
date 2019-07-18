package model

import (
	"time"
)

type B3PipeUser struct {
Model
 ID int `gorm:"type:bigint(20) unsigned;unique;unique_index;not null" json:"id"` // 
 CreatedAt time.Time `gorm:"type:timestamp" json:"created_at"` // 
 UpdatedAt time.Time `gorm:"type:timestamp" json:"updated_at"` // 
 DeletedAt time.Time `gorm:"type:timestamp" json:"deleted_at"` // 
 Name string `gorm:"type:varchar(32)" json:"name"` // 
 Nickname string `gorm:"type:varchar(32)" json:"nickname"` // 
 AvatarUrl string `gorm:"type:varchar(255)" json:"avatar_url"` // 
 B3Key string `gorm:"type:varchar(32)" json:"b3_key"` // 
 Locale string `gorm:"type:varchar(32)" json:"locale"` // 
 TotalArticleCount int `gorm:"type:int(11)" json:"total_article_count"` // 
 GithubID string `gorm:"type:varchar(255)" json:"github_id"` // 
}
