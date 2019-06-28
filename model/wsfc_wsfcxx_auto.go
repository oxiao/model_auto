package model

import "github.com/shopspring/decimal"

type WsfcWsfcxx struct {
	WsfcxxID string          `sql:"index" gorm:"type:varchar(36);primary_key" json:"wsfcxx_id"` // 瓦斯赋存信息编号_wsfcxx_id
	McID     string          `gorm:"type:varchar(36)" json:"mc_id"`                             //
	YswshlSc decimal.Decimal `gorm:"type:decimal(13,3)" json:"yswshl_sc"`                       // 原始瓦斯含量-_YSWSHL-
	YswsylSc decimal.Decimal `gorm:"type:decimal(13,3)" json:"yswsyl_sc"`                       // 原始瓦斯压力_YSWSYL
	YswsylTs decimal.Decimal `gorm:"type:decimal(13,3)" json:"yswsyl_ts"`                       // 原始瓦斯压力_YSWSYL
	Wsdzgl   string          `gorm:"type:varchar(50)" json:"wsdzgl"`                            // 瓦斯地质规律_WSDZGL
	Mcwscs   string          `gorm:"type:varchar(50)" json:"mcwscs"`                            // 煤层瓦斯参数_MCWSCS
	Wsycyc   string          `gorm:"type:varchar(50)" json:"wsycyc"`                            // 瓦斯涌出预测_WSYCYC
	YswshlTs decimal.Decimal `gorm:"type:decimal(13,3)" json:"yswshl_ts"`                       // 推算_TS
}
