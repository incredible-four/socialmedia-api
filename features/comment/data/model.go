package data

import "gorm.io/gorm"

type Comments struct {
	gorm.Model
	Text string
	ContentID uint
	UserID uint
}