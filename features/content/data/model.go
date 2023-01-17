package data

import (
	"incrediblefour/features/comment/data"

	"gorm.io/gorm"
)

type Contents struct {
	gorm.Model
	Image string
	Caption string
	UserID uint
	Comments []data.Comments `gorm:"foreignkey:ContentID"`
}