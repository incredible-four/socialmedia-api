package data

import (
	com "incrediblefour/features/comment/data"
	con "incrediblefour/features/content/data"
	"incrediblefour/features/user"
	"mime/multipart"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Avatar   string
	Banner   string
	Name     string
	Username string
	Bio      string
	Email    string
	Password string
	Contents []con.Contents `gorm:"foreignkey:UserID"`
	Comments []com.Comments `gorm:"foreignkey:UserID"`
}

type File struct {
	File multipart.File `json:"file,omitempty" validate:"required"`
}

type Url struct {
	Url string `json:"url,omitempty" validate:"required"`
}

func ToCore(data Users) user.Core {
	return user.Core{
		ID:       data.ID,
		Avatar:   data.Avatar,
		Banner:   data.Banner,
		Name:     data.Name,
		Username: data.Username,
		Bio:      data.Bio,
		Email:    data.Email,
		Password: data.Password,
	}
}

func CoreToData(data user.Core) Users {
	return Users{
		Model:    gorm.Model{ID: data.ID},
		Avatar:   data.Avatar,
		Banner:   data.Banner,
		Name:     data.Name,
		Username: data.Username,
		Bio:      data.Bio,
		Email:    data.Email,
		Password: data.Password,
	}
}
