package model

import (
	"time"
)

type WssgTcsg struct {
Model
 TcsgID string `gorm:"type:varchar(36);unique;unique_index;not null" json:"tcsg_id"` // 
 Sgmc string `gorm:"type:varchar(255)" json:"sgmc"` // 事故名称
 Sglx string `gorm:"type:varchar(50)" json:"sglx"` // 事故类型编码
 Sgjb string `gorm:"type:varchar(50)" json:"sgjb"` // 事故级别
 Sgsj time.Time `gorm:"type:datetime" json:"sgsj"` // 事故时间
 Sgdd string `gorm:"type:varchar(255)" json:"sgdd"` // 事故地点
 Jdbcs decimal.Decimal `gorm:"type:decimal(18,3)" json:"jdbcs"` // 距地表垂深
 Tfsyt string `gorm:"type:varchar(255)" json:"tfsyt"` // 事故地点通风系统示意图（CAD、图片或PDF等）
 Zyqk string `gorm:"type:varchar(50)" json:"zyqk"` // 作业情况
 Ssrs int `gorm:"type:int(11)" json:"ssrs"` // 受伤人数
 Swrs string `gorm:"type:varchar(100)" json:"swrs"` // 死亡人数
 Bjfw string `gorm:"type:varchar(50)" json:"bjfw"` // 波及范围
 Sscd string `gorm:"type:varchar(36)" json:"sscd"` // 损失程度
 Sgyy string `gorm:"type:varchar(100)" json:"sgyy"` // 事故原因
 Sgcl string `gorm:"type:varchar(100)" json:"sgcl"` // 事故处理
 Jycs string `gorm:"type:varchar(100)" json:"jycs"` // 救援措施
 Sgxctp string `gorm:"type:varchar(255)" json:"sgxctp"` // 事故现场图片（多张图片类型）
 Zhfs string `gorm:"type:varchar(100)" json:"zhfs"` // 支护方式
 Ddbg decimal.Decimal `gorm:"type:decimal(18,3)" json:"ddbg"` // 地点标高（m）
 Yxfl decimal.Decimal `gorm:"type:decimal(18,3)" json:"yxfl"` // 有效风量
 Zcwsnd decimal.Decimal `gorm:"type:decimal(18,3)" json:"zcwsnd"` // 正常瓦斯浓度
 Jdwsycl decimal.Decimal `gorm:"type:decimal(18,3)" json:"jdwsycl"` // 绝对瓦斯涌出量
 Zsrs int `gorm:"type:int(11)" json:"zsrs"` // 重伤人数
 Qsrs int `gorm:"type:int(11)" json:"qsrs"` // 轻伤人数
 Jjss decimal.Decimal `gorm:"type:decimal(18,3)" json:"jjss"` // 直接经济损失（万元）
 Jjjjss decimal.Decimal `gorm:"type:decimal(18,3)" json:"jjjjss"` // 间接经济损失（万元）
 Sgyz string `gorm:"type:varchar(255)" json:"sgyz"` // 事故预兆
 McTz string `gorm:"type:varchar(100)" json:"mc_tz"` // 煤层特征名称
 McHd string `gorm:"type:varchar(50)" json:"mc_hd"` // 煤层厚度
 McQj decimal.Decimal `gorm:"type:decimal(18,3)" json:"mc_qj"` // 煤层倾角
 McYd decimal.Decimal `gorm:"type:decimal(18,3)" json:"mc_yd"` // 煤层硬度
 KcqkSbljc string `gorm:"type:varchar(100)" json:"kcqk_sbljc"` // 上部临近层开采情况
 KcqkXbljc string `gorm:"type:varchar(255)" json:"kcqk_xbljc"` // 下部临近层开采情况
 Ddlx string `gorm:"type:varchar(255)" json:"ddlx"` // 地点类型
 Dzgz string `gorm:"type:varchar(255)" json:"dzgz"` // 地质构造
 TcID string `gorm:"type:varchar(36);unique;unique_index;not null" json:"tc_id"` // 煤与瓦斯突出Id
 FPjjl decimal.Decimal `gorm:"type:decimal(18,3)" json:"f_pjjl"` // 棚间距离
 Dkjl decimal.Decimal `gorm:"type:decimal(18,3)" json:"dkjl"` // 控顶距离
 Tclx string `gorm:"type:varchar(50)" json:"tclx"` // 突出类型
 Mcpmt string `gorm:"type:varchar(255)" json:"mcpmt"` // 突出处煤层剖面图
 Dbwyzzt string `gorm:"type:varchar(255)" json:"dbwyzzt"` // 顶板岩层柱状图
 Dbyczzt string `gorm:"type:varchar(255)" json:"dbyczzt"` // 底板岩层柱状图
 Djqk string `gorm:"type:varchar(255)" json:"djqk"` // 突出孔洞及煤堆积情况
 PcMld string `gorm:"type:varchar(255)" json:"pc_mld"` // 喷出煤的粒度
 Qt string `gorm:"type:varchar(255)" json:"qt"` // 其他
 Jj decimal.Decimal `gorm:"type:decimal(18,3)" json:"jj"` // 孔洞形状轴线与水平面之夹角
 PcMl decimal.Decimal `gorm:"type:decimal(18,3)" json:"pc_ml"` // 喷出煤量
 PcYs decimal.Decimal `gorm:"type:decimal(18,3)" json:"pc_ys"` // 喷出岩石量
 PcJl decimal.Decimal `gorm:"type:decimal(18,3)" json:"pc_jl"` // 煤喷出距离
 PcPd decimal.Decimal `gorm:"type:decimal(18,3)" json:"pc_pd"` // 煤喷出堆积坡度
 PcMfx string `gorm:"type:varchar(255)" json:"pc_mfx"` // 喷出煤的分选情况
 Hyps string `gorm:"type:varchar(255)" json:"hyps"` // 突出地点附近围岩破碎情况
 Mcps string `gorm:"type:varchar(255)" json:"mcps"` // 突出地点附近煤层破碎情况
 Dlxy string `gorm:"type:varchar(255)" json:"dlxy"` // 动力效应
 Tcqyl string `gorm:"type:varchar(255)" json:"tcqyl"` // 突出前瓦斯压力情况
 Tchyc string `gorm:"type:varchar(255)" json:"tchyc"` // 突出后瓦斯涌出情况
}
