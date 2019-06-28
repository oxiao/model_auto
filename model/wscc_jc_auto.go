package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type WsccJc struct {
	WsccxtID string          `sql:"index" gorm:"type:varchar(36);primary_key" json:"wsccxt_id"` //
	FtqyID   string          `gorm:"type:varchar(36)" json:"ftqy_id"`                           // f_id
	Ccdd     string          `gorm:"type:varchar(50)" json:"ccdd"`                              //
	Ccfs     string          `gorm:"type:varchar(50)" json:"ccfs"`                              //
	Ccsb     string          `gorm:"type:varchar(50)" json:"ccsb"`                              //
	Ccgx     string          `gorm:"type:varchar(50)" json:"ccgx"`                              //
	Jczk     string          `gorm:"type:varchar(200)" json:"jczk"`                             // 监测钻孔
	Kssj     time.Time       `gorm:"type:datetime" json:"kssj"`                                 // 开始时间
	Ljsc     decimal.Decimal `gorm:"type:decimal(18,3)" json:"ljsc"`                            // 累计时长
	Wsnd     decimal.Decimal `gorm:"type:decimal(18,3)" json:"wsnd"`                            // 瓦斯浓度
	Hhll     decimal.Decimal `gorm:"type:decimal(18,3)" json:"hhll"`                            // 混合流量
	Cll      decimal.Decimal `gorm:"type:decimal(18,3)" json:"cll"`                             // 纯流量
	Fy       decimal.Decimal `gorm:"type:decimal(18,3)" json:"fy"`                              // 负压
	Wd       decimal.Decimal `gorm:"type:decimal(18,3)" json:"wd"`                              // 温度
	Jcsj     time.Time       `gorm:"type:datetime" json:"jcsj"`                                 // 监测时间
}
