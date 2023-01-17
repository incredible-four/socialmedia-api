package handler

import "incrediblefour/features/content"

type AddContentRequest struct {
	Image   string `json:"image"`
	Caption string `json:"caption"`
}

func ToCore(data interface{}) *content.Core {
	res := content.Core{}

	switch data.(type) {
	case AddContentRequest:
		cnv := data.(AddContentRequest)
		res.Image = cnv.Image
		res.Caption = cnv.Caption
	default:
		return nil
	}

	return &res
}
