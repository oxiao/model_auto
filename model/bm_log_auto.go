package model

import (
	"time"
)

type BmLog struct {
Model
 LogID string `gorm:"type:varchar(36);unique;unique_index;not null" json:"log_id"` // 日志编号
 Type string `gorm:"type:varchar(36)" json:"type"` // 
 Grade string `gorm:"type:varchar(36)" json:"grade"` // INFO,DEBUG,WARN,ERROR
 Mis string `gorm:"type:varchar(128)" json:"mis"` // 
 Caller string `gorm:"type:varchar(512)" json:"caller"` // 
 Content string `gorm:"type:text" json:"content"` // 内容
 CreateDate time.Time `gorm:"type:datetime" json:"create_date"` // 
 Creator string `gorm:"type:varchar(36)" json:"creator"` // 操作人
 Ip string `gorm:"type:varchar(36)" json:"ip"` // 
}
