package handlers

import (
	"encoding/json"
	"net/http"

	"secure-login/middleware"
	"secure-login/storage"
)

func Dashboard(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	// Find user by ID.
	// For this simple project, iterate through the in-memory list.
	user := storage.FindUserByID(userID)

	if user == nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message":  "welcome to dashboard",
		"user_id":  user.ID,
		"username": user.Username,
	})
}
