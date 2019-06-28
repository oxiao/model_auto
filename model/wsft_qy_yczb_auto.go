package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type WsftQyYczb struct {
	ZbID   string          `sql:"index" gorm:"type:varchar(36);primary_key" json:"zb_id"` // f_id
	FtqyID string          `gorm:"type:varchar(36)" json:"ftqy_id"`                       // f_id
	Qymc   string          `gorm:"type:varchar(200)" json:"qymc"`                         // 区域名称
	Cddd   string          `gorm:"type:varchar(200)" json:"cddd"`                         // 测定地点
	Ddbg   decimal.Decimal `gorm:"type:decimal(18,3)" json:"ddbg"`                        // 地点标高
	Ddms   decimal.Decimal `gorm:"type:decimal(18,3)" json:"ddms"`                        // 地点埋深
	Csff   string          `gorm:"type:varchar(200)" json:"csff"`                         // 测试方法
	Zkkh   string          `gorm:"type:varchar(200)" json:"zkkh"`                         // 钻孔孔号
	Qysd   decimal.Decimal `gorm:"type:decimal(18,3)" json:"qysd"`                        // 取样深度
	ZbWsyl decimal.Decimal `gorm:"type:decimal(18,3)" json:"zb_wsyl"`                     // 指标-瓦斯压力
	ZbWshl decimal.Decimal `gorm:"type:decimal(18,3)" json:"zb_wshl"`                     // 指标-瓦斯含量
	ZbYcl  decimal.Decimal `gorm:"type:decimal(18,3)" json:"zb_ycl"`                      // 指标-涌出量
	ZbDbwy decimal.Decimal `gorm:"type:decimal(18,3)" json:"zb_dbwy"`                     // 指标-顶部位移
	ZbQt   string          `gorm:"type:varchar(200)" json:"zb_qt"`                        // 指标-其他
	Cdsj   time.Time       `gorm:"type:datetime" json:"cdsj"`                             // 测定时间
	Qyry   string          `gorm:"type:varchar(200)" json:"qyry"`                         // 取样人员
	Jlsj   time.Time       `gorm:"type:datetime" json:"jlsj"`                             // 记录时间
}
