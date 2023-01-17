package data

import (
	"errors"
	"incrediblefour/features/user"
	"log"

	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.UserData {
	return &userQuery {
		db: db,
	}
}

func (uq *userQuery) isDuplicate(username, email string) (bool, error) {
	res := Users{}
	if err := uq.db.Where("username = ? OR email = ?", username, email).First(&res).Error; err != nil {
		log.Println("Check duplicate error : ", err.Error())
		return true, err
	}

	if res.Username != "" {
		return true, nil
	}

	return false, nil
}

func (uq *userQuery) Register(newUser user.Core) (error) {
	cnv := CoreToData(newUser)
	isDupl, err := uq.isDuplicate(cnv.Username, cnv.Email)
	
	if err != nil {
		log.Println("Check duplicate error : ", err.Error())
	}

	if isDupl != false {
		log.Println("Duplicated data when register")
		return errors.New("Unable to register, duplicated data")
	}

	err = uq.db.Create(&cnv).Error
	if err != nil {
		log.Println("Create query error : ", err.Error())
		return err
	}

	newUser.ID = cnv.ID

	return nil
}