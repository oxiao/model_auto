package model

import (
	"time"
)

type BmEmployeeFamily struct {
Model
 FamilyID string `gorm:"type:varchar(36);unique;unique_index;not null" json:"family_id"` // 
 EmployeeID string `gorm:"type:varchar(36)" json:"employee_id"` // 员工编号
 Relation string `gorm:"type:varchar(36)" json:"relation"` // 
 Name string `gorm:"type:varchar(36)" json:"name"` // 
 Sex string `gorm:"type:varchar(36)" json:"sex"` // 
 Bod time.Time `gorm:"type:datetime" json:"bod"` // 
 Job string `gorm:"type:varchar(36)" json:"job"` // 
 Phone string `gorm:"type:varchar(36)" json:"phone"` // 
 Addr string `gorm:"type:varchar(256)" json:"addr"` // 
 State string `gorm:"type:varchar(3)" json:"state"` // 使用标志
 Idx string `gorm:"type:varchar(3)" json:"idx"` // 排序
 StateTime time.Time `gorm:"type:datetime" json:"state_time"` // 标志更改时的时间
}
