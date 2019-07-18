package model

import (
	"time"
)

type BmEmployeeWork struct {
Model
 WorkID string `gorm:"type:varchar(36);unique;unique_index;not null" json:"work_id"` // f_id
 EmployeeID string `gorm:"type:varchar(36)" json:"employee_id"` // 员工编号
 Name string `gorm:"type:varchar(36)" json:"name"` // 姓名
 WorkerSn string `gorm:"type:varchar(36)" json:"worker_sn"` // 工号
 InDate time.Time `gorm:"type:datetime" json:"in_date"` // 下井时间
 Work string `gorm:"type:varchar(500)" json:"work"` // 作业内容
 Addr string `gorm:"type:varchar(100)" json:"addr"` // 作业地点
 OutDate time.Time `gorm:"type:datetime" json:"out_date"` // 上井时间
 Question string `gorm:"type:varchar(500)" json:"question"` // 异常记录
}
