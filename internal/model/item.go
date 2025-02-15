package model

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Name string `gorm:"index"`
	Type string
}
