package model

import (
	"time"
)

type BmEmployee struct {
	EmployeeID    string    `sql:"index" gorm:"type:varchar(36);primary_key" json:"employee_id"` // 员工编号
	WorktypeID    string    `gorm:"type:varchar(36)" json:"worktype_id"`                         //
	JobtitleID    string    `gorm:"type:varchar(36)" json:"jobtitle_id"`                         //
	WorkerSn      string    `gorm:"type:varchar(36)" json:"worker_sn"`                           //
	Name          string    `gorm:"type:varchar(36)" json:"name"`                                // 姓名
	SimpleName    string    `gorm:"type:varchar(36)" json:"simple_name"`                         // 簡化名
	AliasName     string    `gorm:"type:varchar(36)" json:"alias_name"`                          //
	Sex           string    `gorm:"type:varchar(36)" json:"sex"`                                 // 性别
	Qualification string    `gorm:"type:varchar(36)" json:"qualification"`                       // 资质
	Photo         string    `gorm:"type:text" json:"photo"`                                      // 相片
	Bod           time.Time `gorm:"type:datetime" json:"bod"`                                    // 生日
	BirthdPlace   string    `gorm:"type:varchar(36)" json:"birthd_place"`                        // 籍贯
	Nationality   string    `gorm:"type:varchar(3)" json:"nationality"`                          //
	BloodType     string    `gorm:"type:varchar(3)" json:"blood_type"`                           //
	Height        string    `gorm:"type:varchar(36)" json:"height"`                              //
	Weight        string    `gorm:"type:varchar(36)" json:"weight"`                              //
	Look          string    `gorm:"type:varchar(36)" json:"look"`                                //
	IdentityType  string    `gorm:"type:varchar(36)" json:"identity_type"`                       // 证件
	IdentityNo    string    `gorm:"type:varchar(64)" json:"identity_no"`                         // 证件号码
	Education     string    `gorm:"type:varchar(36)" json:"education"`                           // 學歷
	Marry         string    `gorm:"type:varchar(36)" json:"marry"`                               // 婚姻
	Party         string    `gorm:"type:varchar(36)" json:"party"`                               // 政治面貌
	JoinTime      time.Time `gorm:"type:datetime" json:"join_time"`                              //
	WorkTime      time.Time `gorm:"type:datetime" json:"work_time"`                              //
	Email         string    `gorm:"type:varchar(128)" json:"email"`                              // 邮箱
	HandPhone     string    `gorm:"type:varchar(64)" json:"hand_phone"`                          // 手机
	Phone         string    `gorm:"type:varchar(64)" json:"phone"`                               // 电话
	Im            string    `gorm:"type:varchar(128)" json:"im"`                                 // IM
	Address       string    `gorm:"type:varchar(1024)" json:"address"`                           // 联系地址
	City          string    `gorm:"type:varchar(64)" json:"city"`                                // 城市
	Station       string    `gorm:"type:varchar(64)" json:"station"`                             // 州
	Country       string    `gorm:"type:varchar(64)" json:"country"`                             // 国家
	Post          string    `gorm:"type:varchar(36)" json:"post"`                                // 邮编
	LoginLimit    int       `gorm:"type:int(11)" json:"login_limit"`                             // 0 可登录，1 不可登录
	IsSysPerson   int       `gorm:"type:int(11)" json:"is_sys_person"`                           // 0 非系统人员，1系统人员。系统人员是指仅对平台进行管理的人员。
	CreateDate    time.Time `gorm:"type:datetime" json:"create_date"`                            // 创建时间
	Creator       string    `gorm:"type:varchar(36)" json:"creator"`                             // 创建人
	Note          string    `gorm:"type:varchar(256)" json:"note"`                               // 备注
	Info          string    `gorm:"type:varchar(4000)" json:"info"`                              // 简介
	Idx           int       `gorm:"type:int(11)" json:"idx"`                                     // 排序
	State         string    `gorm:"type:varchar(3)" json:"state"`                                // 0正常，1停用，－1刪除
	StateTime     time.Time `gorm:"type:datetime" json:"state_time"`                             // 标志更改时的时间
}
