package model

import (
	"time"
)

type JxgcCmgzmCkq struct {
Model
 CmgzmID string `gorm:"type:varchar(36)" json:"cmgzm_id"` // 
 CkqID string `gorm:"type:varchar(36);unique;unique_index;not null" json:"ckq_id"` // 
 McID string `gorm:"type:varchar(36)" json:"mc_id"` // 
 Mc string `gorm:"type:varchar(50)" json:"mc"` // 
 Sp string `gorm:"type:varchar(50)" json:"sp"` // 
 Wz string `gorm:"type:varchar(50)" json:"wz"` // 
 DdbYx string `gorm:"type:varchar(50)" json:"ddb_yx"` // 
 Bg decimal.Decimal `gorm:"type:decimal(13,3)" json:"bg"` // 
 ZxCd decimal.Decimal `gorm:"type:decimal(13,3)" json:"zx_cd"` // 
 QxCd decimal.Decimal `gorm:"type:decimal(13,3)" json:"qx_cd"` // 
 JsMb string `gorm:"type:varchar(3)" json:"js_mb"` // 
 JsLy string `gorm:"type:varchar(50)" json:"js_ly"` // 
 JsYjJsl decimal.Decimal `gorm:"type:decimal(13,3)" json:"js_yj_jsl"` // 
 JsZcJsl decimal.Decimal `gorm:"type:decimal(13,3)" json:"js_zc_jsl"` // 
 JsZdJsl decimal.Decimal `gorm:"type:decimal(13,3)" json:"js_zd_jsl"` // 
 JsMj decimal.Decimal `gorm:"type:decimal(13,3)" json:"js_mj"` // 
 JsTcqk string `gorm:"type:varchar(50)" json:"js_tcqk"` // 
 JsSmbg decimal.Decimal `gorm:"type:decimal(13,3)" json:"js_smbg"` // 
 Bhmzls string `gorm:"type:varchar(50)" json:"bhmzls"` // 
 Xdyx string `gorm:"type:varchar(50)" json:"xdyx"` // 
 Tcqk string `gorm:"type:varchar(50)" json:"tcqk"` // 
 Gcwx string `gorm:"type:varchar(3)" json:"gcwx"` // 
 ScKsRq time.Time `gorm:"type:datetime" json:"sc_ks_rq"` // 
 ScJsRq time.Time `gorm:"type:datetime" json:"sc_js_rq"` // 
 Zhfzcs string `gorm:"type:varchar(50)" json:"zhfzcs"` // 
}
