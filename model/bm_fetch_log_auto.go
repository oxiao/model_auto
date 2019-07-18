package model

import (
	"time"
)

type BmFetchLog struct {
Model
 LogID string `gorm:"type:varchar(36);unique;unique_index;not null" json:"log_id"` // 
 TaskID string `gorm:"type:varchar(36)" json:"task_id"` // 
 Jobname string `gorm:"type:varchar(255)" json:"jobname"` // 
 Executiontime time.Time `gorm:"type:datetime" json:"executiontime"` // 
 Executionduration decimal.Decimal `gorm:"type:decimal(18,2)" json:"executionduration"` // 
 Createddatetime time.Time `gorm:"type:datetime" json:"createddatetime"` // 
 Runlog string `gorm:"type:text" json:"runlog"` // 
}
