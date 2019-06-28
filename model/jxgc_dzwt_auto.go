package model

type JxgcDzwt struct {
	DzwtID string `sql:"index" gorm:"type:varchar(36);primary_key" json:"dzwt_id"` //
	Wtff   string `gorm:"type:varchar(50)" json:"wtff"`                            //
	Tcfa   string `gorm:"type:varchar(50)" json:"tcfa"`                            //
	Tccg   string `gorm:"type:varchar(50)" json:"tccg"`                            //
}
