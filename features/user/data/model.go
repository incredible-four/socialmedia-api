package data

import (
	com "incrediblefour/features/comment/data"
	con "incrediblefour/features/content/data"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Avatar   string
	Banner   string
	Name     string
	Username string
	Email    string
	Password string
	Contents []con.Contents `gorm:"foreignkey:UserID"`
	Comments []com.Comments `gorm:"foreignkey:UserID"`
}
