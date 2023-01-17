package services

import (
	"incrediblefour/features/content"

	"github.com/go-playground/validator/v10"
)

type contentSrv struct {
	data content.ContentData
	vld  *validator.Validate
}

func New(c content.ContentData) content.ContentService {
	return &contentSrv{
		data: d,
		vld:  validator.New(),
	}
}
