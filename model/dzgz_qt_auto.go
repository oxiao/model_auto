package model

type DzgzQt struct {
	McID string `gorm:"type:varchar(36)" json:"mc_id"`                         //
	QtID string `sql:"index" gorm:"type:varchar(36);primary_key" json:"qt_id"` //
	Mc   string `gorm:"type:varchar(50)" json:"mc"`                            //
	Lx   string `gorm:"type:varchar(36)" json:"lx"`                            //
	Xz   string `gorm:"type:varchar(36)" json:"xz"`                            // 类型：岩浆岩，断裂，火成岩侵入，降落柱，节理，褶曲
	Fgfl string `gorm:"type:varchar(50)" json:"fgfl"`                          //
	Ly   string `gorm:"type:varchar(36)" json:"ly"`                            //
	Ms   string `gorm:"type:varchar(200)" json:"ms"`                           //
}
