package model

import (
	"time"
)

type BmTemplate struct {
Model
 TemplateID string `gorm:"type:varchar(36);unique;unique_index;not null" json:"template_id"` // 
 Name string `gorm:"type:varchar(36)" json:"name"` // 
 Code string `gorm:"type:varchar(36)" json:"code"` // 
 Content string `gorm:"type:text" json:"content"` // 
 Type string `gorm:"type:varchar(36)" json:"type"` // 采集，调度，报表，
 Ver string `gorm:"type:varchar(36)" json:"ver"` // 
 CreateDate time.Time `gorm:"type:datetime" json:"create_date"` // 
 Status string `gorm:"type:varchar(4)" json:"status"` // 
}
