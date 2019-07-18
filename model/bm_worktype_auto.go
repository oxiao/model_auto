package model

import (
	"time"
)

type BmWorktype struct {
Model
 WorktypeID string `gorm:"type:varchar(36);unique;unique_index;not null" json:"worktype_id"` // 
 CompanyID string `gorm:"type:varchar(36)" json:"company_id"` // 公司编号
 Type string `gorm:"type:varchar(36)" json:"type"` // 
 Name string `gorm:"type:varchar(50)" json:"name"` // 
 Duty string `gorm:"type:varchar(1000)" json:"duty"` // 
 Request string `gorm:"type:varchar(1000)" json:"request"` // 
 CertName string `gorm:"type:varchar(30)" json:"cert_name"` // 
 CertSn string `gorm:"type:varchar(36)" json:"cert_sn"` // 
 CertValidDate time.Time `gorm:"type:datetime" json:"cert_valid_date"` // 
 State string `gorm:"type:varchar(3)" json:"state"` // 使用标志
 Idx string `gorm:"type:varchar(3)" json:"idx"` // 排序
 StateTime time.Time `gorm:"type:datetime" json:"state_time"` // 标志更改时的时间
}
