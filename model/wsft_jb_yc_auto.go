package model

import (
	"time"
)

type WsftJbYc struct {
Model
 YcID string `gorm:"type:varchar(36);unique;unique_index;not null" json:"yc_id"` // f_id
 Gzmbh string `gorm:"type:varchar(36)" json:"gzmbh"` // 
 Gzmmc string `gorm:"type:varchar(50)" json:"gzmmc"` // 工作面名称
 Xhbh string `gorm:"type:varchar(50)" json:"xhbh"` // 循环编号
 Sgrq time.Time `gorm:"type:datetime" json:"sgrq"` // 施工日期
 Sgbc string `gorm:"type:varchar(50)" json:"sgbc"` // 施工班次
 Sgjd string `gorm:"type:varchar(50)" json:"sgjd"` // 施工进度
 JccsKh string `gorm:"type:varchar(50)" json:"jccs_kh"` // 检测参数-孔号
 JccsKs string `gorm:"type:varchar(50)" json:"jccs_ks"` // 检测参数-孔深
 JccsCs string `gorm:"type:varchar(50)" json:"jccs_cs"` // 检测参数-参数
 JccsZ string `gorm:"type:varchar(50)" json:"jccs_z"` // 检测参数-值
 Cdry string `gorm:"type:varchar(50)" json:"cdry"` // 测定人员
 Jlry string `gorm:"type:varchar(50)" json:"jlry"` // 记录人员
 Jlsj time.Time `gorm:"type:datetime" json:"jlsj"` // 记录时间
 Mchd decimal.Decimal `gorm:"type:decimal(18,3)" json:"mchd"` // 煤层厚度
 Rmhd decimal.Decimal `gorm:"type:decimal(18,3)" json:"rmhd"` // 软煤厚度
 Gzphd string `gorm:"type:varchar(50)" json:"gzphd"` // 动力-构造破坏带
 Fchd string `gorm:"type:varchar(50)" json:"fchd"` // 煤层赋存厚度变化
 Fcqj string `gorm:"type:varchar(50)" json:"fcqj"` // 煤层赋存倾角变化
 Fcdb string `gorm:"type:varchar(50)" json:"fcdb"` // 煤层赋存底板起伏
 Yldj string `gorm:"type:varchar(50)" json:"yldj"` // 应力叠加
 TcyzPk string `gorm:"type:varchar(50)" json:"tcyz_pk"` // 突出预兆-喷孔
 TcyzXmp string `gorm:"type:varchar(50)" json:"tcyz_xmp"` // 突出预兆-响煤炮
 TcyzWd string `gorm:"type:varchar(50)" json:"tcyz_wd"` // 突出预兆-温度变化
 TcyzQt string `gorm:"type:varchar(50)" json:"tcyz_qt"` // 突出预兆-其他预兆
 Splx string `gorm:"type:varchar(50)" json:"splx"` // 审批类型
}
