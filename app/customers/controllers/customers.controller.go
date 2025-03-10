package controllers

import (
    "banktest/app/customers/models"
	"banktest/app/customers/services"
	// "banktest/app/auth/middleware"
	"banktest/library"
    "encoding/json"
    "net/http"
)

func GetAllUserData(w http.ResponseWriter, r *http.Request) {
	status, msg, result := services.GetAllUserData()
	if !status{
		library.Res_400(w, msg)
		return
	}
	if result.User == nil{
		library.Res_200(w, msg, []string{})
		return
	}
	library.Res_200(w, msg, result.User)
}

func FindCustomer(w http.ResponseWriter, r *http.Request) {
	var user models.CustomerRequest
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        library.Res_400(w, "Invalid request payload")
        return
    }
    defer r.Body.Close()
	// Get user information from context (set by middleware)
    // userID := r.Context().Value(middleware.UserIDKey).(uint)

	status, msg, result := services.GetUserData(user)
	if !status{
		library.Res_400(w, msg)
		return
	}
	if result.User == nil{
		library.Res_200(w, msg, []string{})
		return
	}
	library.Res_200(w, msg, result.User)
}
