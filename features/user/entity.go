package user

import (
	"github.com/labstack/echo/v4"
)

type Core struct {
	ID uint
	Avatar string
	Banner string
	Name string
	Username string
	Email string
	Password string
}

type UserHandler interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
	Profile() echo.HandlerFunc
	GetProfile() echo.HandlerFunc
	Update() echo.HandlerFunc
	Deactivate() echo.HandlerFunc
}

type UserService interface {
	Register(newUser Core) (Core, error)
	Login(username, password string) (string, Core, error)
	Profile(username string) (Core, error)
	GetProfile(token interface{}) (Core, error)
	Update(token interface{}, updatedProfile Core) (Core, error)
	Deactivate(token interface{}) (error)
}

type UserData interface {
	Register(newUser Core) (Core, error)
	Login(username string) (Core, error)
	Profile(username string) (Core, error)
	GetProfile(id uint) (Core, error)
	Update(updatedProfile Core) (Core, error)
	Deactivate(id uint) (error)
}