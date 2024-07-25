package web

import (
	"backend/pkg/dto"
	"backend/pkg/service/impl"
	"backend/pkg/session"
	"backend/pkg/utils"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type UserController struct {
	UserService impl.UserServiceImpl
}

// Register Create new user controller
func (c *UserController) Register(w http.ResponseWriter, r *http.Request) {
	err := utils.Environment()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, os.Getenv("METHOD_NOT_ALLOWED"), http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != os.Getenv("DEFAULT_API_LINK")+"/register" {
		http.Error(w, os.Getenv("NOT_FOUND"), http.StatusNotFound)
		return
	}

	var userDTO dto.UserDTO
	if err := json.NewDecoder(r.Body).Decode(&userDTO); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.UserService.CreateUser(&userDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (c *UserController) Login(w http.ResponseWriter, r *http.Request) {
	err := utils.Environment()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, os.Getenv("METHOD_NOT_ALLOWED"), http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != os.Getenv("DEFAULT_API_LINK")+"/login" {
		http.Error(w, os.Getenv("NOT_FOUND"), http.StatusNotFound)
		return
	}

	var credentials struct {
		Email    string `json:"email" db:"email"`
		Password string `json:"password" db:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userDTO, err := c.UserService.Connection(credentials.Email, credentials.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// I will add the token generation here
	sessionToken, err := c.UserService.CreateSession(userDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	session.SetSessionCookie(w, sessionToken)
	w.WriteHeader(http.StatusOK)
	w.Header().Set(os.Getenv("CONTENT_TYPE"), os.Getenv("APPLICATION_JSON"))
	// json.NewEncoder(w).Encode(userDTO)
	err = json.NewEncoder(w).Encode(map[string]interface{}{
		"token":  sessionToken,
		"status": "success",
		"user":   userDTO,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c *UserController) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	err := utils.Environment()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method != http.MethodPut {
		http.Error(w, os.Getenv("METHOD_NOT_ALLOWED"), http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != os.Getenv("DEFAULT_API_LINK")+"/profile-update/{id}" {
		http.Error(w, os.Getenv("NOT_FOUND"), http.StatusNotFound)
		return
	}

	id, err := utils.ExtractIDFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var userDTO dto.UserDTO
	if err := json.NewDecoder(r.Body).Decode(&userDTO); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.UserService.UpdateProfile(uint(id), &userDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (c *UserController) GetProfile(w http.ResponseWriter, r *http.Request) {
	err := utils.Environment()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, os.Getenv("METHOD_NOT_ALLOWED"), http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != os.Getenv("DEFAULT_API_LINK")+"/profile/{id}" {
		http.Error(w, os.Getenv("NOT_FOUND"), http.StatusNotFound)
		return
	}

	id, err := utils.ExtractIDFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userDTO, err := c.UserService.GetProfile(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set(os.Getenv("CONTENT_TYPE"), os.Getenv("APPLICATION_JSON"))
	err = json.NewEncoder(w).Encode(userDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c *UserController) Follow(w http.ResponseWriter, r *http.Request) {
	err := utils.Environment()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	/*if r.Method != http.MethodPost {
		http.Error(w, os.Getenv("METHOD_NOT_ALLOWED), http.StatusMethodNotAllowed)
		return
	}*/

	if r.URL.Path != os.Getenv("DEFAULT_API_LINK")+"/follow/{id}" {
		http.Error(w, os.Getenv("NOT_FOUND"), http.StatusNotFound)
		return
	}

	followerID, err := utils.ExtractIDFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	followeeID, err := strconv.Atoi(r.URL.Query().Get("followee_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return

	}

	err = c.UserService.Follow(uint(followerID), uint(followeeID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (c *UserController) Unfollow(w http.ResponseWriter, r *http.Request) {
	err := utils.Environment()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	/*if r.Method != http.MethodPost {
		http.Error(w, os.Getenv("METHOD_NOT_ALLOWED), http.StatusMethodNotAllowed)
		return
	}*/

	if r.URL.Path != os.Getenv("DEFAULT_API_LINK")+"/unfollow/{id}" {
		http.Error(w, os.Getenv("NOT_FOUND"), http.StatusNotFound)
		return
	}

	followerID, err := utils.ExtractIDFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	followeeID, err := strconv.Atoi(r.URL.Query().Get("followee_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return

	}

	err = c.UserService.Unfollow(uint(followerID), uint(followeeID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (c *UserController) GetFollowers(w http.ResponseWriter, r *http.Request) {
	err := utils.Environment()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, os.Getenv("METHOD_NOT_ALLOWED"), http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != os.Getenv("DEFAULT_API_LINK")+"/followers/{id}" {
		http.Error(w, os.Getenv("NOT_FOUND"), http.StatusNotFound)
		return
	}

	id, err := utils.ExtractIDFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	followers, err := c.UserService.GetFollowers(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set(os.Getenv("CONTENT_TYPE"), os.Getenv("APPLICATION_JSON"))
	err = json.NewEncoder(w).Encode(followers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c *UserController) IsUserOnline(w http.ResponseWriter, r *http.Request) {
	err := utils.Environment()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, os.Getenv("METHOD_NOT_ALLOWED"), http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != os.Getenv("DEFAULT_API_LINK")+"/is_online" {
		http.Error(w, os.Getenv("NOT_FOUND"), http.StatusNotFound)
		return
	}

	token, err := session.GetSessionTokenFromRequest(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	isOnline, err := c.UserService.IsUserOnline(token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	response := map[string]bool{"is_online": isOnline}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c *UserController) Logout(w http.ResponseWriter, r *http.Request) {
	token, err := session.GetSessionTokenFromRequest(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	err = c.UserService.Logout(token)
	if err != nil {
		return
	}

	// Clear the session cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    "",
		Expires:  time.Unix(0, 0), // Expire the cookie immediately
		HttpOnly: true,
	})

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("Logged out successfully"))
	if err != nil {
		return
	}
}

func (c *UserController) Users(w http.ResponseWriter, r *http.Request) {
	err := utils.Environment()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, os.Getenv("METHOD_NOT_ALLOWED"), http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != os.Getenv("DEFAULT_API_LINK")+"/users" {
		http.Error(w, os.Getenv("NOT_FOUND"), http.StatusNotFound)
		return
	}

	users, err := c.UserService.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set(os.Getenv("CONTENT_TYPE"), os.Getenv("APPLICATION_JSON"))
	// err = json.NewEncoder(w).Encode(users)
	err = json.NewEncoder(w).Encode(map[string]interface{}{
		"users": users,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// RegisterRoutes Register routes
func (c *UserController) RegisterRoutes(routes *http.ServeMux) *http.ServeMux {
	err := utils.Environment()
	if err != nil {
		log.Println(err)
		return routes
	}

	routes.HandleFunc(os.Getenv("DEFAULT_API_LINK")+"/register", c.Register)
	routes.HandleFunc(os.Getenv("DEFAULT_API_LINK")+"/login", c.Login)
	routes.HandleFunc(os.Getenv("DEFAULT_API_LINK")+"/users", c.Users)
	routes.HandleFunc(os.Getenv("DEFAULT_API_LINK")+"/logout", c.Logout)
	routes.HandleFunc(os.Getenv("DEFAULT_API_LINK")+"/profile/{id}", c.GetProfile)
	routes.HandleFunc(os.Getenv("DEFAULT_API_LINK")+"/profile-update/{id}", c.UpdateProfile)
	routes.HandleFunc(os.Getenv("DEFAULT_API_LINK")+"/follow", c.Follow)
	routes.HandleFunc(os.Getenv("DEFAULT_API_LINK")+"/unfollow", c.Unfollow)
	routes.HandleFunc(os.Getenv("DEFAULT_API_LINK")+"/followers/{id}", c.GetFollowers)
	routes.HandleFunc(os.Getenv("DEFAULT_API_LINK")+"/is_online", c.IsUserOnline)

	return routes
}
