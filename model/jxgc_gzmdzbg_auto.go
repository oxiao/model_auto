package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type JxgcGzmdzbg struct {
	BgID   string          `sql:"index" gorm:"type:varchar(36);primary_key" json:"bg_id"` //
	GzmID  string          `gorm:"type:varchar(200)" json:"gzm_id"`                       // 工作面编号
	GzmMc  string          `gorm:"type:varchar(50)" json:"gzm_mc"`                        //
	Gzrq   time.Time       `gorm:"type:datetime" json:"gzrq"`                             //
	Gzbc   string          `gorm:"type:varchar(6)" json:"gzbc"`                           //
	DzgzZk string          `gorm:"type:varchar(500)" json:"dzgz_zk"`                      //
	WsZk   string          `gorm:"type:varchar(500)" json:"ws_zk"`                        //
	McZk   string          `gorm:"type:varchar(500)" json:"mc_zk"`                        //
	GzmMh  decimal.Decimal `gorm:"type:decimal(13,3)" json:"gzm_mh"`                      //
	Rfchd  decimal.Decimal `gorm:"type:decimal(13,3)" json:"rfchd"`                       //
	Qtqk   string          `gorm:"type:varchar(1000)" json:"qtqk"`                        //
	Bgfj   string          `gorm:"type:varchar(100)" json:"bgfj"`                         //
	Jlrq   string          `gorm:"type:varchar(36)" json:"jlrq"`                          //
	Jlsj   time.Time       `gorm:"type:datetime" json:"jlsj"`                             //
}
