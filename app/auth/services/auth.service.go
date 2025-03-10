package services

import (
    "banktest/app/auth/models"
    "banktest/app/auth/utils"
	"banktest/config"

    "golang.org/x/crypto/bcrypt"
)

func Register(user models.User) (bool, string, models.User) {
    // Check if username already exists
    var exists bool
    err := config.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM customers WHERE username=$1)", user.Username).Scan(&exists)
    if err != nil {
        return false,"Database error",user
    }
    if exists {
        return false,"Username already exists",user
    }

    // Hash the password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
       return false,"Failed to hash password",user
    }

    // Insert user into database
    var userID uint
    err = config.DB.QueryRow(
		"INSERT INTO customers (name, username, password, hp, address) VALUES ($1, $2, $3, $4, $5) RETURNING id",
        user.Name, user.Username, string(hashedPassword), user.Hp, user.Address,
    ).Scan(&userID)
    if err != nil {
        return false,"Failed to create user",user
    }

	// Create response without password
    user.ID = userID
    user.Password = "" 

	return true,"User registered successfully",user
}

func Login(loginReq models.LoginRequest) (bool, string, models.LoginResponse) {
    // Find user
	var loginRes models.LoginResponse
    err := config.DB.QueryRow(
        "SELECT id, name, username, password, hp, address FROM customers WHERE username=$1",
        loginReq.Username,
    ).Scan(&loginRes.ID, &loginRes.Name, &loginRes.Username, &loginRes.Password, &loginRes.Hp, &loginRes.Address)
    if err != nil {
        return false,"Invalid Username",loginRes
    }

    // Compare passwords
    err = bcrypt.CompareHashAndPassword([]byte(loginRes.Password), []byte(loginReq.Password))
    if err != nil {
        return false,"Invalid Password",loginRes
    }

    // Generate JWT token
    token, err := utils.GenerateToken(loginRes.ID, loginRes.Username)
    if err != nil {
        return false,"Failed to generate token",loginRes
    }
	loginRes.Token = token
	loginRes.Password = ""

    return true,"login successfully",loginRes
}
