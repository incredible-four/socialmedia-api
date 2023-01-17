package handler

import (
	"incrediblefour/features/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

type userControll struct {
	srv user.UserService
}

func New(srv user.UserService) user.UserHandler {
	return &userControll {
		srv: srv,
	}
}

func (uc *userControll) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := RegisterRequest{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "Please input correctly")
		}

		err := uc.srv.Register(*ToCore(input))
		if err != nil {
			return c.JSON(PrintErrorResponse(err.Error()))
		}

		return c.JSON(http.StatusCreated, "Registered a new account successfully")
	}
}