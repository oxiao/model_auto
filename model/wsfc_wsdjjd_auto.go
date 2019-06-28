package model

import "github.com/shopspring/decimal"

type WsfcWsdjjd struct {
	JdID     string          `sql:"index" gorm:"type:varchar(36);primary_key" json:"jd_id"` // f_id
	Jdnf     string          `gorm:"type:char(4)" json:"jdnf"`                              // 鉴定年份
	Hdscnl   decimal.Decimal `gorm:"type:decimal(18,3)" json:"hdscnl"`                      // 核定生产能力
	Wsjdycl  decimal.Decimal `gorm:"type:decimal(18,3)" json:"wsjdycl"`                     // 瓦斯绝对涌出量（m3/min）
	Wsxdycl  decimal.Decimal `gorm:"type:decimal(18,3)" json:"wsxdycl"`                     // 瓦斯相对涌出量（m3/t）
	CmZdycl  decimal.Decimal `gorm:"type:decimal(18,3)" json:"cm_zdycl"`                    // 采面最大绝对瓦斯涌出量（m3/min）
	JmZdycl  decimal.Decimal `gorm:"type:decimal(18,3)" json:"jm_zdycl"`                    // 掘面最大绝对瓦斯涌出量（m3/min）
	Co2Jdycl decimal.Decimal `gorm:"type:decimal(18,3)" json:"co2_jdycl"`                   // 二氧化碳绝对涌出量（m3/min）
	Co2Xdycl decimal.Decimal `gorm:"type:decimal(18,5)" json:"co2_xdycl"`                   // 二氧化碳相对涌出量（m3/t）
	Ysjgzts  int             `gorm:"type:int(11)" json:"ysjgzts"`                           // 月实际工作天数d
	Ycml     decimal.Decimal `gorm:"type:decimal(18,3)" json:"ycml"`                        // 月产煤量t
	Fhq      decimal.Decimal `gorm:"type:decimal(18,3)" json:"fhq"`                         // 开采煤层最短发火期d
	Zrfhqx   string          `gorm:"type:varchar(20)" json:"zrfhqx"`                        // 自燃发火倾向性
	Mcbzx    string          `gorm:"type:varchar(20)" json:"mcbzx"`                         // 煤尘爆炸性
	Wsdlxx   string          `gorm:"type:varchar(500)" json:"wsdlxx"`                       // 瓦斯动力现象情况
	Wspcqk   string          `gorm:"type:varchar(500)" json:"wspcqk"`                       // 瓦斯喷出情况
	Jdwsdj   string          `gorm:"type:varchar(20)" json:"jdwsdj"`                        // 矿井鉴定瓦斯等级
	SndDj    string          `gorm:"type:varchar(20)" json:"snd_dj"`                        // 上年度矿井瓦斯等级
}
