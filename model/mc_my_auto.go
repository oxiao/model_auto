package model

import (
	"time"
)

type McMy struct {
Model
 McID string `gorm:"type:varchar(36)" json:"mc_id"` // 
 MyID string `gorm:"type:varchar(36);unique;unique_index;not null" json:"my_id"` // 
 MyBm string `gorm:"type:varchar(36)" json:"my_bm"` // 
 MyMc string `gorm:"type:varchar(50)" json:"my_mc"` // 
 QyDd string `gorm:"type:varchar(50)" json:"qy_dd"` // 
 QySj time.Time `gorm:"type:datetime" json:"qy_sj"` // 
 Qyr string `gorm:"type:varchar(36)" json:"qyr"` // 
}
