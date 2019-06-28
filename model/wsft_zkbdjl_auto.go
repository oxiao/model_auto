package model

import (
	"time"
)

type WsftZkbdjl struct {
	ZkID string    `sql:"index" gorm:"type:varchar(36);primary_key" json:"zk_id"` // f_id
	Bh   string    `gorm:"type:varchar(36)" json:"bh"`                            //
	Lx   string    `gorm:"type:varchar(36)" json:"lx"`                            // 区域，局部
	Jd   string    `gorm:"type:varchar(36)" json:"jd"`                            // 预测，措施，检验
	Sg   string    `gorm:"type:varchar(36)" json:"sg"`                            // 设计，施工
	Jlrq time.Time `gorm:"type:datetime" json:"jlrq"`                             //
}
