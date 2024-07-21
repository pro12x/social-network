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
		http.Error(w, err.Error(), http.StatusInternalServerError)
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

	if r.URL.Path != os.Getenv("DEFAULT_API_LINK")+"/profile/{id}" {
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

// RegisterRoutes Register routes
func (c *UserController) RegisterRoutes(routes *http.ServeMux) *http.ServeMux {
	err := utils.Environment()
	if err != nil {
		log.Println(err)
		return routes
	}

	routes.HandleFunc(os.Getenv("DEFAULT_API_LINK")+"/register", c.Register)
	routes.HandleFunc(os.Getenv("DEFAULT_API_LINK")+"/login", c.Login)
	routes.HandleFunc(os.Getenv("DEFAULT_API_LINK")+"/profile/{id}", c.GetProfile)
	routes.HandleFunc(os.Getenv("DEFAULT_API_LINK")+"/profile/update/{id}", c.UpdateProfile)
	routes.HandleFunc(os.Getenv("DEFAULT_API_LINK")+"/follow", c.Follow)
	routes.HandleFunc(os.Getenv("DEFAULT_API_LINK")+"/unfollow", c.Unfollow)
	routes.HandleFunc(os.Getenv("DEFAULT_API_LINK")+"/followers/{id}", c.GetFollowers)

	return routes
}
