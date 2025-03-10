package utils

import (
    "time"

    "github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("banktest") // In production, use environment variables

func GenerateToken(userID uint, username string) (string, error) {
    claims := jwt.MapClaims{
        "user_id":  userID,
        "username": username,
        // "exp":      time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
        "exp":      time.Now().Add(time.Minute * 5).Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtSecret)
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
    return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return jwtSecret, nil
    })
} 
