package model

type BmEmployeeRole struct {
	EmployeeRoleID string `sql:"index" gorm:"type:varchar(36);primary_key" json:"employee_role_id"` // 人员角色编号
	RoleID         string `gorm:"type:varchar(36)" json:"role_id"`                                  // 角色编号
	EmployeeID     string `gorm:"type:varchar(36)" json:"employee_id"`                              // 员工编号
	MisID          string `gorm:"type:varchar(64)" json:"mis_id"`                                   // 系统编号
}
