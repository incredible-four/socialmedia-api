package handler

import (
	"incrediblefour/features/content"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type contentHandle struct {
	srv content.ContentService
}

func New(cs content.ContentService) content.ContentHandler {
	return &contentHandle{
		srv: cs,
	}
}

func (ch *contentHandle) Add() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := AddContentRequest{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "invalid input")
		}

		cnv := ToCore(input)

		res, err := ch.srv.Add(c.Get("user"), *cnv)
		if err != nil {
			log.Println("error post content : ", err.Error())
			return c.JSON(http.StatusInternalServerError, "unable to process the data")
		}
		return c.JSON(http.StatusCreated, ToResponse(res))
	}
}
