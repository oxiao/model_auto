package model

import (
	"time"
)

type BmJobtitle struct {
Model
 JobtitleID string `gorm:"type:varchar(36);unique;unique_index;not null" json:"jobtitle_id"` // 
 CompanyID string `gorm:"type:varchar(36)" json:"company_id"` // 公司编号
 Type string `gorm:"type:varchar(36)" json:"type"` // 
 Name string `gorm:"type:varchar(50)" json:"name"` // 
 WorkContent string `gorm:"type:varchar(1000)" json:"work_content"` // 
 Level string `gorm:"type:varchar(36)" json:"level"` // 
 Idx string `gorm:"type:varchar(3)" json:"idx"` // 排序
 State string `gorm:"type:varchar(3)" json:"state"` // 使用标志
 StateTime time.Time `gorm:"type:datetime" json:"state_time"` // 标志更改时的时间
}
