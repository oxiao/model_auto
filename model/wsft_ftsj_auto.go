package model

import (
	"time"
)

type WsftFtsj struct {
Model
 SjID string `gorm:"type:varchar(36);unique;unique_index;not null" json:"sj_id"` // f_id
 Sjmc string `gorm:"type:varchar(200)" json:"sjmc"` // 名称
 Sjfl string `gorm:"type:varchar(36)" json:"sjfl"` // 专项设计，措施计划，应急预案
 Sjlx string `gorm:"type:varchar(200)" json:"sjlx"` // 对应每个分类下的类型
 Dd string `gorm:"type:varchar(200)" json:"dd"` // 地点
 Nr string `gorm:"type:varchar(500)" json:"nr"` // 内容
 Zy string `gorm:"type:varchar(200)" json:"zy"` // 摘要
 Sjr string `gorm:"type:varchar(200)" json:"sjr"` // 设计人
 Bzsj time.Time `gorm:"type:datetime" json:"bzsj"` // 编制时间
 Jlsj time.Time `gorm:"type:datetime" json:"jlsj"` // 记录时间
 Wj string `gorm:"type:varchar(500)" json:"wj"` // 文件
}
