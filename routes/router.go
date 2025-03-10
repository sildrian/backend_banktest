package routes

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"banktest/app/auth/middleware"
	customers "banktest/app/customers/controllers"
	auth "banktest/app/auth/controllers"
)

func Router() {
	r := mux.NewRouter()
	// r := router.PathPrefix("/v1").Subrouter()
	r.HandleFunc("/create-customer", auth.Register).Methods("POST")
	r.HandleFunc("/login", auth.Login).Methods("POST")
	r.HandleFunc("/get-customer", middleware.AuthMiddleware(customers.GetAllUserData)).Methods("GET")
	r.HandleFunc("/find-customer", middleware.AuthMiddleware(customers.FindCustomer)).Methods("POST")

	// CORS configuration
	corsMiddleware := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:3000"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization", "Origin"}),
		handlers.AllowCredentials(),
	)

	// Apply middleware
	handler := corsMiddleware(r)

	// Configure server
    srv := &http.Server{
        Handler:      handler,
        Addr:         ":8080",
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
    }

	log.Println("Server starting on port 8080...")
	// log.Fatal(http.ListenAndServe(":8080", handler))
	if err := srv.ListenAndServe(); err != nil {
        log.Fatal("Server failed to start:", err)
    }
}
