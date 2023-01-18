package handler

import (
	"incrediblefour/features/comment"
	"incrediblefour/helper"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type commentControll struct {
	srv comment.CommentService
}

func New(cs comment.CommentService) comment.CommentHandler {
	return &commentControll{
		srv: cs,
	}
}

func (cs *commentControll) Add() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := AddCommentRequest{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "Please input correctly")
		}

		paramID := c.Param("id")
		contentID, err := strconv.Atoi(paramID)
		if err != nil {
			log.Println("Convert id error : ", err.Error())
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Please input number only",
			})
		}

		cnv := ToCore(input)

		res, err := cs.srv.Add(c.Get("user"), *cnv, uint(contentID))
		if err != nil {
			log.Println("Error insert book :  ", err.Error())
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}

		return c.JSON(helper.PrintSuccessReponse(http.StatusCreated, "Added a new comment successfully", ToResponse(res)))
	}
}

func (cs *commentControll) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")

		paramID := c.Param("id")
		commentID, err := strconv.Atoi(paramID)
		if err != nil {
			log.Println("Convert id error : ", err.Error())
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Please input number only",
			})
		}

		err = cs.srv.Delete(token, uint(commentID))
		
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}

		return c.JSON(http.StatusAccepted, "Deleted a comment successfully")
	}
}