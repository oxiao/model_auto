package model

import (
	"time"
)

type DzktDmzk struct {
Model
 DmzkID string `gorm:"type:varchar(36);unique;unique_index;not null" json:"dmzk_id"` // 
 MzID string `gorm:"type:varchar(36)" json:"mz_id"` // 
 ZkBh string `gorm:"type:varchar(50)" json:"zk_bh"` // 
 Zkmc string `gorm:"type:varchar(50)" json:"zkmc"` // 
 Zklx string `gorm:"type:varchar(3)" json:"zklx"` // 
 Kkwz string `gorm:"type:varchar(50)" json:"kkwz"` // 
 Kkrq time.Time `gorm:"type:datetime" json:"kkrq"` // 
 Jgrq time.Time `gorm:"type:datetime" json:"jgrq"` // 
 Zksd decimal.Decimal `gorm:"type:decimal(13,3)" json:"zksd"` // 
 Zkcw string `gorm:"type:varchar(50)" json:"zkcw"` // 
 Sgdw string `gorm:"type:varchar(50)" json:"sgdw"` // 
 Zkjg string `gorm:"type:varchar(50)" json:"zkjg"` // 
 ZkcjYsmc string `gorm:"type:varchar(50)" json:"zkcj_ysmc"` // 
 ZkcjCjhd decimal.Decimal `gorm:"type:decimal(13,3)" json:"zkcj_cjhd"` // 
 ZkcjDjsd decimal.Decimal `gorm:"type:decimal(13,3)" json:"zkcj_djsd"` // 
 Fksj time.Time `gorm:"type:datetime" json:"fksj"` // 
 Fkzl decimal.Decimal `gorm:"type:decimal(13,3)" json:"fkzl"` // 
 Fkdw string `gorm:"type:varchar(50)" json:"fkdw"` // 
 ZkcxKs string `gorm:"type:varchar(50)" json:"zkcx_ks"` // 
 ZkcxFwj string `gorm:"type:varchar(50)" json:"zkcx_fwj"` // 
 ZkcxTdj decimal.Decimal `gorm:"type:decimal(13,3)" json:"zkcx_tdj"` // 
 ZkzlDznd string `gorm:"type:varchar(50)" json:"zkzl_dznd"` // 
 ZkzlDchd decimal.Decimal `gorm:"type:decimal(13,3)" json:"zkzl_dchd"` // 
 ZkzlDcms string `gorm:"type:varchar(50)" json:"zkzl_dcms"` // 
 ZkzlMs string `gorm:"type:varchar(50)" json:"zkzl_ms"` // 
 SzhyHsc string `gorm:"type:varchar(50)" json:"szhy_hsc"` // 
 SzhyCysj time.Time `gorm:"type:datetime" json:"szhy_cysj"` // 
 SzhyHydw string `gorm:"type:varchar(50)" json:"szhy_hydw"` // 
 SzhySzcs string `gorm:"type:varchar(50)" json:"szhy_szcs"` // 
 CssyHsc string `gorm:"type:varchar(50)" json:"cssy_hsc"` // 
 CssyCssj time.Time `gorm:"type:datetime" json:"cssy_cssj"` // 
 CssyLlgc string `gorm:"type:varchar(50)" json:"cssy_llgc"` // 
 CssySwbh string `gorm:"type:varchar(50)" json:"cssy_swbh"` // 
 CssySzcy string `gorm:"type:varchar(50)" json:"cssy_szcy"` // 
 LsHsc string `gorm:"type:varchar(50)" json:"ls_hsc"` // 
 LsLsd string `gorm:"type:varchar(50)" json:"ls_lsd"` // 
 LsLscd string `gorm:"type:varchar(50)" json:"ls_lscd"` // 
 LsYxms string `gorm:"type:varchar(50)" json:"ls_yxms"` // 
 LsYxcd string `gorm:"type:varchar(50)" json:"ls_yxcd"` // 
 CgMcmc string `gorm:"type:varchar(50)" json:"cg_mcmc"` // 
 CgMcwz string `gorm:"type:varchar(50)" json:"cg_mcwz"` // 
 CgMchd decimal.Decimal `gorm:"type:decimal(13,3)" json:"cg_mchd"` // 
 CgMcddb string `gorm:"type:varchar(50)" json:"cg_mcddb"` // 
}
