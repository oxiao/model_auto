package model

import (
	"time"
)

type McMz struct {
Model
 MzID string `gorm:"type:varchar(36);unique;unique_index;not null" json:"mz_id"` // 
 MyID string `gorm:"type:varchar(36)" json:"my_id"` // 
 Amdhxzc string `gorm:"type:varchar(50)" json:"amdhxzc"` // 
 Bmdfzjg string `gorm:"type:varchar(50)" json:"bmdfzjg"` // 
 Mzfx string `gorm:"type:varchar(50)" json:"mzfx"` // 
 Gyfx string `gorm:"type:varchar(50)" json:"gyfx"` // 
 Ysfx string `gorm:"type:varchar(50)" json:"ysfx"` // 
 Mhcffx string `gorm:"type:varchar(50)" json:"mhcffx"` // 
 Mzfxjgdbsff string `gorm:"type:varchar(50)" json:"mzfxjgdbsff"` // 
 Mdgyxsy string `gorm:"type:varchar(50)" json:"mdgyxsy"` // 
 Yhwhl decimal.Decimal `gorm:"type:decimal(13,3)" json:"yhwhl"` // 
 Mdqh string `gorm:"type:varchar(50)" json:"mdqh"` // 
 Mdyh string `gorm:"type:varchar(50)" json:"mdyh"` // 
 Mzzyzbdfjbz string `gorm:"type:varchar(50)" json:"mzzyzbdfjbz"` // 
 Mlx string `gorm:"type:varchar(50)" json:"mlx"` // 
 Sybgb string `gorm:"type:varchar(50)" json:"sybgb"` // 
 Syzrb string `gorm:"type:varchar(50)" json:"syzrb"` // 
 Mdzhly string `gorm:"type:varchar(50)" json:"mdzhly"` // 
 Mzpj string `gorm:"type:varchar(50)" json:"mzpj"` // 
 Msjdzbjyy string `gorm:"type:varchar(50)" json:"msjdzbjyy"` // 
 Cmzy string `gorm:"type:varchar(50)" json:"cmzy"` // 
 Mdphlx string `gorm:"type:varchar(50)" json:"mdphlx"` // 
 Mdjgxxs decimal.Decimal `gorm:"type:decimal(13,3)" json:"mdjgxxs"` // 
}
