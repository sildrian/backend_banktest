package controllers

import (
    "banktest/app/auth/models"
	"banktest/app/auth/services"
	"banktest/library"
    "encoding/json"
    "net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        library.Res_400(w, "Invalid request payload")
        return
    }
    defer r.Body.Close()

	status, msg, result := services.Register(user)
	if !status{
		library.Res_400(w, msg)
		return
	}
	library.Res_200(w, msg, result)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var loginReq models.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
        library.Res_400(w, "Invalid request payload")
        return
    }
	defer r.Body.Close()

	status, msg, result := services.Login(loginReq)
	if !status{
		library.Res_400(w, msg)
		return
	}
	library.Res_200(w, msg, result)
}
