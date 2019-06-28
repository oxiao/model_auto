package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type BmFetchLog struct {
	LogID             string          `sql:"index" gorm:"type:varchar(36);primary_key" json:"log_id"` //
	TaskID            string          `gorm:"type:varchar(36)" json:"task_id"`                        //
	Jobname           string          `gorm:"type:varchar(255)" json:"jobname"`                       //
	Executiontime     time.Time       `gorm:"type:datetime" json:"executiontime"`                     //
	Executionduration decimal.Decimal `gorm:"type:decimal(18,2)" json:"executionduration"`            //
	Createddatetime   time.Time       `gorm:"type:datetime" json:"createddatetime"`                   //
	Runlog            string          `gorm:"type:text" json:"runlog"`                                //
}
