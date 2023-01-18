package handler

import "incrediblefour/features/comment"

type CommentResponse struct {
	ID uint `json:"id"`
	Text string `json:"text"`
}

func ToResponse(data comment.Core) CommentResponse {
	return CommentResponse{
		ID: data.ID,
		Text: data.Text,
	}
}

func CoresToResponse(dataCore comment.Core) CommentResponse {
	return CommentResponse{
		ID: dataCore.ID,
		Text: dataCore.Text,
	}
}
