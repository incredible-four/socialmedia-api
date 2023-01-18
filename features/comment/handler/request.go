package handler

import "incrediblefour/features/comment"

type AddCommentRequest struct {
	Text string `json:"text"`
}

func ToCore(data interface{}) *comment.Core {
	res := comment.Core{}

	cnv := data.(AddCommentRequest)
	res.Text = cnv.Text

	return &res
}