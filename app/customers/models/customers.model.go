package models

import (
	auth_model "banktest/app/auth/models"
)

type CustomerRequest struct {
    Username string `json:"search"`
}

type CustomerResponse struct {
	User []auth_model.User `json:"user,omitempty"`
}
