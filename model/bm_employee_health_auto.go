package model

import (
	"time"
)

type BmEmployeeHealth struct {
	HealthID      string    `sql:"index" gorm:"type:varchar(36);primary_key" json:"health_id"` //
	EmployeeID    string    `gorm:"type:varchar(36)" json:"employee_id"`                       // 员工编号
	ExameDate     time.Time `gorm:"type:datetime" json:"exame_date"`                           //
	ExameHospital string    `gorm:"type:varchar(36)" json:"exame_hospital"`                    //
	ReportSn      string    `gorm:"type:varchar(36)" json:"report_sn"`                         //
	ReportFile    string    `gorm:"type:varchar(256)" json:"report_file"`                      //
	Desc          string    `gorm:"type:varchar(1000)" json:"desc"`                            //
	State         string    `gorm:"type:varchar(3)" json:"state"`                              // 使用标志
	Idx           string    `gorm:"type:varchar(3)" json:"idx"`                                // 排序
	StateTime     time.Time `gorm:"type:datetime" json:"state_time"`                           // 标志更改时的时间
}
