package model

type JxgcGzm struct {
	GzmID string `sql:"index" gorm:"type:varchar(36);primary_key" json:"gzm_id"` //
	McID  string `gorm:"type:varchar(36)" json:"mc_id"`                          //
	SpID  string `gorm:"type:varchar(36)" json:"sp_id"`                          //
	Mc    string `gorm:"type:varchar(50)" json:"mc"`                             //
	Lx    string `gorm:"type:varchar(36)" json:"lx"`                             // 采煤工作面，掘进工作面，巷道工作面
}
