package model

import (
	"time"
)

type BmEmployeeSkill struct {
	SkillID    string    `sql:"index" gorm:"type:varchar(36);primary_key" json:"skill_id"` //
	EmployeeID string    `gorm:"type:varchar(36)" json:"employee_id"`                      // 员工编号
	SkillName  string    `gorm:"type:varchar(36)" json:"skill_name"`                       //
	CertSn     string    `gorm:"type:varchar(36)" json:"cert_sn"`                          //
	CertOrg    string    `gorm:"type:varchar(128)" json:"cert_org"`                        //
	CertLevel  string    `gorm:"type:varchar(36)" json:"cert_level"`                       //
	CertDate   time.Time `gorm:"type:datetime" json:"cert_date"`                           //
	ExpireDate time.Time `gorm:"type:datetime" json:"expire_date"`                         //
	CertType   string    `gorm:"type:varchar(36)" json:"cert_type"`                        // 培训类，鉴定类，资格类，等级类
	UsageLimit string    `gorm:"type:varchar(36)" json:"usage_limit"`                      // 国际，国内，省市，单位
	Note       string    `gorm:"type:varchar(1000)" json:"note"`                           //
	State      string    `gorm:"type:varchar(3)" json:"state"`                             // 使用标志
	Idx        string    `gorm:"type:varchar(3)" json:"idx"`                               // 排序
	StateTime  time.Time `gorm:"type:datetime" json:"state_time"`                          // 标志更改时的时间
}
