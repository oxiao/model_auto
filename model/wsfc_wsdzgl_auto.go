package model

type WsfcWsdzgl struct {
	WsdzglID string `sql:"index" gorm:"type:varchar(36);primary_key" json:"wsdzgl_id"` // 瓦斯地质规律编号_wsdzgl_id
	WsfcxxID string `gorm:"type:varchar(36)" json:"wsfcxx_id"`                         // 瓦斯赋存信息编号_wsfcxx_id
	Cjhj     string `gorm:"type:varchar(50)" json:"cjhj"`                              // 沉积环境_CJHJ
	Gzyx     string `gorm:"type:varchar(50)" json:"gzyx"`                              // 构造影响_GZYX
	Gcyx     string `gorm:"type:varchar(50)" json:"gcyx"`                              // 盖层影响_GCYX
	Msyx     string `gorm:"type:varchar(50)" json:"msyx"`                              // 埋深影响_MSYX
	Dzdyhf   string `gorm:"type:varchar(50)" json:"dzdyhf"`                            // 地质单元划分_DZDYHF
}
