package model

import "github.com/shopspring/decimal"

type WsfcJjmyc struct {
	JjycID string          `sql:"index" gorm:"type:varchar(36);primary_key" json:"jjyc_id"` // f_id
	CqID   string          `gorm:"type:varchar(36)" json:"cq_id"`                           //
	Cq     string          `gorm:"type:varchar(20)" json:"cq"`                              // 生产采区
	McID   string          `gorm:"type:varchar(36)" json:"mc_id"`                           //
	Yclx   string          `gorm:"type:varchar(36)" json:"yclx"`                            // 煤壁瓦斯涌出,落煤瓦斯涌出
	Ycy    string          `gorm:"type:varchar(20)" json:"ycy"`                             // 涌出源
	Xdcd   decimal.Decimal `gorm:"type:decimal(18,3)" json:"xdcd"`                          // 巷道长度（m）
	Mbzc   decimal.Decimal `gorm:"type:decimal(18,3)" json:"mbzc"`                          // 煤壁周长(m)
	Jjsd   decimal.Decimal `gorm:"type:decimal(18,3)" json:"jjsd"`                          // 掘进速度(m/min)
	Wshl   decimal.Decimal `gorm:"type:decimal(18,2)" json:"wshl"`                          // 瓦斯含量（m3/t）
	Csycl  decimal.Decimal `gorm:"type:decimal(18,3)" json:"csycl"`                         // 煤壁初始涌出量
	Wsycl  decimal.Decimal `gorm:"type:decimal(18,2)" json:"wsycl"`                         // 煤壁瓦斯涌出量（m3/min）
	Smd    decimal.Decimal `gorm:"type:decimal(18,3)" json:"smd"`                           // 煤的视密度（t/m3）
	Xddm   decimal.Decimal `gorm:"type:decimal(18,3)" json:"xddm"`                          // 巷道断面（m2）
}
