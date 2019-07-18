package model

import (
	"time"
)

type WsfcMcwsc struct {
Model
 CsID string `gorm:"type:varchar(36);unique;unique_index;not null" json:"cs_id"` // 
 McID string `gorm:"type:varchar(36)" json:"mc_id"` // 
 WsfcxxID string `gorm:"type:varchar(36)" json:"wsfcxx_id"` // 瓦斯赋存信息编号_wsfcxx_id
 Mygyfxzb decimal.Decimal `gorm:"type:decimal(13,3)" json:"mygyfxzb"` // 
 Wsxfcs decimal.Decimal `gorm:"type:decimal(13,3)" json:"wsxfcs"` // 
 Mdkxl decimal.Decimal `gorm:"type:decimal(13,3)" json:"mdkxl"` // 
 Zkwsycl decimal.Decimal `gorm:"type:decimal(13,3)" json:"zkwsycl"` // 
 Mctqxxs decimal.Decimal `gorm:"type:decimal(13,3)" json:"mctqxxs"` // 
 Mdjgxxs decimal.Decimal `gorm:"type:decimal(13,3)" json:"mdjgxxs"` // 
 Zkwsllsjxs decimal.Decimal `gorm:"type:decimal(13,3)" json:"zkwsllsjxs"` // 
 Bmmkzrwsll decimal.Decimal `gorm:"type:decimal(13,3)" json:"bmmkzrwsll"` // 
 Wsfscsd decimal.Decimal `gorm:"type:decimal(13,3)" json:"wsfscsd"` // 
 Csdw string `gorm:"type:varchar(20)" json:"csdw"` // 测试单位
 Cssj time.Time `gorm:"type:datetime" json:"cssj"` // 测试时间
}
