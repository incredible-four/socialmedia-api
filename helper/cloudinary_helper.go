package helper

import (
	"context"
	config "incrediblefour/config"
	"io"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/go-playground/validator/v10"
)

func ImageUploadHelper(input interface{}) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//create cloudinary instance
	cld, err := cloudinary.NewFromParams(config.EnvCloudName(), config.EnvCloudAPIKey(), config.EnvCloudAPISecret())
	if err != nil {
		return "", err
	}

	//upload file
	uploadParam, err := cld.Upload.Upload(ctx, input, uploader.UploadParams{Folder: config.EnvCloudUploadFolder()})
	if err != nil {
		return "", err
	}
	return uploadParam.SecureURL, nil
}

func TypeFile(test multipart.File) bool {
	fileByte, _ := io.ReadAll(test)
	fileType := http.DetectContentType(fileByte)
	if fileType == "image/png" || fileType == "image/jpeg" {
		return true
	}
	return false
}

type File struct {
	File multipart.File `json:"file,omitempty" validate:"required"`
	// Banner multipart.File `json:"banner,omitempty" validate:"required"`
}

type Banner struct {
	Banner multipart.File `json:"banner,omitempty" validate:"required"`
}

type Url struct {
	Url string `json:"url,omitempty" validate:"required"`
}

var (
	validate = validator.New()
)

type mediaUpload interface {
	FileUpload(file File) (string, error)
	RemoteUpload(url Url) (string, error)
	BannerUpload(banner Banner) (string, error)
}

func (*media) RemoteUpload(url Url) (string, error) {
	//validate
	err := validate.Struct(url)
	if err != nil {
		return "", err
	}

	//upload
	uploadUrl, errUrl := ImageUploadHelper(url.Url)
	if errUrl != nil {
		return "", err
	}
	return uploadUrl, nil
}

type media struct{}

func NewMediaUpload() mediaUpload {
	return &media{}
}
func (*media) BannerUpload(banner Banner) (string, error) {
	//validate
	err := validate.Struct(banner)
	if err != nil {
		return "", err
	}

	//upload
	uploadUrl, err := ImageUploadHelper(banner.Banner)
	if err != nil {
		return "", err
	}
	return uploadUrl, nil
}

func (*media) FileUpload(file File) (string, error) {
	//validate
	err := validate.Struct(file)
	if err != nil {
		return "", err
	}

	//upload
	uploadUrl, err := ImageUploadHelper(file.File)
	if err != nil {
		return "", err
	}
	return uploadUrl, nil
}