package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type JxgcJjgzmJc struct {
	JcID    string          `sql:"index" gorm:"type:varchar(36);primary_key" json:"jc_id"` // f_id
	JjgzmID string          `gorm:"type:varchar(36)" json:"jjgzm_id"`                      //
	Jjm     string          `gorm:"type:varchar(50)" json:"jjm"`                           // 掘进面名称
	Jjrq    time.Time       `gorm:"type:datetime" json:"jjrq"`                             // 掘进日期
	Jcbc    string          `gorm:"type:varchar(20)" json:"jcbc"`                          // 掘进班次
	Jjdz    string          `gorm:"type:varchar(20)" json:"jjdz"`                          // 掘进对组
	Jjjc    decimal.Decimal `gorm:"type:decimal(18,3)" json:"jjjc"`                        // 掘进进尺数
	Cqrs    int             `gorm:"type:int(11)" json:"cqrs"`                              // 出勤人数
	Dbgx    string          `gorm:"type:varchar(20)" json:"dbgx"`                          // 当班工序
	Xbgx    string          `gorm:"type:varchar(20)" json:"xbgx"`                          // 下班工序
	Jlry    string          `gorm:"type:varchar(50)" json:"jlry"`                          // 记录人编码
	Bz      string          `gorm:"type:varchar(100)" json:"bz"`                           // 备注
}
