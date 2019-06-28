package model

import (
	"time"
)

type BmEmployeeBank struct {
	BankID     string    `sql:"index" gorm:"type:varchar(36);primary_key" json:"bank_id"` //
	EmployeeID string    `gorm:"type:varchar(36)" json:"employee_id"`                     // 员工编号
	BankName   string    `gorm:"type:varchar(128)" json:"bank_name"`                      //
	Account    string    `gorm:"type:varchar(36)" json:"account"`                         //
	Usage      string    `gorm:"type:varchar(36)" json:"usage"`                           // 工资，报销，补助，奖金，其它
	State      string    `gorm:"type:varchar(3)" json:"state"`                            // 使用标志
	Idx        string    `gorm:"type:varchar(3)" json:"idx"`                              // 排序
	StateTime  time.Time `gorm:"type:datetime" json:"state_time"`                         // 标志更改时的时间
}
