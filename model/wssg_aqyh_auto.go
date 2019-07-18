package model

import (
	"time"
)

type WssgAqyh struct {
Model
 YhID string `gorm:"type:varchar(36);unique;unique_index;not null" json:"yh_id"` // f_id
 FPlace string `gorm:"type:varchar(50)" json:"f_place"` // 地点
 Sj time.Time `gorm:"type:datetime" json:"sj"` // 时间
 Yhjb string `gorm:"type:varchar(50)" json:"yhjb"` // 隐患级别
 Yhnr string `gorm:"type:varchar(500)" json:"yhnr"` // 隐患内容
 Yhtp string `gorm:"type:varchar(500)" json:"yhtp"` // 隐患图片
 Jcbm string `gorm:"type:varchar(50)" json:"jcbm"` // 检查部门
 Jcry string `gorm:"type:varchar(50)" json:"jcry"` // 检查人员
 Clzt string `gorm:"type:varchar(50)" json:"clzt"` // 处理状态
 Clnr string `gorm:"type:varchar(50)" json:"clnr"` // 处理内容
 Clbm string `gorm:"type:varchar(50)" json:"clbm"` // 处理部门
 Clry string `gorm:"type:varchar(50)" json:"clry"` // 处理人员
 Clsj time.Time `gorm:"type:datetime" json:"clsj"` // 处理时间
 Jlsj time.Time `gorm:"type:datetime" json:"jlsj"` // 记录时间
}
