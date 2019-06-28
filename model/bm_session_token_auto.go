package model

import (
	"time"
)

type BmSessionToken struct {
	EmployeeID string    `sql:"index" gorm:"type:varchar(36);primary_key" json:"employee_id"` // 员工编号
	Token      string    `gorm:"type:varchar(1024)" json:"token"`                             // 会话令牌
	ExpireDate time.Time `gorm:"type:datetime" json:"expire_date"`                            // 会话有效时间
}
