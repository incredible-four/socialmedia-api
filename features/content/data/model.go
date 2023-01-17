package data

import (
	"incrediblefour/features/content"

	"gorm.io/gorm"
)

type Contents struct {
	gorm.Model
	Image   string
	Caption string
	UserID  uint
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
