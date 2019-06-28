package model

import "github.com/shopspring/decimal"

type WsfcHcmyc struct {
	YcID    string          `sql:"index" gorm:"type:varchar(36);primary_key" json:"yc_id"` // f_id
	CqID    string          `gorm:"type:varchar(36)" json:"cq_id"`                         //
	Cq      string          `gorm:"type:varchar(50)" json:"cq"`                            // 采区
	McID    string          `gorm:"type:varchar(36)" json:"mc_id"`                         //
	Mc      string          `gorm:"type:varchar(20)" json:"mc"`                            // 煤层
	Yclx    string          `gorm:"type:varchar(36)" json:"yclx"`                          // 开采（围岩）层，邻近层
	Ycy     string          `gorm:"type:varchar(20)" json:"ycy"`                           // 涌出源
	Mh      decimal.Decimal `gorm:"type:decimal(18,3)" json:"mh"`                          // 煤厚（m)
	Wsyl    decimal.Decimal `gorm:"type:decimal(18,3)" json:"wsyl"`                        // 瓦斯压力
	Wshl    decimal.Decimal `gorm:"type:decimal(18,3)" json:"wshl"`                        // 瓦斯含量（m3/t）
	Wsycl   decimal.Decimal `gorm:"type:decimal(18,3)" json:"wsycl"`                       // 瓦斯涌出量（m3/t）
	Kcgd    decimal.Decimal `gorm:"type:decimal(18,3)" json:"kcgd"`                        // 开采高度(m)
	Pfl     decimal.Decimal `gorm:"type:decimal(18,3)" json:"pfl"`                         // 排放率
	Xdwsycl decimal.Decimal `gorm:"type:decimal(18,3)" json:"xdwsycl"`                     // 相对瓦斯涌出量（m3/t）
	Ccwsl   decimal.Decimal `gorm:"type:decimal(18,3)" json:"ccwsl"`                       // 残存瓦斯含量（m3/t）
	Ljclx   string          `gorm:"type:varchar(20)" json:"ljclx"`                         // 邻近层类型
	Ljcbh   string          `gorm:"type:varchar(20)" json:"ljcbh"`                         // 邻近层编号
	Jkccjl  decimal.Decimal `gorm:"type:decimal(18,3)" json:"jkccjl"`                      // 距开采层距离
	Ljcycl  decimal.Decimal `gorm:"type:decimal(18,3)" json:"ljcycl"`                      // 邻近层总涌出量（m3/t）
}
