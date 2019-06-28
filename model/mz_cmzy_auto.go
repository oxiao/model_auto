package model

type MzCmzy struct {
	CmzyID   string `sql:"index" gorm:"type:varchar(36);primary_key" json:"cmzy_id"` //
	MzID     string `gorm:"type:varchar(36)" json:"mz_id"`                           //
	Jmqi     string `gorm:"type:varchar(50)" json:"jmqi"`                            //
	Jmd      string `gorm:"type:varchar(50)" json:"jmd"`                             //
	Jmzx     string `gorm:"type:varchar(50)" json:"jmzx"`                            //
	Jmq      string `gorm:"type:varchar(50)" json:"jmq"`                             //
	Hmq      string `gorm:"type:varchar(50)" json:"hmq"`                             //
	Fmzx     string `gorm:"type:varchar(50)" json:"fmzx"`                            //
	Fmd      string `gorm:"type:varchar(50)" json:"fmd"`                             //
	Fmdzbxs  string `gorm:"type:varchar(50)" json:"fmdzbxs"`                         //
	Cmwz     string `gorm:"type:varchar(50)" json:"cmwz"`                            //
	Cmzylx   string `gorm:"type:varchar(50)" json:"cmzylx"`                          //
	Cmwzdjfs string `gorm:"type:varchar(50)" json:"cmwzdjfs"`                        //
	Cmjd     string `gorm:"type:varchar(50)" json:"cmjd"`                            //
	Mhxl     string `gorm:"type:varchar(50)" json:"mhxl"`                            //
	Mhgj     string `gorm:"type:varchar(50)" json:"mhgj"`                            //
	Mhjd     string `gorm:"type:varchar(50)" json:"mhjd"`                            //
	Mhyb     string `gorm:"type:varchar(50)" json:"mhyb"`                            //
	Mhtd     string `gorm:"type:varchar(50)" json:"mhtd"`                            //
	Rgmhsy   string `gorm:"type:varchar(50)" json:"rgmhsy"`                          //
	Mhxs     string `gorm:"type:varchar(50)" json:"mhxs"`                            //
	Cmcs     string `gorm:"type:varchar(50)" json:"cmcs"`                            //
}
