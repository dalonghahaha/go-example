package models

import (
	"time"
)

//Model基类，定义通用属性
type BaseModel struct {
	ID             int64 `gorm:"column:id;primary_key" json:"id"`
	CreateDatetime int64 `gorm:"column:create_datetime" json:"create_datetime"`
	Creator        int64 `gorm:"column:creator" json:"creator"`
	UpdateDatetime int64 `gorm:"column:update_datetime" json:"update_datetime"`
	Updator        int64 `gorm:"column:updator" json:"updator"`
}

//BaseModel的BeforeCreate钩子
func (c *BaseModel) BeforeCreate() error {
	c.CreateDatetime = time.Now().Unix()
	return nil
}

//BaseModel的BeforeUpdate钩子
func (c *BaseModel) BeforeUpdate() error {
	c.UpdateDatetime = time.Now().Unix()
	return nil
}
