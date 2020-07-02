package comjia

import (
	"julive/components/db"
	"julive/models"
)

type CjBlackWord struct {
	models.BaseModel
	Name     string `gorm:"column:name" json:"name" form:"name" binding:"required"`
	CityID   int    `gorm:"column:city_id" json:"city_id" form:"city_id" binding:"required"`
	WordType int    `gorm:"column:word_type" json:"word_type" form:"word_type" binding:"required"`
}

func (c *CjBlackWord) TableName() string {
	return "cj_black_words"
}

func (c *CjBlackWord) GetAll() ([]*CjBlackWord, error) {
	list := []*CjBlackWord{}
	errors := db.Get("comjia").Find(&list).GetErrors()
	if len(errors) > 0 {
		return nil, errors[0]
	}
	return list, nil
}

func (c *CjBlackWord) GetByID(id int64) error {
	err := db.Get("comjia").Where("id = ? ", id).First(&c).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *CjBlackWord) Create() error {
	errors := db.Get("comjia").Create(c).GetErrors()
	if len(errors) > 0 {
		return errors[0]
	}
	return nil
}

func (c *CjBlackWord) Update() error {
	errors := db.Get("comjia").Save(c).GetErrors()
	if len(errors) > 0 {
		return errors[0]
	}
	return nil
}

func (c *CjBlackWord) Delete() error {
	errors := db.Get("comjia").Delete(c).GetErrors()
	if len(errors) > 0 {
		return errors[0]
	}
	return nil
}
