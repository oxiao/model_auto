package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type JxgcJxgzmJc struct {
	JcID  string          `sql:"index" gorm:"type:varchar(36);primary_key" json:"jc_id"` // f_id
	GzmID string          `gorm:"type:varchar(36)" json:"gzm_id"`                        // f_id
	Jjm   string          `gorm:"type:varchar(50)" json:"jjm"`                           // 掘进面名称
	Jjrq  time.Time       `gorm:"type:datetime" json:"jjrq"`                             // 掘进日期
	Jjbc  string          `gorm:"type:varchar(20)" json:"jjbc"`                          // 掘进班次
	Jjdz  string          `gorm:"type:varchar(20)" json:"jjdz"`                          // 掘进对组
	Jjjc  decimal.Decimal `gorm:"type:decimal(18,3)" json:"jjjc"`                        // 掘进进尺数
	Ljjc  decimal.Decimal `gorm:"type:decimal(18,3)" json:"ljjc"`                        // 累计进尺数
	Jmcfx decimal.Decimal `gorm:"type:decimal(18,3)" json:"jmcfx"`                       // 距煤层法线距离（m）
	Cqrs  int             `gorm:"type:int(11)" json:"cqrs"`                              // 出勤人数
	Dbgx  string          `gorm:"type:varchar(20)" json:"dbgx"`                          // 当班工序
	Xbgx  string          `gorm:"type:varchar(20)" json:"xbgx"`                          // 下班工序
	Jlry  string          `gorm:"type:varchar(20)" json:"jlry"`                          // 记录人员
	Jlsj  time.Time       `gorm:"type:datetime" json:"jlsj"`                             // 记录时间
	Bz    string          `gorm:"type:varchar(200)" json:"bz"`                           // 备注
}
