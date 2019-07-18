package model

import (
	"time"
)

type BmEmployeeEquipment struct {
Model
 EquipmentID string `gorm:"type:varchar(36);unique;unique_index;not null" json:"equipment_id"` // 
 EmployeeID string `gorm:"type:varchar(36)" json:"employee_id"` // 员工编号
 CardSn string `gorm:"type:varchar(36)" json:"card_sn"` // 
 LightSn string `gorm:"type:varchar(36)" json:"light_sn"` // 
 Phone string `gorm:"type:varchar(36)" json:"phone"` // 
 SelfSaveSn string `gorm:"type:varchar(36)" json:"self_save_sn"` // 
 State string `gorm:"type:varchar(3)" json:"state"` // 
 StateTime time.Time `gorm:"type:datetime" json:"state_time"` // 标志更改时的时间
}
