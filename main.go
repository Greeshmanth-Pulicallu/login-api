package main

import (
	"github.com/Greeshmanth-Pulicallu/login-api/db"
	"github.com/Greeshmanth-Pulicallu/login-api/models"
	"github.com/Greeshmanth-Pulicallu/login-api/router"
)

func main() {
	// connect to db
	db.Connect()
	db.DB.AutoMigrate(&models.User{})

	r := router.SetupRouter()
	r.Run(":8080")
}
