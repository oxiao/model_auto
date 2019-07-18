package model

import (
	"time"
)

type BmJob struct {
Model
 EmployeeID string `gorm:"type:varchar(36);unique;unique_index;not null" json:"employee_id"` // 员工编号
 Zy string `gorm:"type:varchar(20)" json:"zy"` // 专业
 Bz string `gorm:"type:varchar(20)" json:"bz"` // 班制
 Zcgzsj int `gorm:"type:int(11)" json:"zcgzsj"` // 最长工作时间
 Gdxjcs int `gorm:"type:int(11)" json:"gdxjcs"` // 规定下井次数
 Cyzgzbh string `gorm:"type:varchar(36)" json:"cyzgzbh"` // 从业资格证编号
}
