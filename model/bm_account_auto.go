package model

type BmAccount struct {
	EmployeeID string `gorm:"type:varchar(36)" json:"employee_id"` // 员工编号
	Code       string `gorm:"type:varchar(36)" json:"code"`        //
	Pwd        string `gorm:"type:varchar(100)" json:"pwd"`        //
	Security   string `gorm:"type:varchar(200)" json:"security"`   //
	LastLogin  string `gorm:"type:varchar(500)" json:"last_login"` //
}
