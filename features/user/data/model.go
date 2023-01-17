package data

import (
	com "incrediblefour/features/comment/data"
	con "incrediblefour/features/content/data"
	"incrediblefour/features/user"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Avatar string
	Banner string
	Name string
	Username string
	Email string
	Password string
	Contents []con.Contents `gorm:"foreignkey:UserID"`
	Comments []com.Comments `gorm:"foreignkey:UserID"`
}

func ToCore(data Users) user.Core {
	return user.Core{
		ID: data.ID,
		Avatar: data.Avatar,
		Banner: data.Banner,
		Name: data.Name,
		Username: data.Username,
		Email: data.Email,
		Password: data.Password,
	}
}

func CoreToData(data user.Core) Users {
	return Users{
		Model: gorm.Model{ID: data.ID},
		Avatar: data.Avatar,
		Banner: data.Banner,
		Name: data.Name,
		Username: data.Username,
		Email: data.Email,
		Password: data.Password,
	}
}