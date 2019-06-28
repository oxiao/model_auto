package model

type DzclCldx struct {
	XdxxID string `gorm:"type:varchar(36)" json:"xdxx_id"` //
	Bh     string `gorm:"type:varchar(36)" json:"bh"`      //
	Mc     string `gorm:"type:varchar(50)" json:"mc"`      //
	Xd     string `gorm:"type:varchar(50)" json:"xd"`      //
}
