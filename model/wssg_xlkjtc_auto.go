package model

import (
	"time"
)

type WssgXlkjtc struct {
Model
 XlkjtcxxID string `gorm:"type:varchar(36);unique;unique_index;not null" json:"xlkjtcxx_id"` // 
 Tcsg string `gorm:"type:varchar(50)" json:"tcsg"` // 
 Dlxx string `gorm:"type:varchar(50)" json:"dlxx"` // 
 Wsyl decimal.Decimal `gorm:"type:decimal(13,3)" json:"wsyl"` // 
 Wshl decimal.Decimal `gorm:"type:decimal(13,3)" json:"wshl"` // 
 Wsycl decimal.Decimal `gorm:"type:decimal(13,3)" json:"wsycl"` // 
}
