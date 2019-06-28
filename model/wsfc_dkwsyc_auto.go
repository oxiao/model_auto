package model

import (
	"time"
)

type WsfcDkwsyc struct {
	YcID string    `sql:"index" gorm:"type:varchar(36);primary_key" json:"yc_id"` // f_id
	Sj   time.Time `gorm:"type:datetime" json:"sj"`                               // 时间
	McID string    `gorm:"type:varchar(20)" json:"mc_id"`                         // 煤层
	Dd   string    `gorm:"type:varchar(50)" json:"dd"`                            // 地点
	Xx   string    `gorm:"type:varchar(500)" json:"xx"`                           // 异常现象
}
