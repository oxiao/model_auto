package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type WsftQyQyxx struct {
	FtqyID string          `sql:"index" gorm:"type:varchar(36);primary_key" json:"ftqy_id"` // f_id
	CqID   string          `gorm:"type:varchar(36)" json:"cq_id"`                           //
	McID   string          `gorm:"type:varchar(36)" json:"mc_id"`                           //
	Mcmc   string          `gorm:"type:varchar(50)" json:"mcmc"`                            // 煤层名称
	Cqmc   string          `gorm:"type:varchar(50)" json:"cqmc"`                            // 采区名称
	Ftcs   string          `gorm:"type:varchar(500)" json:"ftcs"`                           // 防突措施
	Kssj   time.Time       `gorm:"type:datetime" json:"kssj"`                               // 开始时间
	Jssj   time.Time       `gorm:"type:datetime" json:"jssj"`                               // 结束时间
	WshlCc decimal.Decimal `gorm:"type:decimal(13,3)" json:"wshl_cc"`                       // 瓦斯含量-_WSHL-
	WshlYs decimal.Decimal `gorm:"type:decimal(13,3)" json:"wshl_ys"`                       // 瓦斯含量-_WSHL-
	WsylYs decimal.Decimal `gorm:"type:decimal(13,3)" json:"wsyl_ys"`                       // 瓦斯压力-_WSYL-
	WsylCc string          `gorm:"type:varchar(50)" json:"wsyl_cc"`                         // 残存_CC
	Qtcs   string          `gorm:"type:varchar(500)" json:"qtcs"`                           // 其他瓦斯参数
}
