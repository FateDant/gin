package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type GormModel struct {
	gorm.Model
	Name     string
	Birthday time.Time
}
