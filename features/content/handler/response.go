package handler

import (
	"incrediblefour/features/content"
	"time"
)

type ContentResponse struct {
	ID       uint   `json:"id"`
	Avatar   string `json:"avatar"`
	Username string `json:"username"`
	Image    string `json:"image"`
	Caption  string `json:"caption"`
	CreatedAt time.Time `json:"created_at"`
}

type AddResponse struct {
	ID       uint   `json:"id"`
	Image    string `json:"image"`
	Caption  string `json:"caption"`
	CreatedAt time.Time `json:"created_at"`
}

func ToAddResponse(data content.Core) AddResponse {
	return AddResponse{
		ID:       data.ID,
		Image:    data.Image,
		Caption:  data.Caption,
		CreatedAt: time.Now(),
	}
}

func ToResponse(data content.Core) ContentResponse {
	return ContentResponse{
		ID:       data.ID,
		Avatar:   data.Avatar,
		Username: data.Username,
		Image:    data.Image,
		Caption:  data.Caption,
	}
}

func CoresToResponse(dataCore content.Core) ContentResponse {
	return ContentResponse{
		ID:       dataCore.ID,
		Avatar:   dataCore.Avatar,
		Username: dataCore.Username,
		Image:    dataCore.Image,
		Caption:  dataCore.Caption,
	}
}

func ListCoreToResp(data []content.Core) []ContentResponse {
	var dataResp []ContentResponse
	for _, v := range data {
		dataResp = append(dataResp, CoresToResponse(v))
	}
	return dataResp
}
