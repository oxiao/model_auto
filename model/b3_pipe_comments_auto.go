package model

import (
	"time"
)

type B3PipeComment struct {
Model
 ID int `gorm:"type:bigint(20) unsigned;unique;unique_index;not null" json:"id"` // 
 CreatedAt time.Time `gorm:"type:timestamp" json:"created_at"` // 
 UpdatedAt time.Time `gorm:"type:timestamp" json:"updated_at"` // 
 DeletedAt time.Time `gorm:"type:timestamp" json:"deleted_at"` // 
 ArticleID int `gorm:"type:bigint(20) unsigned" json:"article_id"` // 
 AuthorID int `gorm:"type:bigint(20) unsigned" json:"author_id"` // 
 Content string `gorm:"type:text" json:"content"` // 
 ParentCommentID int `gorm:"type:bigint(20) unsigned" json:"parent_comment_id"` // 
 Ip string `gorm:"type:varchar(128)" json:"ip"` // 
 UserAgent string `gorm:"type:varchar(255)" json:"user_agent"` // 
 PushedAt time.Time `gorm:"type:timestamp" json:"pushed_at"` // 
 AuthorName string `gorm:"type:varchar(32)" json:"author_name"` // 
 AuthorAvatarUrl string `gorm:"type:varchar(255)" json:"author_avatar_url"` // 
 AuthorUrl string `gorm:"type:varchar(255)" json:"author_url"` // 
 BlogID int `gorm:"type:bigint(20) unsigned" json:"blog_id"` // 
}
