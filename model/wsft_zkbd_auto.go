package model

import (
	"time"
)

type WsftZkbd struct {
Model
 ZkID string `gorm:"type:varchar(36);unique;unique_index;not null" json:"zk_id"` // f_id
 Ddmc string `gorm:"type:varchar(50)" json:"ddmc"` // 地点名称
 Zklx string `gorm:"type:varchar(50)" json:"zklx"` // 钻孔类型
 Sbsj time.Time `gorm:"type:datetime" json:"sbsj"` // 施工时间
 Sgmc string `gorm:"type:varchar(50)" json:"sgmc"` // 施工班次
 KkcsJdb decimal.Decimal `gorm:"type:decimal(18,3)" json:"kkcs_jdb"` // 开孔参数
 KkcsJzx decimal.Decimal `gorm:"type:decimal(18,3)" json:"kkcs_jzx"` // 
 KkcsZkkj decimal.Decimal `gorm:"type:decimal(18,3)" json:"kkcs_zkkj"` // 
 ZkcsPj decimal.Decimal `gorm:"type:decimal(18,3)" json:"zkcs_pj"` // 
 ZkcsQj decimal.Decimal `gorm:"type:decimal(18,3)" json:"zkcs_qj"` // 
 ZkcsZksd decimal.Decimal `gorm:"type:decimal(18,3)" json:"zkcs_zksd"` // 
 DlxxWz decimal.Decimal `gorm:"type:decimal(18,3)" json:"dlxx_wz"` // 
 DlxxYclx string `gorm:"type:varchar(50)" json:"dlxx_yclx"` // 
 DlxxYzcd string `gorm:"type:varchar(50)" json:"dlxx_yzcd"` // 动力现象
 Xgdz string `gorm:"type:varchar(50)" json:"xgdz"` // 施工队组
 Jlry string `gorm:"type:varchar(50)" json:"jlry"` // 记录人员
 Jlsj time.Time `gorm:"type:datetime" json:"jlsj"` // 记录时间
 Xplx string `gorm:"type:varchar(50)" json:"xplx"` // 审批类型
}
