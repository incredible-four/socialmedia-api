package handler

import "incrediblefour/features/user"

type RegisterRequest struct {
	Name string `json:"name" form:"name"`
	Username string `json:"username" form:"username"`
	Email string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UpdateRequest struct {
	Avatar string `json:"avatar" form:"avatar"`
	Banner string `json:"banner" form:"banner"`
	Name string `json:"name" form:"name"`
	Username string `json:"username" form:"username"`
	Bio string `json:"bio" form:"bio"`
	Email string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type LoginRequest struct {
	Username string `json:"username"  form:"username"`
	Password string `json:"password" form:"password"`
}

func ToCore(data interface{}) *user.Core {
	res := user.Core{}

	switch data.(type) {
	case RegisterRequest:
		cnv := data.(RegisterRequest)
		res.Name = cnv.Name
		res.Username = cnv.Username
		res.Email = cnv.Email
		res.Password = cnv.Password
	case UpdateRequest:
		cnv := data.(UpdateRequest)
		res.Avatar = cnv.Avatar
		res.Banner = cnv.Banner
		res.Name = cnv.Name
		res.Username = cnv.Username
		res.Bio = cnv.Bio
		res.Email = cnv.Email
		res.Password = cnv.Password
	default:
		return nil
	}

	return &res
}