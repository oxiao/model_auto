package model

import (
	"time"
)

type BmEmployeeRole struct {
Model
 EmployeeRoleID string `gorm:"type:varchar(36);unique;unique_index;not null" json:"employee_role_id"` // 人员角色编号
 RoleID string `gorm:"type:varchar(36)" json:"role_id"` // 角色编号
 EmployeeID string `gorm:"type:varchar(36)" json:"employee_id"` // 员工编号
 MisID string `gorm:"type:varchar(64)" json:"mis_id"` // 系统编号
}
