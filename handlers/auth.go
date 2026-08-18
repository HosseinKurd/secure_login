package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"

	"secure-login/models"
	"secure-login/security"
	"secure-login/storage"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	var credentials Credentials

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	if credentials.Username == "" || credentials.Password == "" {
		http.Error(w, "username and password required", http.StatusBadRequest)
		return
	}

	if storage.FindUser(credentials.Username) != nil {
		http.Error(w, "registration failed", http.StatusBadRequest)
		return
	}

	hash, err := security.HashPassword(credentials.Password)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	user := models.User{
		ID:       uuid.NewString(),
		Username: credentials.Username,
		Password: hash,
	}

	storage.AddUser(user)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "user registered",
	})
}

func Login(w http.ResponseWriter, r *http.Request) {
	var credentials Credentials

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	user := storage.FindUser(credentials.Username)

	if user == nil || !security.CheckPassword(user.Password, credentials.Password) {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	sessionID, err := security.GenerateSessionID()
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	storage.CreateSession(sessionID, user.ID)

	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    sessionID,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
	})

	json.NewEncoder(w).Encode(map[string]string{
		"message": "login successful",
	})
}

func Logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err == nil {
		storage.DeleteSession(cookie.Value)
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    "",
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
	})

	json.NewEncoder(w).Encode(map[string]string{
		"message": "logout successful",
	})
}
