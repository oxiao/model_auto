package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type JxgcCmgzmJc struct {
	JcBh    string          `sql:"index" gorm:"type:varchar(36);primary_key" json:"jc_bh"` // f_id
	CmgzmID string          `gorm:"type:varchar(36)" json:"cmgzm_id"`                      //
	Hcrq    time.Time       `gorm:"type:datetime" json:"hcrq"`                             // 回采日期
	Hcbc    string          `gorm:"type:varchar(20)" json:"hcbc"`                          // 回采班次
	Cmdz    string          `gorm:"type:varchar(20)" json:"cmdz"`                          // 采煤队组
	Ds      decimal.Decimal `gorm:"type:decimal(18,3)" json:"ds"`                          // 刀数
	Pjjc    decimal.Decimal `gorm:"type:decimal(18,3)" json:"pjjc"`                        // 平均进尺
	Ljjc    decimal.Decimal `gorm:"type:decimal(18,3)" json:"ljjc"`                        // 累计进尺
	Cg      decimal.Decimal `gorm:"type:decimal(18,3)" json:"cg"`                          // 采高
	Jttj    decimal.Decimal `gorm:"type:decimal(18,3)" json:"jttj"`                        // 机头推进
	Jwtj    decimal.Decimal `gorm:"type:decimal(18,3)" json:"jwtj"`                        // 机尾推进
	Cl      decimal.Decimal `gorm:"type:decimal(18,3)" json:"cl"`                          // 产量（吨）
	Dbgx    string          `gorm:"type:varchar(20)" json:"dbgx"`                          // 当班工序
	Xbgx    string          `gorm:"type:varchar(20)" json:"xbgx"`                          // 下班工序
	Jlry    string          `gorm:"type:varchar(20)" json:"jlry"`                          // 记录人编码
	Jlsj    time.Time       `gorm:"type:datetime" json:"jlsj"`                             // 记录时间
	Bz      string          `gorm:"type:varchar(200)" json:"bz"`                           // 备注
}
