package storage

import (
	"secure-login/models"
	"sync"
)

var (
	users = []models.User{}
	mu    sync.RWMutex
)

func AddUser(user models.User) {
	mu.Lock()
	defer mu.Unlock()

	users = append(users, user)
}

func FindUser(username string) *models.User {
	mu.RLock()
	defer mu.RUnlock()

	for _, user := range users {
		if user.Username == username {
			u := user
			return &u
		}
	}

	return nil
}

func FindUserByID(id string) *models.User {
	mu.RLock()
	defer mu.RUnlock()

	for _, user := range users {
		if user.ID == id {
			u := user
			return &u
		}
	}

	return nil
}
