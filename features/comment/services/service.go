package services

import (
	"errors"
	"incrediblefour/features/comment"
	"incrediblefour/helper"
	"log"
	"strings"
)

type commentSrv struct {
	data comment.CommentData
}

func New(d comment.CommentData) comment.CommentService {
	return &commentSrv{
		data: d,
	}
}

func (cs *commentSrv) Add(token interface{}, newComment comment.Core, contentID uint) (comment.Core, error) {
	userID := helper.ExtractToken(token)
	if userID <= 0 {
		return comment.Core{}, errors.New("User not found")
	}

	res, err := cs.data.Add(uint(userID), newComment, contentID)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "Comment not found"
		} else {
			msg = "There's a server eror"
		}
		return comment.Core{}, errors.New(msg)
	}

	res.UserID = uint(userID)

	return res, nil
}

func (cs *commentSrv) Delete(token interface{}, commentID uint) error {
	id := helper.ExtractToken(token)

	if id <= 0 {
		return errors.New("Data not found")
	}

	err := cs.data.Delete(uint(id), commentID)
	if err != nil {
		log.Println("Delete query error : ", err.Error())
		return err
	}

	return nil
}