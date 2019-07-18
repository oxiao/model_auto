package model

import (
	"time"
)

type B3PipeArticle struct {
Model
 ID int `gorm:"type:bigint(20) unsigned;unique;unique_index;not null" json:"id"` // 
 CreatedAt time.Time `gorm:"type:timestamp" json:"created_at"` // 
 UpdatedAt time.Time `gorm:"type:timestamp" json:"updated_at"` // 
 DeletedAt time.Time `gorm:"type:timestamp" json:"deleted_at"` // 
 AuthorID int `gorm:"type:bigint(20) unsigned" json:"author_id"` // 
 Title string `gorm:"type:varchar(128)" json:"title"` // 
 Abstract mediumtext `gorm:"type:mediumtext" json:"abstract"` // 
 Tags string `gorm:"type:text" json:"tags"` // 
 Content mediumtext `gorm:"type:mediumtext" json:"content"` // 
 Path string `gorm:"type:varchar(255)" json:"path"` // 
 Status int `gorm:"type:int(11)" json:"status"` // 
 Topped int8 `gorm:"type:tinyint(1)" json:"topped"` // 
 Commentable int8 `gorm:"type:tinyint(1)" json:"commentable"` // 
 ViewCount int `gorm:"type:int(11)" json:"view_count"` // 
 CommentCount int `gorm:"type:int(11)" json:"comment_count"` // 
 Ip string `gorm:"type:varchar(128)" json:"ip"` // 
 UserAgent string `gorm:"type:varchar(255)" json:"user_agent"` // 
 PushedAt time.Time `gorm:"type:timestamp" json:"pushed_at"` // 
 BlogID int `gorm:"type:bigint(20) unsigned" json:"blog_id"` // 
}
