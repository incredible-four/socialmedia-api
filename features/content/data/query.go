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

func (cd *contentData) MyContent(contentID uint) ([]content.Core, error) {
	res := []Contents{}
	if err := cd.db.Where("content_id = ?", contentID).Find(&res).Error; err != nil {
		log.Println("Get User Content by User ID query error : ", err.Error())
		return []content.Core{}, err
	}

	return ListToCore(res), nil
}

func (cd *contentData) ContentList() ([]content.Core, error) {
	res := []AllContents{}
	if err := cd.db.Table("contents").Joins("JOIN users ON users.id = content.user_id").Select("contents.id, contents.avatar, contents.username, contents.image, contents.caption, users.username as Owner").Find(&res).Error; err != nil {
		log.Println("get all content query error : ", err.Error())
		return []content.Core{}, err
	}
	return AllListToCore(res), nil
}
