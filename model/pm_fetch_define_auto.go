package model

import (
	"time"
)

type PmFetchDefine struct {
Model
 DefineID string `gorm:"type:varchar(36);unique;unique_index;not null" json:"define_id"` // 
 TaskID string `gorm:"type:varchar(36)" json:"task_id"` // 
 MonitorMis string `gorm:"type:varchar(36)" json:"monitor_mis"` // 
 Category string `gorm:"type:varchar(36)" json:"category"` // 
 Account string `gorm:"type:varchar(36)" json:"account"` // 
 Pwd string `gorm:"type:varchar(36)" json:"pwd"` // 
 FetchUrl string `gorm:"type:varchar(512)" json:"fetch_url"` // 
 FetchType string `gorm:"type:varchar(36)" json:"fetch_type"` // txt,ftp,wcf,sqlite,file,vedio
 DataMap string `gorm:"type:text" json:"data_map"` // 
}