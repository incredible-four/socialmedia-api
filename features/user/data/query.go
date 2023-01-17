package data

import (
	"errors"
	"incrediblefour/features/user"
	"log"
	"strings"

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

func (uq *userQuery) isDuplicate(username, email string) (bool) {
	res := Users{}
	if err := uq.db.Where("username = ? OR email = ?", username, email).First(&res).Error; err != nil {
		log.Println("Check data error : ", err.Error())
		if strings.Contains(err.Error(), "not found") {
			return false
		}
		
	}
	return true
}

func (uq *userQuery) Register(newUser user.Core) (error) {
	cnv := CoreToData(newUser)
	isDupl := uq.isDuplicate(cnv.Username, cnv.Email)

	if isDupl != false {
		log.Println("Duplicated data when register")
		return errors.New("Unable to register, duplicated data")
	}

	err := uq.db.Create(&cnv).Error
	if err != nil {
		log.Println("Create query error : ", err.Error())
		return err
	}

	newUser.ID = cnv.ID

	return nil
}

func (uq *userQuery) Login(username string) (user.Core, error) {
	res := Users{}

	if err := uq.db.Where("username = ?", username).First(&res).Error; err != nil {
		log.Println("login query error : ", err.Error())
		return user.Core{}, errors.New("Data not found")
	}

	return ToCore(res), nil
}

func (uq *userQuery) Profile(id uint) (user.Core, error) {
	res := Users{}
	if err := uq.db.Where("id = ?", id).First(&res).Error; err != nil {
		log.Println("Get profile by ID query error : ", err.Error())
		return user.Core{}, err
	}
	return ToCore(res), nil
}

func (uq *userQuery) Update(updatedProfile user.Core) (user.Core, error) {
	cnvUpdated := CoreToData(updatedProfile)
	qry := uq.db.Model(Users{}).Where("id = ?", cnvUpdated.ID).Updates(cnvUpdated)
	err := qry.Error

	affRow := qry.RowsAffected

	if affRow <= 0 {
		log.Println("No rows affected")
		return user.Core{}, errors.New("Failed to update, no new record or data not found")
	}

	if err != nil {
		log.Println("update query error : ", err.Error())
		return user.Core{}, errors.New("Unable to update profile")
	}

	return ToCore(cnvUpdated), nil
}

func (uq *userQuery) Deactivate(id uint) (error) {
	qryDelete := uq.db.Delete(&Users{}, id)

	affRow := qryDelete.RowsAffected

	if affRow <= 0 {
		log.Println("No rows affected")
		return errors.New("Failed to delete, data not found")
	}

	return nil
}