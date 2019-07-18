package model

import (
	"time"
)

type JxgcCmgzmSc struct {
Model
 CmgzmID string `gorm:"type:varchar(36)" json:"cmgzm_id"` // 
 CmmMc string `gorm:"type:varchar(50)" json:"cmm_mc"` // 
 ScN string `gorm:"type:varchar(4)" json:"sc_n"` // 
 ScY string `gorm:"type:varchar(2)" json:"sc_y"` // 
 ScX string `gorm:"type:varchar(2)" json:"sc_x"` // 
 JxYsjc decimal.Decimal `gorm:"type:decimal(13,3)" json:"jx_ysjc"` // 
 JcSc decimal.Decimal `gorm:"type:decimal(13,3)" json:"jc_sc"` // 
 FxYsjc decimal.Decimal `gorm:"type:decimal(13,3)" json:"fx_ysjc"` // 
 FxSc decimal.Decimal `gorm:"type:decimal(13,3)" json:"fx_sc"` // 
 Sfxy string `gorm:"type:varchar(3)" json:"sfxy"` // 
 Kcmj decimal.Decimal `gorm:"type:decimal(13,3)" json:"kcmj"` // 
 Jcml decimal.Decimal `gorm:"type:decimal(13,3)" json:"jcml"` // 
 Fdmxs decimal.Decimal `gorm:"type:decimal(13,3)" json:"fdmxs"` // 
 Jdsj time.Time `gorm:"type:datetime" json:"jdsj"` // 
}
