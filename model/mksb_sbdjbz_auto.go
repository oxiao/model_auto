package model

import (
	"time"
)

type MksbSbdjbz struct {
Model
 SbdjbzID string `gorm:"type:varchar(36);unique;unique_index;not null" json:"sbdjbz_id"` // 
 SbID string `gorm:"type:varchar(36)" json:"sb_id"` // 
 Sbmc string `gorm:"type:varchar(50)" json:"sbmc"` // 
 Djbw string `gorm:"type:varchar(50)" json:"djbw"` // 
 DjID string `gorm:"type:varchar(50)" json:"dj_id"` // 
 Djrq time.Time `gorm:"type:datetime" json:"djrq"` // 
 Djfs string `gorm:"type:varchar(50)" json:"djfs"` // 
 Djzt string `gorm:"type:varchar(50)" json:"djzt"` // 
 Djzq string `gorm:"type:varchar(50)" json:"djzq"` // 
 Djry string `gorm:"type:varchar(50)" json:"djry"` // 
}
