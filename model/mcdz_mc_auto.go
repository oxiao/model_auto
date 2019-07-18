package model

import (
	"time"
)

type McdzMc struct {
Model
 McID string `gorm:"type:varchar(36);unique;unique_index;not null" json:"mc_id"` // 
 Mc string `gorm:"type:varchar(50)" json:"mc"` // 
 Pjhd decimal.Decimal `gorm:"type:decimal(13,3)" json:"pjhd"` // 
 Zdhd decimal.Decimal `gorm:"type:decimal(13,3)" json:"zdhd"` // 
 Zxhd decimal.Decimal `gorm:"type:decimal(13,3)" json:"zxhd"` // 
 Kchd decimal.Decimal `gorm:"type:decimal(13,3)" json:"kchd"` // 
 Dblx string `gorm:"type:varchar(50)" json:"dblx"` // 
 Dbhd string `gorm:"type:varchar(3)" json:"dbhd"` // 
 Dbyx string `gorm:"type:varchar(50)" json:"dbyx"` // 
 Diblx string `gorm:"type:varchar(50)" json:"diblx"` // 
 Dibhd decimal.Decimal `gorm:"type:decimal(13,3)" json:"dibhd"` // 
 Dibyx string `gorm:"type:varchar(50)" json:"dibyx"` // 
 Jsjl decimal.Decimal `gorm:"type:decimal(13,3)" json:"jsjl"` // 
 Jxjl decimal.Decimal `gorm:"type:decimal(13,3)" json:"jxjl"` // 
 Jgyx string `gorm:"type:varchar(50)" json:"jgyx"` // 
 Jghd string `gorm:"type:varchar(50)" json:"jghd"` // 
 Jgl decimal.Decimal `gorm:"type:decimal(13,3)" json:"jgl"` // 
 Jhxt string `gorm:"type:varchar(50)" json:"jhxt"` // 
 Fcxs string `gorm:"type:varchar(50)" json:"fcxs"` // 
 Jmfs string `gorm:"type:varchar(50)" json:"jmfs"` // 
 Hdfj decimal.Decimal `gorm:"type:decimal(13,3)" json:"hdfj"` // 
 Qxfj string `gorm:"type:varchar(50)" json:"qxfj"` // 
 Mcjg string `gorm:"type:varchar(50)" json:"mcjg"` // 
 Mcqj decimal.Decimal `gorm:"type:decimal(13,3)" json:"mcqj"` // 
 Pjqj decimal.Decimal `gorm:"type:decimal(13,3)" json:"pjqj"` // 
 Btjcff string `gorm:"type:varchar(50)" json:"btjcff"` // 
 Kcx string `gorm:"type:varchar(50)" json:"kcx"` // 
 Mjkcl decimal.Decimal `gorm:"type:decimal(13,3)" json:"mjkcl"` // 
 Ltx string `gorm:"type:varchar(50)" json:"ltx"` // 
 Ltkd decimal.Decimal `gorm:"type:decimal(13,3)" json:"ltkd"` // 
 Bzwxx string `gorm:"type:varchar(50)" json:"bzwxx"` // 
 Zrqxfl string `gorm:"type:varchar(50)" json:"zrqxfl"` // 
 Fhzq decimal.Decimal `gorm:"type:decimal(13,3)" json:"fhzq"` // 
 Cmdznd string `gorm:"type:varchar(50)" json:"cmdznd"` // 
 Cwwdx string `gorm:"type:varchar(50)" json:"cwwdx"` // 
 Wyxz string `gorm:"type:varchar(50)" json:"wyxz"` // 
 Dbkkcd string `gorm:"type:varchar(50)" json:"dbkkcd"` // 
}
