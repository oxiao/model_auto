package model

import (
	"time"
)

type WsftJbAqfh struct {
	FhbhID string    `sql:"index" gorm:"type:varchar(36);primary_key" json:"fhbh_id"` // f_id
	GzmID  string    `gorm:"type:varchar(36)" json:"gzm_id"`                          //
	Cjdd   string    `gorm:"type:varchar(200)" json:"cjdd"`                           // 采掘地点
	Fhcs   string    `gorm:"type:varchar(201)" json:"fhcs"`                           // 防护措施
	Pzjc   string    `gorm:"type:varchar(202)" json:"pzjc"`                           // 批准进尺
	Sjjc   string    `gorm:"type:varchar(203)" json:"sjjc"`                           // 实际进尺
	Sj     time.Time `gorm:"type:datetime" json:"sj"`                                 // 时间
	Bc     string    `gorm:"type:varchar(50)" json:"bc"`                              // 班次
	Jlry   string    `gorm:"type:varchar(50)" json:"jlry"`                            // 记录人员
	Jlsj   time.Time `gorm:"type:datetime" json:"jlsj"`                               // 记录时间
}
