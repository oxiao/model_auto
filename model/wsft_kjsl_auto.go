package model

import "github.com/shopspring/decimal"

type WsftKjsl struct {
	SlID  string          `sql:"index" gorm:"type:varchar(36);primary_key" json:"sl_id"` // f_id
	Tjaq  string          `gorm:"type:varchar(50)" json:"tjaq"`                          // 统计期间
	Cqmc  string          `gorm:"type:varchar(50)" json:"cqmc"`                          // 采区名称
	Gzmmc string          `gorm:"type:varchar(50)" json:"gzmmc"`                         // 工作面名称
	Bg    decimal.Decimal `gorm:"type:decimal(18,3)" json:"bg"`                          // 标高
	Zxcd  decimal.Decimal `gorm:"type:decimal(18,3)" json:"zxcd"`                        // 走向长度（m）
	Qxcd  decimal.Decimal `gorm:"type:decimal(18,3)" json:"qxcd"`                        // 倾向长度（m）
	Qj    decimal.Decimal `gorm:"type:decimal(18,3)" json:"qj"`                          // 倾角（°）
	Pjmh  decimal.Decimal `gorm:"type:decimal(18,3)" json:"pjmh"`                        // 平均煤厚（m）
	Rz    decimal.Decimal `gorm:"type:decimal(18,3)" json:"rz"`                          // 容重（t/m）
	Dsl   decimal.Decimal `gorm:"type:decimal(18,3)" json:"dsl"`                         // 地损量（t）
	Dzl   decimal.Decimal `gorm:"type:decimal(18,3)" json:"dzl"`                         // 呆滞量（t）
	Hcl   decimal.Decimal `gorm:"type:decimal(18,3)" json:"hcl"`                         // 回采率
	Ktml  decimal.Decimal `gorm:"type:decimal(18,3)" json:"ktml"`                        // 开拓煤量（t）
	Zbml  decimal.Decimal `gorm:"type:decimal(18,3)" json:"zbml"`                        // 准备煤量（t）
	Hcml  decimal.Decimal `gorm:"type:decimal(18,3)" json:"hcml"`                        // 回采煤量（t）
}
