package services

import (
	"errors"
	"incrediblefour/features/content"
	"incrediblefour/helper"
	"log"
	"mime/multipart"
	"strings"

	"github.com/go-playground/validator/v10"
)

type contentSrv struct {
	data content.ContentData
	vld  *validator.Validate
}

func New(c content.ContentData) content.ContentService {
	return &contentSrv{
		data: c,
		vld:  validator.New(),
	}
}

func (cs *contentSrv) Add(formHeader multipart.FileHeader, token interface{}, newContent content.Core) (content.Core, error) {
	userID := helper.ExtractToken(token)

	if userID <= 0 {
		return content.Core{}, errors.New("user not found")
	}
	err := cs.vld.Struct(newContent)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			log.Println(err)
		}
		return content.Core{}, errors.New("invalid input")
	}

	if formHeader.Size > 5000000 {
		return content.Core{}, errors.New("file size is too big")
	}

	formFile, err := formHeader.Open()
	if err != nil {
		return content.Core{}, errors.New("open formheader error")
	}

	if !helper.TypeFile(formFile) {
		return content.Core{}, errors.New("use jpg or png type file")
	}
	defer formFile.Close()
	formFile, _ = formHeader.Open()
	uploadUrl, err := helper.NewMediaUpload().ContentUpload(helper.Content{Content: formFile})

	if err != nil {
		return content.Core{}, errors.New("server error")
	}

	newContent.Image = uploadUrl

	res, err := cs.data.Add(uint(userID), newContent)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "user not found"
		} else {
			msg = "unable to process the data"
		}
		return content.Core{}, errors.New(msg)
	}
	res.UserID = uint(userID)

	return res, nil
}

func (cs *contentSrv) ContentDetail(contentID uint) (content.Core, error) {

	res, err := cs.data.ContentDetail(contentID)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "content not found"
		} else {
			msg = "unable to process the data"
		}
		return content.Core{}, errors.New(msg)
	}
	return res, nil
}

func (cs *contentSrv) ContentList() ([]content.Core, error) {
	res, err := cs.data.ContentList()
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "content not found"
		} else {
			msg = "unable to process the data"
		}
		return []content.Core{}, errors.New(msg)
	}
	return res, nil
}

func (cs *contentSrv) Update(token interface{}, contentID uint, updatedContent content.Core) (content.Core, error) {
	id := helper.ExtractToken(token)

	if id <= 0 {
		return content.Core{}, errors.New("data not found")
	}

	res, err := cs.data.Update(uint(id), contentID, updatedContent)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "Failed to update, no new record or data not found"
		} else if strings.Contains(err.Error(), "Unauthorized") {
			msg = "Unauthorized request"
		} else {
			msg = "unable to process the data"
		}
		return content.Core{}, errors.New(msg)
	}

	res.ID = contentID
	res.UserID = uint(id)

	return res, nil
}

func (cs *contentSrv) Delete(token interface{}, contentID uint) error {
	id := helper.ExtractToken(token)

	if id <= 0 {
		return errors.New("data not found")
	}

	err := cs.data.Delete(uint(id), contentID)

	if err != nil {
		log.Println("delete query error", err.Error())
		return err
	}

	return nil
}

func (cs *contentSrv) GetProfile(username string) (interface{}, error) {
	res, err := cs.data.GetProfile(username)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "content not found"
		} else {
			msg = "unable to process the data"
		}
		return nil, errors.New(msg)
	}
	return res, nil
}
