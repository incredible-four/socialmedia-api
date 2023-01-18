package data

import (
	"incrediblefour/features/comment"

	"gorm.io/gorm"
)

type Comments struct {
	gorm.Model
	Text 	  string
	UserID	  uint
	ContentID uint
}

func ToCore(data Comments) comment.Core {
	return comment.Core{
		ID: data.ID,
		Text: data.Text,
		UserID: data.UserID,
		ContentID: data.ContentID,
	}
}

func CoreToData(data comment.Core) Comments {
	return Comments{
		Model: gorm.Model{ID: data.ID},
		Text: data.Text,
		UserID: data.UserID,
		ContentID: data.ContentID,
	}
}
