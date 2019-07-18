package model

import (
	"time"
)

type BmFetchTask struct {
Model
 TaskID string `gorm:"type:varchar(36);unique;unique_index;not null" json:"task_id"` // 
 Category string `gorm:"type:varchar(36)" json:"category"` // 系统任务，采集任务
 JobType string `gorm:"type:varchar(50)" json:"job_type"` // 以CJ开始表示采集类任务,如CJ-WCF,CJ-OPC,CJ-TXT,分别对应不同的采集任务.
 Name string `gorm:"type:varchar(255)" json:"name"` // 
 Note string `gorm:"type:varchar(255)" json:"note"` // 
 Assembly string `gorm:"type:varchar(255)" json:"assembly"` // 
 Clazz string `gorm:"type:varchar(255)" json:"clazz"` // 
 JobArgs string `gorm:"type:varchar(255)" json:"job_args"` // 
 Content string `gorm:"type:text" json:"content"` // 
 Cron string `gorm:"type:varchar(255)" json:"cron"` // 
 CronNote string `gorm:"type:varchar(255)" json:"cron_note"` // 
 LastRuntime time.Time `gorm:"type:datetime" json:"last_runtime"` // 
 NextRuntime time.Time `gorm:"type:datetime" json:"next_runtime"` // 
 RunNow string `gorm:"type:varchar(3)" json:"run_now"` // 是否立即执行.0否,1是.
 RunCount int `gorm:"type:int(11)" json:"run_count"` // 
 State int `gorm:"type:int(11)" json:"state"` // 0-运行   2-停止
 Idx int `gorm:"type:int(11)" json:"idx"` // 
 Creator string `gorm:"type:varchar(36)" json:"creator"` // 
 CreatorName string `gorm:"type:varchar(80)" json:"creator_name"` // 
 CreateDate time.Time `gorm:"type:datetime" json:"create_date"` // 
 Modifier string `gorm:"type:varchar(36)" json:"modifier"` // 
 ModifierName string `gorm:"type:varchar(80)" json:"modifier_name"` // 
 ModifiedDate time.Time `gorm:"type:timestamp" json:"modified_date"` // 
 IsDelete int `gorm:"type:int(11)" json:"is_delete"` // 
}
