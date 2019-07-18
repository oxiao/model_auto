package model

import (
	"time"
)

type JxgcJxgzm struct {
Model
 GzmID string `gorm:"type:varchar(36);unique;unique_index;not null" json:"gzm_id"` // f_id
 CqID string `gorm:"type:varchar(36)" json:"cq_id"` // 
 McID string `gorm:"type:varchar(36)" json:"mc_id"` // 
 Mc string `gorm:"type:varchar(50)" json:"mc"` // 工作面名称
 Lx string `gorm:"type:varchar(36)" json:"lx"` // 石门工作面，揭煤工作面
 Sgd string `gorm:"type:varchar(20)" json:"sgd"` // 施工队名称
 Yjmc string `gorm:"type:varchar(36)" json:"yjmc"` // 预揭煤层
 Jmfs string `gorm:"type:varchar(36)" json:"jmfs"` // 揭煤方式
 Jjfs string `gorm:"type:varchar(36)" json:"jjfs"` // 掘进方式
 Dmxx string `gorm:"type:varchar(36)" json:"dmxx"` // 断面信息
 Zhfs string `gorm:"type:varchar(36)" json:"zhfs"` // 支护方式
 Xdqj decimal.Decimal `gorm:"type:decimal(18,3)" json:"xdqj"` // 巷道倾角
 Xdfwj decimal.Decimal `gorm:"type:decimal(18,3)" json:"xdfwj"` // 巷道方位角
 Kjrq time.Time `gorm:"type:datetime" json:"kjrq"` // 开掘日期
 Sfjcmc int `gorm:"type:int(11)" json:"sfjcmc"` // 是否揭穿煤层
 Ttrq time.Time `gorm:"type:datetime" json:"ttrq"` // 停头日期
 Jmcfx decimal.Decimal `gorm:"type:decimal(18,3)" json:"jmcfx"` // 距煤层法线（m）
 Fhcs string `gorm:"type:varchar(200)" json:"fhcs"` // 安全防护措施
}
