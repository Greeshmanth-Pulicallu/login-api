package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Greeshmanth-Pulicallu/login-api/db"
	"github.com/Greeshmanth-Pulicallu/login-api/models"
	"github.com/Greeshmanth-Pulicallu/login-api/utils"
	"golang.org/x/crypto/bcrypt"
)

// Sends JWT if credentials are valid, else 401.
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	var user models.User
	fmt.Println(req)
	result := db.DB.Where("id = ?", req.ID).First(&user)
	if result.Error != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	fmt.Printf("user found %v\n", user.PasswordHash)

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		fmt.Printf("Unable to GenerateJWT %v\n", err)
	}

	resp := LoginResponse{Token: token}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
