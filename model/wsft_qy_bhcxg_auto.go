package model

import (
	"time"
)

type WsftQyBhcxg struct {
Model
 XgID string `gorm:"type:varchar(36);unique;unique_index;not null" json:"xg_id"` // f_id
 FtqyID string `gorm:"type:varchar(36)" json:"ftqy_id"` // f_id
 BfcID string `gorm:"type:varchar(36)" json:"bfc_id"` // 
 Kccs string `gorm:"type:varchar(50)" json:"kccs"` // 考察参数
 Wsyl decimal.Decimal `gorm:"type:decimal(18,3)" json:"wsyl"` // 瓦斯浓度
 Wshl decimal.Decimal `gorm:"type:decimal(18,3)" json:"wshl"` // 混合流量
 Zkwsll decimal.Decimal `gorm:"type:decimal(18,3)" json:"zkwsll"` // 纯流量
 Tqxxs decimal.Decimal `gorm:"type:decimal(18,3)" json:"tqxxs"` // 负压
 Bxcs decimal.Decimal `gorm:"type:decimal(18,3)" json:"bxcs"` // 温度
 Kcdd string `gorm:"type:varchar(50)" json:"kcdd"` // 考察地点
 Kcsj time.Time `gorm:"type:datetime" json:"kcsj"` // 考察时间
 Kcry string `gorm:"type:varchar(50)" json:"kcry"` // 考察人员
 Kcdw string `gorm:"type:varchar(50)" json:"kcdw"` // 考察单位
 Jlsj time.Time `gorm:"type:datetime" json:"jlsj"` // 记录时间
}
