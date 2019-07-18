package model

import (
	"time"
)

type JxgcDzwt struct {
Model
 DzwtID string `gorm:"type:varchar(36);unique;unique_index;not null" json:"dzwt_id"` // 
 Wtff string `gorm:"type:varchar(50)" json:"wtff"` // 
 Tcfa string `gorm:"type:varchar(50)" json:"tcfa"` // 
 Tccg string `gorm:"type:varchar(50)" json:"tccg"` // 
}
