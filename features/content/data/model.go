package data

import "gorm.io/gorm"

type Content struct {
	gorm.Model
	Image   string
	Caption string
	UserID  uint
}
