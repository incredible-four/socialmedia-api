package content

import "github.com/labstack/echo/v4"

type Core struct {
	ID      uint
	Image   string
	Caption string
	UserID  uint
}

type ContentHandler interface {
	Add() echo.HandlerFunc
	ContentList() echo.HandlerFunc
	MyContent() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
}
