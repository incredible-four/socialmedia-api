package data

import (
	"incrediblefour/features/comment/data"
	"incrediblefour/features/content"
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

func ToCore(data Contents) content.Core {
	return content.Core{
		ID:      data.ID,
		Image:   data.Image,
		Caption: data.Caption,
	}
}

func CoreToData(data content.Core) Contents {
	return Contents{
		Model:   gorm.Model{ID: data.ID},
		Image:   data.Image,
		Caption: data.Caption,
	}
}
