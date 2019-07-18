package model

import (
	"time"
)

type BmEmployeeOrg struct {
Model
 EmployeeOrgID string `gorm:"type:varchar(36);unique;unique_index;not null" json:"employee_org_id"` // 人员部门编号
 EmployeeID string `gorm:"type:varchar(36)" json:"employee_id"` // 员工编号
 CompanyID string `gorm:"type:varchar(36)" json:"company_id"` // 公司编号
 Idx int `gorm:"type:int(11)" json:"idx"` // 排序
 Rylx string `gorm:"type:varchar(3)" json:"rylx"` // 1主管，2副主管，3一般人员
}
