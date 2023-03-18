package models

import "github.com/jinzhu/gorm"

type TagModel struct {
	gorm.Model
	Title   string `gorm:"not null;size:64"`
	Content string `gorm:"not null;index:title_index"`
	Test    string `gorm:"-"` //忽略
	Desc    string `gorm:"type:text"`
}
