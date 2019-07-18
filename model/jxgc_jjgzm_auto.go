package model

import (
	"time"
)

type JxgcJjgzm struct {
Model
 JjgzmID string `gorm:"type:varchar(36);unique;unique_index;not null" json:"jjgzm_id"` // 
 XddmID string `gorm:"type:varchar(36)" json:"xddm_id"` // 
 CqID string `gorm:"type:varchar(36)" json:"cq_id"` // 
 McID string `gorm:"type:varchar(36)" json:"mc_id"` // 
 Mc string `gorm:"type:varchar(50)" json:"mc"` // 
 ID string `gorm:"type:varchar(50)" json:"_id"` // 
 Lb string `gorm:"type:varchar(50)" json:"lb"` // 
 Zy string `gorm:"type:varchar(50)" json:"zy"` // 
 Zh string `gorm:"type:varchar(50)" json:"zh"` // 
 Jjff string `gorm:"type:varchar(50)" json:"jjff"` // 
 Jjsd decimal.Decimal `gorm:"type:decimal(13,3)" json:"jjsd"` // 
 Jjsb string `gorm:"type:varchar(50)" json:"jjsb"` // 
 Gfl decimal.Decimal `gorm:"type:decimal(13,3)" json:"gfl"` // 
 Cd decimal.Decimal `gorm:"type:decimal(13,3)" json:"cd"` // 
 Kssj time.Time `gorm:"type:datetime" json:"kssj"` // 
 Sfft string `gorm:"type:varchar(3)" json:"sfft"` // 
 Sfjw string `gorm:"type:varchar(3)" json:"sfjw"` // 
 JjdBh string `gorm:"type:varchar(36)" json:"jjd_bh"` // 
 SjxdBh string `gorm:"type:varchar(36)" json:"sjxd_bh"` // 
 SjxdMc string `gorm:"type:varchar(50)" json:"sjxd_mc"` // 
 DwQd string `gorm:"type:varchar(50)" json:"dw_qd"` // 
 DwFxd string `gorm:"type:varchar(50)" json:"dw_fxd"` // 
 DwZl string `gorm:"type:varchar(3)" json:"dw_zl"` // 
 DwJl decimal.Decimal `gorm:"type:decimal(13,3)" json:"dw_jl"` // 
 DwQsd string `gorm:"type:varchar(50)" json:"dw_qsd"` // 
 DwDxd decimal.Decimal `gorm:"type:decimal(13,3)" json:"dw_dxd"` // 
 XdFw string `gorm:"type:varchar(3)" json:"xd_fw"` // 
 XdPd decimal.Decimal `gorm:"type:decimal(13,3)" json:"xd_pd"` // 
 Rxhs decimal.Decimal `gorm:"type:decimal(13,3)" json:"rxhs"` // 
 Xhjc decimal.Decimal `gorm:"type:decimal(13,3)" json:"xhjc"` // 
 Xhl decimal.Decimal `gorm:"type:decimal(13,3)" json:"xhl"` // 
 Yjc decimal.Decimal `gorm:"type:decimal(13,3)" json:"yjc"` // 
 Sgq int `gorm:"type:int(11)" json:"sgq"` // 
 Edxs decimal.Decimal `gorm:"type:decimal(13,3)" json:"edxs"` // 
 Cqrs int `gorm:"type:int(11)" json:"cqrs"` // 
 Clxh decimal.Decimal `gorm:"type:decimal(13,3)" json:"clxh"` // 
}
