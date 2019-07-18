package model

import (
	"time"
)

type BmDict struct {
Model
 DictID string `gorm:"type:varchar(36);unique;unique_index;not null" json:"dict_id"` // 流水号
 ParentDictID string `gorm:"type:varchar(36)" json:"parent_dict_id"` // 上级流水号
 Category string `gorm:"type:varchar(64)" json:"category"` // 字典分类
 Key string `gorm:"type:varchar(64)" json:"key"` // 
 Name string `gorm:"type:varchar(64)" json:"name"` // 字典名称
 Value string `gorm:"type:varchar(4000)" json:"value"` // 字典值
 Value1 string `gorm:"type:varchar(64)" json:"value1"` // 字典值1
 Value2 string `gorm:"type:varchar(64)" json:"value2"` // 字典值2
 Value3 string `gorm:"type:varchar(64)" json:"value3"` // 字典值3
 Value4 string `gorm:"type:varchar(64)" json:"value4"` // 字典值4
 Format string `gorm:"type:varchar(64)" json:"format"` // 字典格式
 Idx int `gorm:"type:int(11)" json:"idx"` // 排序
 Note string `gorm:"type:varchar(256)" json:"note"` // 备注
 Editable string `gorm:"type:varchar(3)" json:"editable"` // 
 IsSys string `gorm:"type:varchar(3)" json:"is_sys"` // 0系統，1用戶
 User string `gorm:"type:varchar(1000)" json:"user"` // 用于字典目录记录上。记录使用此字典目录的表、字段。            格式：表.字段，...
 Node string `gorm:"type:varchar(255)" json:"node"` // 节点编号
 State string `gorm:"type:varchar(3)" json:"state"` // 使用标志
 StateTime time.Time `gorm:"type:datetime" json:"state_time"` // 标志更改时的时间
}
