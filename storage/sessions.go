package storage

import "sync"

var (
	sessions  = make(map[string]string)
	sessionMu sync.RWMutex
)

func CreateSession(sessionID, userID string) {
	sessionMu.Lock()
	defer sessionMu.Unlock()

	sessions[sessionID] = userID
}

func GetUserID(sessionID string) (string, bool) {
	sessionMu.RLock()
	defer sessionMu.RUnlock()

	userID, exists := sessions[sessionID]
	return userID, exists
}

func DeleteSession(sessionID string) {
	sessionMu.Lock()
	defer sessionMu.Unlock()

	delete(sessions, sessionID)
}
