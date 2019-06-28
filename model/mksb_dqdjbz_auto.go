package model

import (
	"time"
)

type MksbDqdjbz struct {
	DqdjbzID string    `sql:"index" gorm:"type:varchar(36);primary_key" json:"dqdjbz_id"` //
	SbID     string    `gorm:"type:varchar(36)" json:"sb_id"`                             //
	Sbmc     string    `gorm:"type:varchar(50)" json:"sbmc"`                              //
	Sbdh     string    `gorm:"type:varchar(50)" json:"sbdh"`                              //
	Sbdmc    string    `gorm:"type:varchar(50)" json:"sbdmc"`                             //
	Djbz     string    `gorm:"type:varchar(50)" json:"djbz"`                              //
	Djzq     string    `gorm:"type:varchar(50)" json:"djzq"`                              //
	Djrq     time.Time `gorm:"type:datetime" json:"djrq"`                                 //
	Djfs     string    `gorm:"type:varchar(50)" json:"djfs"`                              //
	Djry     string    `gorm:"type:varchar(50)" json:"djry"`                              //
}
