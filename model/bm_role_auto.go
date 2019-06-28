package model

type BmRole struct {
	RoleID       string `sql:"index" gorm:"type:varchar(36);primary_key" json:"role_id"` // 角色编号
	ParentRoleID string `gorm:"type:varchar(36)" json:"parent_role_id"`                  // 上级角色编号
	CompanyID    string `gorm:"type:varchar(36)" json:"company_id"`                      // 公司编号
	Name         string `gorm:"type:varchar(64)" json:"name"`                            // 角色名称
	State        int    `gorm:"type:int(11)" json:"state"`                               // 使用标志
	Cmt          string `gorm:"type:varchar(64)" json:"cmt"`                             // 备注
	IsSys        int    `gorm:"type:int(11)" json:"is_sys"`                              // 0 普通角色，1系統角色.仅对平台进行管理的角色。
	Idx          int    `gorm:"type:int(11)" json:"idx"`                                 // 排序
	Node         string `gorm:"type:varchar(255)" json:"node"`                           // 节点编号
}
