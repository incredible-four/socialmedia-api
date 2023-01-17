package data

import (
	"incrediblefour/features/comment/data"
	"time"

	"gorm.io/gorm"
)

type Contents struct {
	gorm.Model
	Image   string
	Caption string
	UserID  uint
	Comment []data.Comments `gorm:"foreignkey:ContentID"`
}

type AllContents struct {
	ID        uint
	Image     string
	Caption   string
	Username  string
	UpdatedAt time.Time
	Comment   []data.Comments `gorm:"foreignkey:ContentID"`
}
