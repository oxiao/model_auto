package model

type BmTrain struct {
	TrainID     string `sql:"index" gorm:"type:varchar(36);primary_key" json:"train_id"` // f_id
	EmployeeID  string `gorm:"type:varchar(36)" json:"employee_id"`                      // 员工编号
	Name        string `gorm:"type:varchar(20)" json:"name"`                             // 姓名
	WorkerSn    string `gorm:"type:varchar(36)" json:"worker_sn"`                        // 工号
	Post        string `gorm:"type:varchar(36)" json:"post"`                             // 岗位
	Contents    string `gorm:"type:varchar(200)" json:"contents"`                        // 培训内容
	Duration    string `gorm:"type:varchar(20)" json:"duration"`                         // 培训时长
	Cycle       string `gorm:"type:varchar(20)" json:"cycle"`                            // 培训周期
	Unit        string `gorm:"type:varchar(20)" json:"unit"`                             // 培训单位
	Date        string `gorm:"type:date" json:"date"`                                    // 培训时间
	Result      string `gorm:"type:varchar(20)" json:"result"`                           // 培训结果
	CertName    string `gorm:"type:varchar(36)" json:"cert_name"`                        // 证书名称
	CertSn      string `gorm:"type:varchar(36)" json:"cert_sn"`                          // 证书编号
	InvalidDate string `gorm:"type:date" json:"invalid_date"`                            // 有效时间
	Publisher   string `gorm:"type:varchar(50)" json:"publisher"`                        // 颁发单位
}
