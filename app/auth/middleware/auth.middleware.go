package middleware

import (
    "banktest/app/auth/utils"
    "context"
    "net/http"
    "strings"

    "github.com/golang-jwt/jwt/v5"
)

type ContextKey string

const UserIDKey ContextKey = "userID"
const UsernameKey ContextKey = "username"

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Get token from Authorization header
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            http.Error(w, "Authorization header is required", http.StatusUnauthorized)
            return
        }

        // Check if the header has the Bearer prefix
        bearerToken := strings.Split(authHeader, " ")
        if len(bearerToken) != 2 || strings.ToLower(bearerToken[0]) != "bearer" {
            http.Error(w, "Invalid authorization format. Use: Bearer <token>", http.StatusUnauthorized)
            return
        }

        // Validate token
        token, err := utils.ValidateToken(bearerToken[1])
        if err != nil {
            http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
            return
        }

        // Extract claims
        claims, ok := token.Claims.(jwt.MapClaims)
        if !ok || !token.Valid {
            http.Error(w, "Invalid token claims", http.StatusUnauthorized)
            return
        }

        // Add user information to request context
        ctx := context.WithValue(r.Context(), UserIDKey, uint(claims["user_id"].(float64)))
        ctx = context.WithValue(ctx, UsernameKey, claims["username"].(string))
        
        // Call next handler with new context
        next.ServeHTTP(w, r.WithContext(ctx))
    }
} 
