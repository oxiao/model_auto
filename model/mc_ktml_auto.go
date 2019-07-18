package model

import (
	"time"
)

type McKtml struct {
Model
 McID string `gorm:"type:varchar(36)" json:"mc_id"` // 
 Zxcd decimal.Decimal `gorm:"type:decimal(13,3)" json:"zxcd"` // 
 CqPjxqc decimal.Decimal `gorm:"type:decimal(13,3)" json:"cq_pjxqc"` // 
 KtqMcpjhd decimal.Decimal `gorm:"type:decimal(13,3)" json:"ktq_mcpjhd"` // 
 Smd decimal.Decimal `gorm:"type:decimal(13,3)" json:"smd"` // 
 Dzss decimal.Decimal `gorm:"type:decimal(13,3)" json:"dzss"` // 
 Dzml decimal.Decimal `gorm:"type:decimal(13,3)" json:"dzml"` // 
 CqCcl decimal.Decimal `gorm:"type:decimal(13,3)" json:"cq_ccl"` // 
}
