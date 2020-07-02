package hr

import "julive/components/db"

type YwEmployee struct {
	ID                    int64  `gorm:"column:id;primary_key" json:"id"`
	EmployeeName          string `gorm:"column:employee_name" json:"employee_name"`
	Mobile                string `gorm:"column:mobile" json:"mobile"`
	PersonalMobile        string `gorm:"column:personal_mobile" json:"personal_mobile"`
	PersonalWechat        string `gorm:"column:personal_wechat" json:"personal_wechat"`
	OldMobiles            string `gorm:"column:old_mobiles" json:"old_mobiles"`
	EliteFlag             int64  `gorm:"column:elite_flag" json:"elite_flag"`
	JobNumber             int64  `gorm:"column:job_number" json:"job_number"`
	CreateDatetime        int64  `gorm:"column:create_datetime" json:"create_datetime"`
	UpdateDatetime        int64  `gorm:"column:update_datetime" json:"update_datetime"`
	Creator               int64  `gorm:"column:creator" json:"creator"`
	Updator               int64  `gorm:"column:updator" json:"updator"`
	Email                 string `gorm:"column:email" json:"email"`
	Password              string `gorm:"column:password" json:"password"`
	UpdatePasswordTime    int    `gorm:"column:update_password_time" json:"update_password_time"`
	PasswordSafety        int64  `gorm:"column:password_safety" json:"password_safety"`
	PasswordResetToken    string `gorm:"column:password_reset_token" json:"password_reset_token"`
	AuthKey               string `gorm:"column:auth_key" json:"auth_key"`
	Sex                   int64  `gorm:"column:sex" json:"sex"`
	Status                int64  `gorm:"column:status" json:"status"`
	OffjobDatetime        int    `gorm:"column:offjob_datetime" json:"offjob_datetime"`
	Academy               string `gorm:"column:academy" json:"academy"`
	Education             string `gorm:"column:education" json:"education"`
	Birth                 int    `gorm:"column:birth" json:"birth"`
	MoreMobiles           string `gorm:"column:more_mobiles" json:"more_mobiles"`
	DepartmentName        string `gorm:"column:department_name" json:"department_name"`
	DingdingID            string `gorm:"column:dingding_id" json:"dingding_id"`
	DepartmentID          int64  `gorm:"column:department_id" json:"department_id"`
	CityID                int    `gorm:"column:city_id" json:"city_id"`
	AccToken              string `gorm:"column:acc_token" json:"acc_token"`
	PostID                int    `gorm:"column:post_id" json:"post_id"`
	DidiMobile            string `gorm:"column:didi_mobile" json:"didi_mobile"`
	LastOperationDatetime int64  `gorm:"column:last_operation_datetime" json:"last_operation_datetime"`
	BdpUsername           string `gorm:"column:bdp_username" json:"bdp_username"`
	ConfluenceUsername    string `gorm:"column:confluence_username" json:"confluence_username"`
	ZentaoUsername        string `gorm:"column:zentao_username" json:"zentao_username"`
	AskUsername           string `gorm:"column:ask_username" json:"ask_username"`
	YxtUsername           string `gorm:"column:yxt_username" json:"yxt_username"`
	SopDays               int    `gorm:"column:sop_days" json:"sop_days"`
	HeaderType            int64  `gorm:"column:header_type" json:"header_type"`
	LastCityID            int    `gorm:"column:last_city_id" json:"last_city_id"`
	Cwiki                 string `gorm:"column:cwiki" json:"cwiki"`
	TotalAllocNum         int    `gorm:"column:total_alloc_num" json:"total_alloc_num"`
	AllocPolicy           int    `gorm:"column:alloc_policy" json:"alloc_policy"`
	IsProhibitYxt         int64  `gorm:"column:is_prohibit_yxt" json:"is_prohibit_yxt"`
	AdjustCityID          int64  `gorm:"column:adjust_city_id" json:"adjust_city_id"`
}

func (y *YwEmployee) TableName() string {
	return "yw_employee"
}

func (y *YwEmployee) GeByID(id int64) error {
	err := db.Get("hr").Where("id = ? ", id).First(&y).Error
	if err != nil {
		return err
	}
	return nil
}
