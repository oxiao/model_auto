package model

type MksbSbqy struct {
	SbqyID string `sql:"index" gorm:"type:varchar(36);primary_key" json:"sbqy_id"` // 区域编号
	KjqyBm string `gorm:"type:varchar(36)" json:"kjqy_bm"`                         //
	Qybm   string `gorm:"type:varchar(50)" json:"qybm"`                            // 区域编码
	Qymc   string `gorm:"type:varchar(50)" json:"qymc"`                            // 区域名称
	Sbfl   string `gorm:"type:varchar(3)" json:"sbfl"`                             // 设备分类
	QyLx   string `gorm:"type:varchar(3)" json:"qy_lx"`                            // 区域类型
}
