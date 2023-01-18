package comment

import "github.com/labstack/echo/v4"

type Core struct {
	ID uint
	Text string
	UserID uint
	ContentID uint
}

type CommentHandler interface {
	Add() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type CommentService interface {
	Add(token interface{}, newComment Core, contentID uint) (Core, error)
	Delete(token interface{}, commentID uint) (error)
}

type CommentData interface {
	Add(userID uint, newComment Core, contentID uint) (Core, error)
	Delete(userID uint, commentID uint) (error)
}