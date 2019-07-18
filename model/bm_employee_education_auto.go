package model

import (
	"time"
)

type BmEmployeeEducation struct {
Model
 EducationID string `gorm:"type:varchar(36);unique;unique_index;not null" json:"education_id"` // 
 EmployeeID string `gorm:"type:varchar(36)" json:"employee_id"` // 员工编号
 StartDate time.Time `gorm:"type:datetime" json:"start_date"` // 
 EndDate time.Time `gorm:"type:datetime" json:"end_date"` // 
 School string `gorm:"type:varchar(128)" json:"school"` // 
 Education string `gorm:"type:varchar(36)" json:"education"` // 
 Witness string `gorm:"type:varchar(36)" json:"witness"` // 
 Phone string `gorm:"type:varchar(36)" json:"phone"` // 
 CertFile string `gorm:"type:varchar(256)" json:"cert_file"` // 
 CertSn string `gorm:"type:varchar(36)" json:"cert_sn"` // 
 State string `gorm:"type:varchar(3)" json:"state"` // 使用标志
 Idx string `gorm:"type:varchar(3)" json:"idx"` // 排序
 StateTime time.Time `gorm:"type:datetime" json:"state_time"` // 标志更改时的时间
}
