package model

type JxgcXdbz struct {
	QyID    string `sql:"index" gorm:"type:varchar(36);primary_key" json:"qy_id"` //
	CmgzmID string `gorm:"type:varchar(36)" json:"cmgzm_id"`                      //
	Lb      string `gorm:"type:varchar(3)" json:"lb"`                             // 切眼,风巷,机巷
	Bh      string `gorm:"type:varchar(50)" json:"bh"`                            //
	Mc      string `gorm:"type:varchar(50)" json:"mc"`                            //
	Tjx     string `gorm:"type:varchar(50)" json:"tjx"`                           //
}
