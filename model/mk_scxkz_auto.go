package model

import (
	"time"
)

type MkScxkz struct {
	ScxkzID  string    `sql:"index" gorm:"type:varchar(36);primary_key" json:"scxkz_id"` //
	MkjbxxID string    `gorm:"type:varchar(36)" json:"mkjbxx_id"`                        //
	Zh       string    `gorm:"type:varchar(50)" json:"zh"`                               //
	Yxq      string    `gorm:"type:varchar(50)" json:"yxq"`                              //
	Cspfwh   string    `gorm:"type:varchar(50)" json:"cspfwh"`                           //
	Ztgcyswh string    `gorm:"type:varchar(50)" json:"ztgcyswh"`                         //
	Xkscnl   string    `gorm:"type:varchar(50)" json:"xkscnl"`                           //
	Xkkcmc   string    `gorm:"type:varchar(50)" json:"xkkcmc"`                           //
	Bzjg     string    `gorm:"type:varchar(50)" json:"bzjg"`                             //
	Bzsj     time.Time `gorm:"type:datetime" json:"bzsj"`                                //
}
