package data

import (
	"incrediblefour/features/comment/data"
	"incrediblefour/features/content"
	"time"

	"gorm.io/gorm"
)

type Contents struct {
	gorm.Model
	Avatar   string
	Username string
	Image    string
	Caption  string
	UserID   uint
	CreatedAt time.Time
	Comment  []data.Comments `gorm:"foreignkey:ContentID"`
}

type AllContents struct {
	ID        uint
	Avatar    string
	Username  string
	Image     string
	Caption   string
	CreatedAt time.Time
	Comments   []data.Comments `gorm:"foreignkey:ContentID"`
}

func ToCore(data Contents) content.Core {
	return content.Core{
		ID:       data.ID,
		Avatar:   data.Avatar,
		Username: data.Username,
		Image:    data.Image,
		Caption:  data.Caption,
		UserID:   data.UserID,
		CreatedAt: data.CreatedAt,
	}
}

func CoreToData(data content.Core) Contents {
	return Contents{
		Model:    gorm.Model{ID: data.ID},
		Avatar:   data.Avatar,
		Username: data.Username,
		Image:    data.Image,
		Caption:  data.Caption,
		UserID:   data.UserID,
		CreatedAt: data.CreatedAt,
	}
}

func (dataModel *Contents) ModelsToCore() content.Core {
	return content.Core{
		ID:       dataModel.ID,
		Avatar:   dataModel.Avatar,
		Username: dataModel.Username,
		Image:    dataModel.Image,
		Caption:  dataModel.Caption,
		UserID:   dataModel.UserID,
		CreatedAt: dataModel.CreatedAt,
	}
}

func ListToCore(data []Contents) []content.Core {
	var dataCore []content.Core
	for _, v := range data {
		dataCore = append(dataCore, v.ModelsToCore())
	}
	return dataCore
}

func (dataModel *AllContents) AllModelsToCore() content.Core {
	return content.Core{
		ID:       dataModel.ID,
		Avatar:   dataModel.Avatar,
		Username: dataModel.Username,
		Image:    dataModel.Image,
		Caption:  dataModel.Caption,
		CreatedAt: dataModel.CreatedAt,
	}
}

func AllListToCore(data []AllContents) []content.Core {
	var dataCore []content.Core
	for _, v := range data {
		dataCore = append(dataCore, v.AllModelsToCore())
	}
	return dataCore
}
