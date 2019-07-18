package model

import (
	"time"
)

type BmParm struct {
Model
 ParmID string `gorm:"type:varchar(36);unique;unique_index;not null" json:"parm_id"` // 参数编号
 MisID string `gorm:"type:varchar(64)" json:"mis_id"` // 系统编号
 Category string `gorm:"type:varchar(64)" json:"category"` // 参数分类
 Name string `gorm:"type:varchar(64)" json:"name"` // 参数名称
 DefaultValue string `gorm:"type:varchar(1024)" json:"default_value"` // 参数默认值
 Cs1 string `gorm:"type:varchar(64)" json:"cs1"` // 参数当前值
 Cs2 string `gorm:"type:varchar(64)" json:"cs2"` // 参数运行值
 Cs3 string `gorm:"type:varchar(64)" json:"cs3"` // 参数最小值
 Cs4 string `gorm:"type:varchar(64)" json:"cs4"` // 参数最大值
 Format string `gorm:"type:varchar(64)" json:"format"` // 参数格式
 Note string `gorm:"type:varchar(256)" json:"note"` // 备注
 Idx int `gorm:"type:int(11)" json:"idx"` // 排序
 State string `gorm:"type:varchar(3)" json:"state"` // 使用标志
 StateTime time.Time `gorm:"type:datetime" json:"state_time"` // 标志更改时的时间
}
