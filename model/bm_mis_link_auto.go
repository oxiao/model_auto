package model

type BmMisLink struct {
	SrcMisID string `sql:"index" gorm:"type:varchar(36);primary_key" json:"src_mis_id"` //
	DstMisID string `sql:"index" gorm:"type:varchar(36);primary_key" json:"dst_mis_id"` //
	Parms    string `gorm:"type:varchar(100)" json:"parms"`                             //
}
