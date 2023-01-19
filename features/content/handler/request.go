package handler

import "incrediblefour/features/content"

type AddContentRequest struct {
	Image   string `json:"image" form:"image"`
	Caption string `json:"caption" form:"caption"`
}

type UpdateContentRequest struct {
	Caption string `json:"caption" form:"caption"`
}

func ToCore(data interface{}) *content.Core {
	res := content.Core{}

	switch data.(type) {
	case AddContentRequest:
		cnv := data.(AddContentRequest)
		res.Image = cnv.Image
		res.Caption = cnv.Caption
	case UpdateContentRequest:
		cnv := data.(UpdateContentRequest)
		res.Caption = cnv.Caption
	default:
		return nil
	}

	return &res
}
