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

type ContentService interface {
	Add(token interface{}, newContent Core) (Core, error)
	ContentList() ([]Core, error)
	MyContent(token interface{}) ([]Core, error)
	Update(token interface{}, contentID uint, updatedContent Core) (Core, error)
	Delete(token interface{}, contentID uint) (Core, error)
}

type ContentData interface {
	Add(userID uint, newContent Core) (Core, error)
	ContentList() ([]Core, error)
	MyContent(userID uint) ([]Core, error)
	Update(userID uint, contentID uint, updatedContent Core) (Core, error)
	Delete(userID uint, contentID uint) (Core, error)
}
