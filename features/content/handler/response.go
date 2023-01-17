package handler

import "incrediblefour/features/content"

type ContentResponse struct {
	ID       uint   `json:"id"`
	Avatar   string `json:"avatar"`
	Username string `json:"username"`
	Image    string `json:"image"`
	Caption  string `json:"caption"`
	UserID   uint   `json:"user_id"`
}

func ToResponse(data content.Core) ContentResponse {
	return ContentResponse{
		ID:       data.ID,
		Avatar:   data.Avatar,
		Username: data.Username,
		Image:    data.Image,
		Caption:  data.Caption,
		UserID:   data.UserID,
	}
}
