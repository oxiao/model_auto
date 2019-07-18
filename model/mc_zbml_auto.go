package model

import (
	"time"
)

type McZbml struct {
Model
 McID string `gorm:"type:varchar(36)" json:"mc_id"` // 
 CqID string `gorm:"type:varchar(36)" json:"cq_id"` // 
 TCMcID string `gorm:"type:varchar(36)" json:"t_c_mc_id"` // 
 CqZxcd decimal.Decimal `gorm:"type:decimal(13,3)" json:"cq_zxcd"` // 
 CqQxcd decimal.Decimal `gorm:"type:decimal(13,3)" json:"cq_qxcd"` // 
 Cqmcpjhd decimal.Decimal `gorm:"type:decimal(13,3)" json:"cqmcpjhd"` // 
 Smd decimal.Decimal `gorm:"type:decimal(13,3)" json:"smd"` // 
 Dzss decimal.Decimal `gorm:"type:decimal(13,3)" json:"dzss"` // 
 Dzml decimal.Decimal `gorm:"type:decimal(13,3)" json:"dzml"` // 
 CqCcl decimal.Decimal `gorm:"type:decimal(13,3)" json:"cq_ccl"` // 
}
