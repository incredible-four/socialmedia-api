package data

import (
	"errors"
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
	if err := cd.db.Where("id = ?", contentID).Find(&res).Error; err != nil {
		log.Println("Get User Content by User ID query error : ", err.Error())
		return []content.Core{}, err
	}

	return ListToCore(res), nil
}

func (cd *contentData) ContentList() ([]content.Core, error) {
	res := []AllContents{}
	if err := cd.db.Table("contents").Joins("JOIN users ON users.id = contents.user_id").Select("contents.id, contents.avatar, contents.username, contents.image, contents.caption, users.username as Owner").Find(&res).Error; err != nil {
		log.Println("get all content query error : ", err.Error())
		return []content.Core{}, err
	}
	return AllListToCore(res), nil
}

func (cd *contentData) Update(userID uint, contentID uint, updatedContent content.Core) (content.Core, error) {
	getID := Contents{}
	err := cd.db.Where("id = ?", contentID).First(&getID).Error

	if err != nil {
		log.Println("get content error : ", err.Error())
		return content.Core{}, err
	}

	if getID.UserID != userID {
		log.Println("Unauthorized request")
		return content.Core{}, errors.New("Unauthorized request")
	}

	cnv := CoreToData(updatedContent)
	qry := cd.db.Where("id = ?", contentID).Updates(&cnv)
	if qry.RowsAffected <= 0 {
		log.Println("update content query error : data not found")
		return content.Core{}, errors.New("not found")
	}

	if err := qry.Error; err != nil {
		log.Println("update content query error : ", err.Error())
	}
	return updatedContent, nil
}

func (cd *contentData) Delete(userID uint, contentID uint) error {
	getID := Contents{}
	err := cd.db.Where("id = ?", contentID).First(&getID).Error

	if err != nil {
		log.Println("get content error : ", err.Error())
		return errors.New("failed to get content data")
	}

	if getID.UserID != userID {
		log.Println("unauthorized request")
		return errors.New("inauthorized request")
	}

	qryDelete := cd.db.Delete(&Contents{}, contentID)

	affRow := qryDelete.RowsAffected

	if affRow <= 0 {
		log.Println("No rows affected")
		return errors.New("failed to delete user content, data not found")
	}

	return nil
}
