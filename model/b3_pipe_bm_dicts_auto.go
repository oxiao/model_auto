package model

import (
	"time"
)

type B3PipeBmDict struct {
Model
 DictID string `gorm:"type:varchar(36);unique;unique_index;not null" json:"dict_id"` // 
 ParentDictID string `gorm:"type:varchar(36)" json:"parent_dict_id"` // 
 Category string `gorm:"type:varchar(64)" json:"category"` // 
 Key string `gorm:"type:varchar(64)" json:"key"` // 
 Name string `gorm:"type:varchar(64)" json:"name"` // 
 Value string `gorm:"type:varchar(4000)" json:"value"` // 
 Value1 string `gorm:"type:varchar(64)" json:"value1"` // 
 Value2 string `gorm:"type:varchar(64)" json:"value2"` // 
 Value3 string `gorm:"type:varchar(64)" json:"value3"` // 
 Value4 string `gorm:"type:varchar(64)" json:"value4"` // 
 Format string `gorm:"type:varchar(64)" json:"format"` // 
 Idx int `gorm:"type:int(11)" json:"idx"` // 
 Note string `gorm:"type:varchar(256)" json:"note"` // 
 Editable string `gorm:"type:varchar(3)" json:"editable"` // 
 IsSys string `gorm:"type:varchar(3)" json:"is_sys"` // 
 User string `gorm:"type:varchar(1000)" json:"user"` // 
 Node string `gorm:"type:varchar(255)" json:"node"` // 
 State string `gorm:"type:varchar(3)" json:"state"` // 
 StateTime time.Time `gorm:"type:datetime" json:"state_time"` // 
}
