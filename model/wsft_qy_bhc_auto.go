package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type WsftQyBhc struct {
	FtqyID string          `sql:"index" gorm:"type:varchar(36);primary_key" json:"ftqy_id"` // f_id
	BfcID  string          `sql:"index" gorm:"type:varchar(36);primary_key" json:"bfc_id"`  //
	Lx     string          `gorm:"type:varchar(50)" json:"lx"`                              // 保护层类型
	Mc     string          `gorm:"type:varchar(50)" json:"mc"`                              // 保护层名称
	Tcwxx  string          `gorm:"type:varchar(50)" json:"tcwxx"`                           // 突出危险性
	Mh     decimal.Decimal `gorm:"type:decimal(13,3)" json:"mh"`                            //
	McID   string          `gorm:"type:varchar(50)" json:"mc_id"`                           // 保护层煤层
	Qj     decimal.Decimal `gorm:"type:decimal(18,3)" json:"qj"`                            // 保护层倾角
	Mz     string          `gorm:"type:varchar(50)" json:"mz"`                              // 保护层煤柱
	Jl     decimal.Decimal `gorm:"type:decimal(18,3)" json:"jl"`                            // 保护层距离
	Zdbhcj decimal.Decimal `gorm:"type:decimal(18,3)" json:"zdbhcj"`                        // 最大保护垂距
	Zxcjjl decimal.Decimal `gorm:"type:decimal(18,3)" json:"zxcjjl"`                        // 最小层间距离
	Xysbhj decimal.Decimal `gorm:"type:decimal(18,3)" json:"xysbhj"`                        // 泄压上保护角
	Xyxbhj decimal.Decimal `gorm:"type:decimal(18,3)" json:"xyxbhj"`                        // 泄压下保护角
	Cjcqsj time.Time       `gorm:"type:datetime" json:"cjcqsj"`                             // 采掘超前时间
	Zxcqj  decimal.Decimal `gorm:"type:decimal(18,3)" json:"zxcqj"`                         // 最小超前距
}
