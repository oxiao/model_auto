package model

import (
	"time"
)

type DzclCldxd struct {
Model
 DxdID string `gorm:"type:varchar(36);unique;unique_index;not null" json:"dxd_id"` // 
 Dxdlx string `gorm:"type:varchar(3)" json:"dxdlx"` // 
 Dxdzt string `gorm:"type:varchar(3)" json:"dxdzt"` // 
 Dxxxkz string `gorm:"type:varchar(3)" json:"dxxxkz"` // 
 Dm string `gorm:"type:varchar(50)" json:"dm"` // 
 X int `gorm:"type:int(11)" json:"x"` // 
 Y int `gorm:"type:int(11)" json:"y"` // 
 Z int `gorm:"type:int(11)" json:"z"` // 
 Zbjl decimal.Decimal `gorm:"type:decimal(18,3)" json:"zbjl"` // 
 Ybjl decimal.Decimal `gorm:"type:decimal(18,3)" json:"ybjl"` // 
 Dxjb string `gorm:"type:varchar(3)" json:"dxjb"` // 
 Gcrq time.Time `gorm:"type:datetime" json:"gcrq"` // 
 Gcz string `gorm:"type:varchar(50)" json:"gcz"` // 
 Gcjd decimal.Decimal `gorm:"type:decimal(13,3)" json:"gcjd"` // 
 Jsrq time.Time `gorm:"type:datetime" json:"jsrq"` // 
 Jcrq time.Time `gorm:"type:datetime" json:"jcrq"` // 
 Jcz string `gorm:"type:varchar(50)" json:"jcz"` // 
 Kcsp string `gorm:"type:varchar(50)" json:"kcsp"` // 
 Szcq string `gorm:"type:varchar(50)" json:"szcq"` // 
 Szxd string `gorm:"type:varchar(50)" json:"szxd"` // 
 Gzdd string `gorm:"type:varchar(50)" json:"gzdd"` // 
 Mdfs string `gorm:"type:varchar(50)" json:"mdfs"` // 
}
