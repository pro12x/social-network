package web

import (
	"backend/pkg/dto"
	"backend/pkg/service/impl"
	"backend/pkg/utils"
	"encoding/json"
	"net/http"
	"os"
)

type UserController struct {
	userService impl.UserServiceImpl
}

// GetUserById Get user by id
func (c *UserController) GetUserById(w http.ResponseWriter, r *http.Request) {
	//
}

// Register Create new user controller
func (c *UserController) Register(w http.ResponseWriter, r *http.Request) {
	err := utils.Environment()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != os.Getenv("DEFAULT_API_LINK")+"/register" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	var userDTO dto.UserDTO
	if err := json.NewDecoder(r.Body).Decode(&userDTO); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = c.userService.CreateUser(&userDTO)
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
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != os.Getenv("DEFAULT_API_LINK")+"/login" {
		http.Error(w, "Not found", http.StatusNotFound)
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

	userDTO, err := c.userService.Connection(credentials.Email, credentials.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// I will add the token generation here

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(userDTO)
	err = json.NewEncoder(w).Encode(map[string]interface{}{
		"token": "",
		"user":  userDTO,
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
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != os.Getenv("DEFAULT_API_LINK")+"/profile/{id}" {
		http.Error(w, "Not found", http.StatusNotFound)
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

	err = c.userService.UpdateProfile(uint(id), &userDTO)
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
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != os.Getenv("DEFAULT_API_LINK")+"/profile/{id}" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	id, err := utils.ExtractIDFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userDTO, err := c.userService.GetProfile(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(userDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// RegisterRoutes Register routes
func (c *UserController) RegisterRoutes() {
	http.HandleFunc("/user/{id}", c.GetUserById)
	http.HandleFunc("/register", c.Register)
}
