package main

import (
	"fmt"
	"net/http"

	"github.com/Greeshmanth-Pulicallu/login-api/db"
	"github.com/Greeshmanth-Pulicallu/login-api/handlers"
	"github.com/Greeshmanth-Pulicallu/login-api/models"
)

func main() {
	// connect to db
	db.Connect()
	db.DB.AutoMigrate(&models.User{})

	fmt.Println("Server running on :8080")
	http.HandleFunc("/login", handlers.LoginHandler)
	http.ListenAndServe(":8080", nil)
}
