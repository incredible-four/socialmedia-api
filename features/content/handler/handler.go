package handler

import (
	"incrediblefour/features/content"
	"incrediblefour/helper"
	"log"
	"net/http"
	"strconv"

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
		return c.JSON(helper.PrintSuccessReponse(http.StatusCreated, "success post content", ToResponse(res)))
	}
}

func (ch *contentHandle) MyContent() echo.HandlerFunc {
	return func(c echo.Context) error {
		paramID := c.Param("id")
		contentID, err := strconv.Atoi(paramID)
		if err != nil {
			log.Println("convert id error : ", err.Error())
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "invalid input",
			})
		}
		res, err := ch.srv.MyContent(uint(contentID))
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}

		return c.JSON(helper.PrintSuccessReponse(http.StatusCreated, "success get user content", ListCoreToResp(res)))
	}
}
