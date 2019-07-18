package model

import (
	"time"
)

type B3PipeCorrelation struct {
Model
 ID int `gorm:"type:bigint(20) unsigned;unique;unique_index;not null" json:"id"` // 
 CreatedAt time.Time `gorm:"type:timestamp" json:"created_at"` // 
 UpdatedAt time.Time `gorm:"type:timestamp" json:"updated_at"` // 
 DeletedAt time.Time `gorm:"type:timestamp" json:"deleted_at"` // 
 Id1 int `gorm:"type:bigint(20) unsigned" json:"id1"` // 
 Id2 int `gorm:"type:bigint(20) unsigned" json:"id2"` // 
 Str1 string `gorm:"type:varchar(255)" json:"str1"` // 
 Str2 string `gorm:"type:varchar(255)" json:"str2"` // 
 Str3 string `gorm:"type:varchar(255)" json:"str3"` // 
 Str4 string `gorm:"type:varchar(255)" json:"str4"` // 
 Int1 int `gorm:"type:int(11)" json:"int1"` // 
 Int2 int `gorm:"type:int(11)" json:"int2"` // 
 Int3 int `gorm:"type:int(11)" json:"int3"` // 
 Int4 int `gorm:"type:int(11)" json:"int4"` // 
 Type int `gorm:"type:int(11)" json:"type"` // 
 BlogID int `gorm:"type:bigint(20) unsigned" json:"blog_id"` // 
}
