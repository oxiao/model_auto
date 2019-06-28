package model

import (
	"time"
)

type MksbSb struct {
	SbID   string    `sql:"index" gorm:"type:varchar(36);primary_key" json:"sb_id"` //
	XsbID  string    `gorm:"type:varchar(36)" json:"xsb_id"`                        //
	SbglID string    `gorm:"type:varchar(36)" json:"sbgl_id"`                       //
	Fjsb   string    `gorm:"type:varchar(36)" json:"fjsb"`                          //
	Sbgl   string    `gorm:"type:varchar(36)" json:"sbgl"`                          // 井下大型设备，井下便携设备
	Zyfl   string    `gorm:"type:varchar(3)" json:"zyfl"`                           //
	Sblx   string    `gorm:"type:varchar(36)" json:"sblx"`                          //
	McID   string    `gorm:"type:varchar(36)" json:"mc_id"`                         //
	SbqyID string    `gorm:"type:varchar(36)" json:"sbqy_id"`                       //
	Sbbm   string    `gorm:"type:varchar(50)" json:"sbbm"`                          //
	Sbmc   string    `gorm:"type:varchar(50)" json:"sbmc"`                          //
	Sbsm   string    `gorm:"type:varchar(50)" json:"sbsm"`                          //
	Sbwz   string    `gorm:"type:varchar(50)" json:"sbwz"`                          //
	Kjzb   string    `gorm:"type:varchar(50)" json:"kjzb"`                          //
	Cdh    string    `gorm:"type:varchar(50)" json:"cdh"`                           //
	Azdd   string    `gorm:"type:varchar(50)" json:"azdd"`                          //
	Azry   string    `gorm:"type:varchar(50)" json:"azry"`                          //
	Azsj   time.Time `gorm:"type:datetime" json:"azsj"`                             //
	Yt     string    `gorm:"type:varchar(200)" json:"yt"`                           // 用途
	Whzq   string    `gorm:"type:varchar(200)" json:"whzq"`                         // 维护周期
	Zjjxsj time.Time `gorm:"type:datetime" json:"zjjxsj"`                           // 最近检修时间
}
