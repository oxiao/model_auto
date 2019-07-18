package model

import (
	"time"
)

type ZyclMcqcl struct {
Model
 McqclID string `gorm:"type:varchar(36);unique;unique_index;not null" json:"mcqcl_id"` // 
 McID string `gorm:"type:varchar(36)" json:"mc_id"` // 
 Zykcmc string `gorm:"type:varchar(50)" json:"zykcmc"` // 
 Mcqhl decimal.Decimal `gorm:"type:decimal(13,3)" json:"mcqhl"` // 
 Mcqyl decimal.Decimal `gorm:"type:decimal(13,3)" json:"mcqyl"` // 
 Mctqx string `gorm:"type:varchar(50)" json:"mctqx"` // 
 Kcl074 decimal.Decimal `gorm:"type:decimal(13,3)" json:"kcl_074"` // 
 Kcl01 decimal.Decimal `gorm:"type:decimal(13,3)" json:"kcl_01"` // 
 Hjkcl decimal.Decimal `gorm:"type:decimal(13,3)" json:"hjkcl"` // 
 Cl074 decimal.Decimal `gorm:"type:decimal(13,3)" json:"cl_074"` // 
 Cl01 decimal.Decimal `gorm:"type:decimal(13,3)" json:"cl_01"` // 
 Hjcl decimal.Decimal `gorm:"type:decimal(13,3)" json:"hjcl"` // 
 Mchqmj decimal.Decimal `gorm:"type:decimal(13,3)" json:"mchqmj"` // 
 Mzhf decimal.Decimal `gorm:"type:decimal(13,3)" json:"mzhf"` // 
 Mdpjhf decimal.Decimal `gorm:"type:decimal(13,3)" json:"mdpjhf"` // 
 Mdymjhql decimal.Decimal `gorm:"type:decimal(13,3)" json:"mdymjhql"` // 
 Mdkqgzjhql decimal.Decimal `gorm:"type:decimal(13,3)" json:"mdkqgzjhql"` // 
 FqylMcqhl decimal.Decimal `gorm:"type:decimal(13,3)" json:"fqyl_mcqhl"` // 
 YsccMcqhl decimal.Decimal `gorm:"type:decimal(13,3)" json:"yscc_mcqhl"` // 
 Mdkqgzjzlmd decimal.Decimal `gorm:"type:decimal(13,3)" json:"mdkqgzjzlmd"` // 
 Mdgzwhjmd decimal.Decimal `gorm:"type:decimal(13,3)" json:"mdgzwhjmd"` // 
 Mcqdzcl decimal.Decimal `gorm:"type:decimal(13,3)" json:"mcqdzcl"` // 
 Dzcl decimal.Decimal `gorm:"type:decimal(13,3)" json:"dzcl"` // 
 Qjljqtcl decimal.Decimal `gorm:"type:decimal(13,3)" json:"qjljqtcl"` // 
 Mcqkccl decimal.Decimal `gorm:"type:decimal(13,3)" json:"mcqkccl"` // 
 Mcjhd decimal.Decimal `gorm:"type:decimal(13,3)" json:"mcjhd"` // 
 Mcms decimal.Decimal `gorm:"type:decimal(13,3)" json:"mcms"` // 
 Stl decimal.Decimal `gorm:"type:decimal(13,3)" json:"stl"` // 
 Mzymjsf decimal.Decimal `gorm:"type:decimal(13,3)" json:"mzymjsf"` // 
 Mdphsf decimal.Decimal `gorm:"type:decimal(13,3)" json:"mdphsf"` // 
 Ysccyl decimal.Decimal `gorm:"type:decimal(13,3)" json:"ysccyl"` // 
 Fqccyl decimal.Decimal `gorm:"type:decimal(13,3)" json:"fqccyl"` // 
 Bhyl decimal.Decimal `gorm:"type:decimal(13,3)" json:"bhyl"` // 
 Lsyl decimal.Decimal `gorm:"type:decimal(13,3)" json:"lsyl"` // 
 Ljjxyl decimal.Decimal `gorm:"type:decimal(13,3)" json:"ljjxyl"` // 
 Plyl decimal.Decimal `gorm:"type:decimal(13,3)" json:"plyl"` // 
 Yltd decimal.Decimal `gorm:"type:decimal(13,3)" json:"yltd"` // 
 Jdyl decimal.Decimal `gorm:"type:decimal(13,3)" json:"jdyl"` // 
 Djrcql decimal.Decimal `gorm:"type:decimal(13,3)" json:"djrcql"` // 
 Djrcsl decimal.Decimal `gorm:"type:decimal(13,3)" json:"djrcsl"` // 
 Csl decimal.Decimal `gorm:"type:decimal(13,3)" json:"csl"` // 
 Pjjkmj decimal.Decimal `gorm:"type:decimal(13,3)" json:"pjjkmj"` // 
 Xfsj decimal.Decimal `gorm:"type:decimal(13,3)" json:"xfsj"` // 
 Ccwd decimal.Decimal `gorm:"type:decimal(13,3)" json:"ccwd"` // 
 Lstj decimal.Decimal `gorm:"type:decimal(13,3)" json:"lstj"` // 
 Dcqj decimal.Decimal `gorm:"type:decimal(13,3)" json:"dcqj"` // 
 Qxxl decimal.Decimal `gorm:"type:decimal(13,3)" json:"qxxl"` // 
}
