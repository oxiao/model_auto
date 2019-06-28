package model

import (
	"time"
)

type BmOrg struct {
	CompanyID       string    `sql:"index" gorm:"type:varchar(36);primary_key" json:"company_id"` // 公司编号
	ParentCompanyID string    `gorm:"type:varchar(36)" json:"parent_company_id"`                  // 上级公司编号
	RegionID        string    `gorm:"type:varchar(36)" json:"region_id"`                          // 区划编号
	Name            string    `gorm:"type:varchar(512)" json:"name"`                              // 公司名称
	SimpleName      string    `gorm:"type:varchar(32)" json:"simple_name"`                        // 簡化名
	Code            string    `gorm:"type:varchar(64)" json:"code"`                               // 机构代码
	Owner           string    `gorm:"type:varchar(256)" json:"owner"`                             // 法人
	Qualification   string    `gorm:"type:varchar(128)" json:"qualification"`                     // 资质
	Capital         string    `gorm:"type:varchar(32)" json:"capital"`                            // 注册资金
	Address         string    `gorm:"type:varchar(512)" json:"address"`                           // 联系地址
	City            string    `gorm:"type:varchar(64)" json:"city"`                               // 城市
	Station         string    `gorm:"type:varchar(64)" json:"station"`                            // 州
	Country         string    `gorm:"type:varchar(64)" json:"country"`                            // 国家
	Phone           string    `gorm:"type:varchar(64)" json:"phone"`                              // 电话
	Fax             string    `gorm:"type:varchar(32)" json:"fax"`                                // 传真
	Photo           string    `gorm:"type:text" json:"photo"`                                     // 徽标
	WebSite         string    `gorm:"type:varchar(256)" json:"web_site"`                          // 网站
	Post            string    `gorm:"type:varchar(64)" json:"post"`                               // 邮编
	Industry        string    `gorm:"type:varchar(3)" json:"industry"`                            // 行业
	Orient          string    `gorm:"type:varchar(3)" json:"orient"`                              // 方向
	RegistryDate    time.Time `gorm:"type:datetime" json:"registry_date"`                         // 注册时间
	UnregistryDate  time.Time `gorm:"type:datetime" json:"unregistry_date"`                       // 撤消时间
	Note            string    `gorm:"type:varchar(256)" json:"note"`                              // 备注
	OrgType         string    `gorm:"type:varchar(3)" json:"org_type"`                            // 0 总公司 1 分公司 2 部门
	Type            string    `gorm:"type:varchar(3)" json:"type"`                                // 公司类型
	Idx             string    `gorm:"type:varchar(3)" json:"idx"`                                 // 排序
	Node            string    `gorm:"type:varchar(255)" json:"node"`                              // 节点编号
	State           string    `gorm:"type:varchar(3)" json:"state"`                               // 使用标志
	StateTime       time.Time `gorm:"type:datetime" json:"state_time"`                            // 标志更改时的时间
}
