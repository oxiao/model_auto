package model

import (
	"time"
)

type WsftJbYj struct {
Model
 XjID string `gorm:"type:varchar(36);unique;unique_index;not null" json:"xj_id"` // f_id
 GzmID string `gorm:"type:varchar(36)" json:"gzm_id"` // 
 Gzmmc string `gorm:"type:varchar(50)" json:"gzmmc"` // 工作面名称
 Sj time.Time `gorm:"type:datetime" json:"sj"` // 时间
 Bc string `gorm:"type:varchar(50)" json:"bc"` // 班次
 S string `gorm:"type:varchar(50)" json:"s"` // 参数-s
 DeltaH2 string `gorm:"type:varchar(50)" json:"delta_h2"` // 参数-⊿h2
 K1 string `gorm:"type:varchar(50)" json:"k1"` // 参数-K1
 Q string `gorm:"type:varchar(50)" json:"q"` // 参数-q
 Qtcs string `gorm:"type:varchar(50)" json:"qtcs"` // 参数-其他
 Sd decimal.Decimal `gorm:"type:decimal(18,3)" json:"sd"` // 深度
 Cdry string `gorm:"type:varchar(36)" json:"cdry"` // 测定人员
 Jlsj time.Time `gorm:"type:datetime" json:"jlsj"` // 记录时间
 Splx string `gorm:"type:varchar(50)" json:"splx"` // 审批类型
}
