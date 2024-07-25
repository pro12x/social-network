package session

import (
	"errors"
	"github.com/google/uuid"
	"net/http"
	"sync"
	"time"
)

type Session struct {
	UserID    uint
	ExpiresAt time.Time
}

var (
	sessionStore    = map[string]*Session{}
	mutex           = &sync.Mutex{}
	sessionDuration = 24 * time.Hour
)

type StoreSessions struct {
	session map[string]uint
	mu      sync.Mutex
}

const sessionName = "session_token"

func NewSessionStore() *StoreSessions {
	return &StoreSessions{
		session: make(map[string]uint),
	}
}

func (s *StoreSessions) StoreSession(token string, userID uint) {
	s.mu.Lock()
	s.session[token] = userID
	s.mu.Unlock()
}

func CreateSession(userID uint) (string, error) {
	sessionToken := uuid.NewString()
	expiresAt := time.Now().Add(sessionDuration)

	mutex.Lock()
	sessionStore[sessionToken] = &Session{
		UserID:    userID,
		ExpiresAt: expiresAt,
	}
	mutex.Unlock()

	return sessionToken, nil
}

func GetSession(token string) (*Session, error) {
	mutex.Lock()
	session, exists := sessionStore[token]
	mutex.Unlock()

	if !exists {
		return nil, errors.New("session not found")
	}
	if session.ExpiresAt.Before(time.Now()) {
		err := DeleteSession(token)
		if err != nil {
			return nil, err
		}
		return nil, errors.New("session expired")
	}

	return session, nil
}

func DeleteSession(token string) error {
	if _, exists := sessionStore[token]; !exists {
		return errors.New("session not found")
	}
	mutex.Lock()
	delete(sessionStore, token)
	mutex.Unlock()
	return nil
}

func (s *StoreSessions) GetUserID(token string) (uint, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	userID, exists := s.session[token]
	return userID, exists
}

func (s *StoreSessions) ClearSession(token string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.session, token)
}

func SetSessionCookie(w http.ResponseWriter, token string) {
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    token,
		Expires:  time.Now().Add(sessionDuration),
		HttpOnly: true,
	})
}

func GetSessionTokenFromRequest(r *http.Request) (string, error) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		return "", err
	}
	return cookie.Value, nil
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
