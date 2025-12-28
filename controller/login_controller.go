package controller

import (
	utils "belajargolang/services"
	"encoding/json"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {

	// contoh hasil validasi user
	userID := 1
	email := "admin@mail.com"

	token, err := utils.GenerateToken(userID, email)
	if err != nil {
		http.Error(w, "Failed generate token", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"token": token,
	}

	json.NewEncoder(w).Encode(response)
}

func Profile(w http.ResponseWriter, r *http.Request) {

	userID := r.Context().Value("user_id")
	email := r.Context().Value("email")

	response := map[string]interface{}{
		"user_id": userID,
		"email":   email,
	}

	json.NewEncoder(w).Encode(response)
}
