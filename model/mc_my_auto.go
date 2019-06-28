package model

import (
	"time"
)

type McMy struct {
	McID string    `gorm:"type:varchar(36)" json:"mc_id"`                         //
	MyID string    `sql:"index" gorm:"type:varchar(36);primary_key" json:"my_id"` //
	MyBm string    `gorm:"type:varchar(36)" json:"my_bm"`                         //
	MyMc string    `gorm:"type:varchar(50)" json:"my_mc"`                         //
	QyDd string    `gorm:"type:varchar(50)" json:"qy_dd"`                         //
	QySj time.Time `gorm:"type:datetime" json:"qy_sj"`                            //
	Qyr  string    `gorm:"type:varchar(36)" json:"qyr"`                           //
}
