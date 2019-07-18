package model

import (
	"time"
)

type BmSessionToken struct {
Model
 EmployeeID string `gorm:"type:varchar(36);unique;unique_index;not null" json:"employee_id"` // 员工编号
 Token string `gorm:"type:varchar(1024)" json:"token"` // 会话令牌
 ExpireDate time.Time `gorm:"type:datetime" json:"expire_date"` // 会话有效时间
}
