package handler

import (
	"incrediblefour/features/user"
	"net/http"
	"strings"
)

type UserResponse struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
}

func ToResponse(data user.Core) UserResponse {
	return UserResponse{
		ID: data.ID,
		Name: data.Name,
		Username: data.Username,
		Email: data.Email,
		Password: data.Password,
	}
}

func PrintSuccessReponse(code int, message string, data ...interface{}) (int, interface{}) {
	resp := map[string]interface{}{}
	if len(data) < 2 {
		resp["data"] = ToResponse(data[0].(user.Core))
	} else {
		resp["data"] = ToResponse(data[0].(user.Core))
		resp["token"] = data[1].(string)
	}

	if message != "" {
		resp["message"] = message
	}

	return code, resp
}

func PrintErrorResponse(msg string) (int, interface{}) {
	resp := map[string]interface{}{}
	code := -1
	if msg != "" {
		resp["message"] = msg
	}

	if strings.Contains(msg, "server") {
		code = http.StatusInternalServerError
	} else if strings.Contains(msg, "format") {
		code = http.StatusBadRequest
	} else if strings.Contains(msg, "match") {
		code = http.StatusUnauthorized
	} else {
		code = http.StatusNotFound
	} 

	return code, resp
}