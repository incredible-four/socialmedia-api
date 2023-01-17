package services

import (
	"errors"
	"incrediblefour/features/user"
	"log"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	qry user.UserData
}

func New(ud user.UserData) user.UserService {
	return &userUseCase {
		qry: ud,
	}
}

func (uuc *userUseCase) Register(newUser user.Core) (error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("generate bcrypt error : ", err.Error())
		return errors.New("Unable to process password")
	}
	newUser.Password = string(hashed)
	err = uuc.qry.Register(newUser)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "duplicated") {
			msg = "Username or email already exist"
		} else if strings.Contains(err.Error(), "password") {
			msg = "There is a problem with the server"
		} else {
			msg = "There is a problem with the server"
		}
		return errors.New(msg)
	}
	return nil
}