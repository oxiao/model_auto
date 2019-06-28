package model

import (
	"time"
)

type MksbWx struct {
	WxID string    `sql:"index" gorm:"type:varchar(36);primary_key" json:"wx_id"` //
	SbID string    `gorm:"type:varchar(36)" json:"sb_id"`                         //
	Gzlr string    `gorm:"type:varchar(500)" json:"gzlr"`                         // 故障内容
	Bxry string    `gorm:"type:varchar(50)" json:"bxry"`                          // 报修人员
	Bxsj time.Time `gorm:"type:datetime" json:"bxsj"`                             // 报修时间
	Wxdw string    `gorm:"type:varchar(50)" json:"wxdw"`                          // 维修单位
	Wxry string    `gorm:"type:varchar(36)" json:"wxry"`                          // 维修人员
	Wxfs string    `gorm:"type:varchar(36)" json:"wxfs"`                          // 维修方式
	Wxjg string    `gorm:"type:varchar(50)" json:"wxjg"`                          // 维修结果
	Wxzq string    `gorm:"type:varchar(10)" json:"wxzq"`                          // 维修周期
	Bz   string    `gorm:"type:varchar(500)" json:"bz"`                           // 备注
}
