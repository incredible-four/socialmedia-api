package data

import (
	"incrediblefour/features/content"
	"log"

	"gorm.io/gorm"
)

type contentData struct {
	db *gorm.DB
}

func New(db *gorm.DB) content.ContentData {
	return &contentData{
		db: db,
	}
}

func (cd *contentData) Add(userID uint, newContent content.Core) (content.Core, error) {
	cnv := CoreToData(newContent)
	cnv.UserID = uint(userID)
	err := cd.db.Create(&cnv).Error
	if err != nil {
		return content.Core{}, err
	}

	newContent.ID = cnv.ID

	return newContent, nil
}

func (cd *contentData) MyContent(userID uint) ([]content.Core, error) {
	res := []Contents{}
	if err := cd.db.Where("user_id = ?", userID).Find(&res).Error; err != nil {
		log.Println("Get User Content by User ID query error : ", err.Error())
		return []content.Core{}, err
	}

	return ListToCore(res), nil
}
