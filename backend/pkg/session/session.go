package session

import (
	"net/http"
	"sync"
	"time"
)

type SessionStore struct {
	session map[string]uint
	mu      sync.Mutex
}

const sessionName = "session_token"

func NewSessionStore() *SessionStore {
	return &SessionStore{
		session: make(map[string]uint),
	}
}

func (s *SessionStore) StoreSession(token string, userID uint) {
	s.mu.Lock()
	s.session[token] = userID
	s.mu.Unlock()
}

func (s *SessionStore) GetUserID(token string) (uint, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	userID, exists := s.session[token]
	return userID, exists
}

func (s *SessionStore) ClearSession(token string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.session, token)
}

func SetSessionCookie(w http.ResponseWriter, token string) {
	http.SetCookie(w, &http.Cookie{
		Name:     sessionName,
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		Path:     "/",
	})
}

func GetSessionCookie(r *http.Request) (string, error) {
	cookie, err := r.Cookie(sessionName)
	if err != nil {
		return "", err
	}
	return cookie.Value, nil
}

func ClearSessionCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:    sessionName,
		Value:   "",
		Expires: time.Now().Add(-1 * time.Hour),
		Path:    "/",
	})
}
