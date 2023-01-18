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

type AllComments struct {
	gorm.Model
	ID        uint
	Avatar    string
	Username  string
	Text   	  string
	UserID    uint
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

func (dataModel *Comments) ModelsToCore() comment.Core {
	return comment.Core{
		ID: dataModel.ID,
		Text: dataModel.Text,
		UserID: dataModel.UserID,
		ContentID: dataModel.ContentID,
	}
}

func ListToCore(data []Comments) []comment.Core {
	var dataCore []comment.Core
	for _, v := range data {
		dataCore = append(dataCore, v.ModelsToCore())
	}
	return dataCore
}

func (dataModel *AllComments) AllModelsToCore() comment.Core {
	return comment.Core{
		ID: dataModel.ID,
		Avatar: dataModel.Avatar,
		Username: dataModel.Username,
		Text: dataModel.Text,
		UserID: dataModel.UserID,
		ContentID: dataModel.ContentID,
	}
}

func AllListToCore(data []AllComments) []comment.Core {
	var dataCore []comment.Core
	for _, v := range data {
		dataCore = append(dataCore, v.AllModelsToCore())
	}
	return dataCore
}

