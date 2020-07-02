package comjia

import (
	"julive/components/db"
	"julive/models"
)

type CjVideoInfo struct {
	models.BaseModel
	VideoLong        string `gorm:"column:video_long" json:"video_long"`
	VideoHeaderImg   string `gorm:"column:video_header_img" json:"video_header_img"`
	VideoImg         string `gorm:"column:video_img" json:"video_img"`
	VideoURL         string `gorm:"column:video_url" json:"video_url"`
	VideoSuffix      string `gorm:"column:video_suffix" json:"video_suffix"`
	ObjectID         int64  `gorm:"column:object_id" json:"object_id"`
	ClickPlayNum     int64  `gorm:"column:click_play_num" json:"click_play_num"`
	PlayCompletedNum int    `gorm:"column:play_completed_num" json:"play_completed_num"`
	VideoWidth       string `gorm:"column:video_width" json:"video_width"`
	VideoHeight      string `gorm:"column:video_height" json:"video_height"`
}

func (c *CjVideoInfo) TableName() string {
	return "cj_video_info"
}

func (c *CjVideoInfo) GetByID(id int64) error {
	err := db.Get("comjia").Where("id = ? ", id).First(&c).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *CjVideoInfo) Create() error {
	errors := db.Get("comjia").Create(c).GetErrors()
	if len(errors) > 0 {
		return errors[0]
	}
	return nil
}

func (c *CjVideoInfo) Update() error {
	errors := db.Get("comjia").Save(c).GetErrors()
	if len(errors) > 0 {
		return errors[0]
	}
	return nil
}

func (c *CjVideoInfo) Delete() error {
	errors := db.Get("comjia").Delete(c).GetErrors()
	if len(errors) > 0 {
		return errors[0]
	}
	return nil
}
