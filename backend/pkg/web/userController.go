package web

import (
	"backend/pkg/dto"
	"backend/pkg/service/impl"
	"backend/pkg/session"
	"backend/pkg/utils"
	"encoding/json"
	"net/http"
	"os"
	"strings"
	"time"
)

type UserController struct {
	UserService impl.UserServiceImpl
}

// Register Create new user controller
func (c *UserController) Register(w http.ResponseWriter, r *http.Request) {
	err := utils.Environment()
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusInternalServerError, "-", err.Error()+utils.Reset)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method != http.MethodPost {
		utils.LoggerError.Println(utils.Error, http.StatusMethodNotAllowed, "-", os.Getenv("METHOD_NOT_ALLOWED")+utils.Reset)
		http.Error(w, os.Getenv("METHOD_NOT_ALLOWED"), http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != os.Getenv("DEFAULT_API_LINK")+"/register" {
		utils.LoggerError.Println(utils.Error, http.StatusNotFound, "-", os.Getenv("NOT_FOUND")+utils.Reset)
		http.Error(w, os.Getenv("NOT_FOUND"), http.StatusNotFound)
		return
	}

	var userDTO dto.UserDTO
	if err := json.NewDecoder(r.Body).Decode(&userDTO); err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusBadRequest, "-", err.Error(), utils.Reset)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.UserService.CreateUser(&userDTO)
	if err != nil {
		utils.LoggerInfo.Println(http.StatusInternalServerError, "-", err.Error())
		err = json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	w.WriteHeader(http.StatusCreated)
	utils.LoggerInfo.Println(utils.Info, http.StatusCreated, "-", "User created successfully"+utils.Reset)
}

func (c *UserController) Login(w http.ResponseWriter, r *http.Request) {
	err := utils.Environment()
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusInternalServerError, "-", err.Error(), utils.Reset)
		err = json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	if r.Method != http.MethodPost {
		utils.LoggerError.Println(utils.Error, http.StatusMethodNotAllowed, "-", os.Getenv("METHOD_NOT_ALLOWED")+utils.Reset)
		http.Error(w, os.Getenv("METHOD_NOT_ALLOWED"), http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != os.Getenv("DEFAULT_API_LINK")+"/login" {
		utils.LoggerError.Println(utils.Error, http.StatusNotFound, "-", os.Getenv("NOT_FOUND")+utils.Reset)
		http.Error(w, os.Getenv("NOT_FOUND"), http.StatusNotFound)
		return
	}

	var credentials dto.UserConnectionDTO

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusBadRequest, "-", err.Error(), utils.Reset)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userDTO, err := c.UserService.Connection(credentials.Email, credentials.Password)
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusUnauthorized, "-", err.Error(), utils.Reset)
		err = json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  http.StatusUnauthorized,
			"message": "Email or password is incorrect",
		})
		return
	}

	// I will add the token generation here
	sessionToken, err := c.UserService.CreateSession(userDTO)
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusUnauthorized, "-", err.Error(), utils.Reset)
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	session.SetSessionCookie(w, sessionToken)
	w.WriteHeader(http.StatusOK)
	utils.LoggerInfo.Println(http.StatusOK, "-", "User logged in successfully")
	w.Header().Set(os.Getenv("CONTENT_TYPE"), os.Getenv("APPLICATION_JSON"))
	err = json.NewEncoder(w).Encode(map[string]interface{}{
		"status": http.StatusOK,
		"token":  sessionToken,
		"user":   userDTO,
	})
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusInternalServerError, "-", err.Error(), utils.Reset)
		err = json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
}

func (c *UserController) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	err := utils.Environment()
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusInternalServerError, "-", err.Error(), utils.Reset)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method != http.MethodPut {
		utils.LoggerError.Println(utils.Error, http.StatusMethodNotAllowed, "-", os.Getenv("METHOD_NOT_ALLOWED")+utils.Reset)
		http.Error(w, os.Getenv("METHOD_NOT_ALLOWED"), http.StatusMethodNotAllowed)
		return
	}

	if !strings.HasPrefix(r.URL.Path, os.Getenv("DEFAULT_API_LINK")+"/profile-update/") {
		utils.LoggerError.Println(utils.Error, http.StatusNotFound, "-", os.Getenv("NOT_FOUND")+utils.Reset)
		http.Error(w, os.Getenv("NOT_FOUND"), http.StatusNotFound)
		return
	}

	id, err := utils.ExtractIDFromRequest(r)
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusBadRequest, "- User ID is required", err.Error(), utils.Reset)
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	limit, err := c.UserService.CountUsers()
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusInternalServerError, "-", err.Error(), utils.Reset)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if id > limit {
		utils.LoggerError.Println(utils.Error, http.StatusNotFound, "-", "User not exists", utils.Reset)
		http.Error(w, "Invalid Request", http.StatusNotFound)
		return
	}

	var userDTO dto.UserDTO
	if err := json.NewDecoder(r.Body).Decode(&userDTO); err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusBadRequest, "-", err.Error(), utils.Reset)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.UserService.UpdateProfile(id, &userDTO)
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusInternalServerError, "-", err.Error(), utils.Reset)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (c *UserController) GetProfile(w http.ResponseWriter, r *http.Request) {
	err := utils.Environment()
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusInternalServerError, "-", err.Error(), utils.Reset)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method != http.MethodGet {
		utils.LoggerError.Println(utils.Error, http.StatusMethodNotAllowed, "-", os.Getenv("METHOD_NOT_ALLOWED")+utils.Reset)
		http.Error(w, os.Getenv("METHOD_NOT_ALLOWED"), http.StatusMethodNotAllowed)
		return
	}

	if !strings.HasPrefix(r.URL.Path, os.Getenv("DEFAULT_API_LINK")+"/profile/") {
		utils.LoggerError.Println(utils.Error, http.StatusNotFound, "-", os.Getenv("NOT_FOUND")+utils.Reset)
		http.Error(w, os.Getenv("NOT_FOUND"), http.StatusNotFound)
		return
	}

	id, err := utils.ExtractIDFromRequest(r)
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusBadRequest, "- User ID is required", err.Error(), utils.Reset)
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	limit, err := c.UserService.CountUsers()
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusInternalServerError, "-", err.Error(), utils.Reset)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if id > limit {
		utils.LoggerError.Println(utils.Error, http.StatusNotFound, "-", "User not exists", utils.Reset)
		http.Error(w, "Invalid Request", http.StatusNotFound)
		return
	}

	userDTO, err := c.UserService.GetProfile(id)
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusInternalServerError, "-", err.Error(), utils.Reset)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	utils.LoggerInfo.Println(utils.Info, http.StatusOK, "-", "User profile retrieved successfully"+utils.Reset)
	w.Header().Set(os.Getenv("CONTENT_TYPE"), os.Getenv("APPLICATION_JSON"))
	//err = json.NewEncoder(w).Encode(userDTO)
	if userDTO != nil {
		err = json.NewEncoder(w).Encode(map[string]interface{}{
			"status": http.StatusOK,
			"user":   userDTO,
		})
	} else {
		err = json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  http.StatusNoContent,
			"message": "No user found",
		})
	}
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusInternalServerError, "-", err.Error(), utils.Reset)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c *UserController) IsUserOnline(w http.ResponseWriter, r *http.Request) {
	err := utils.Environment()
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusInternalServerError, "-", err.Error(), utils.Reset)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method != http.MethodPost {
		utils.LoggerError.Println(utils.Error, http.StatusMethodNotAllowed, "-", os.Getenv("METHOD_NOT_ALLOWED")+utils.Reset)
		http.Error(w, os.Getenv("METHOD_NOT_ALLOWED"), http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != os.Getenv("DEFAULT_API_LINK")+"/is_online" {
		utils.LoggerError.Println(utils.Error, http.StatusNotFound, "-", os.Getenv("NOT_FOUND")+utils.Reset)
		http.Error(w, os.Getenv("NOT_FOUND"), http.StatusNotFound)
		return
	}

	var data struct {
		Token string `json:"token"`
	}

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusBadRequest, "-", err.Error(), utils.Reset)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	isOnline, err := c.UserService.IsUserOnline(data.Token)
	if err != nil {
		utils.LoggerInfo.Println(utils.Warn, http.StatusUnauthorized, "-", err.Error(), utils.Reset)
		err = json.NewEncoder(w).Encode(map[string]interface{}{
			"status":    http.StatusUnauthorized,
			"is_online": isOnline,
			"message":   err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	utils.LoggerInfo.Println(utils.Info, http.StatusOK, "-", "User is online"+utils.Reset)
	w.Header().Set(os.Getenv("CONTENT_TYPE"), os.Getenv("APPLICATION_JSON"))
	err = json.NewEncoder(w).Encode(map[string]interface{}{
		"status":    http.StatusOK,
		"is_online": isOnline,
	})
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusInternalServerError, "-", err.Error(), utils.Reset)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c *UserController) Logout(w http.ResponseWriter, r *http.Request) {
	err := utils.Environment()
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusInternalServerError, "-", err.Error(), utils.Reset)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method != http.MethodPost {
		utils.LoggerError.Println(utils.Error, http.StatusMethodNotAllowed, "-", os.Getenv("METHOD_NOT_ALLOWED")+utils.Reset)
		http.Error(w, os.Getenv("METHOD_NOT_ALLOWED"), http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != os.Getenv("DEFAULT_API_LINK")+"/logout" {
		utils.LoggerError.Println(utils.Error, http.StatusNotFound, "-", os.Getenv("NOT_FOUND")+utils.Reset)
		http.Error(w, os.Getenv("NOT_FOUND"), http.StatusNotFound)
		return
	}

	var data struct {
		Token string `json:"token"`
	}

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusBadRequest, "-", err.Error(), utils.Reset)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.UserService.Logout(data.Token)
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusUnauthorized, "-", err.Error(), utils.Reset)
		http.Error(w, err.Error(), http.StatusUnauthorized)
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
	utils.LoggerInfo.Println(utils.Info, http.StatusOK, "-", "Logged out successfully"+utils.Reset)
	w.Header().Set(os.Getenv("CONTENT_TYPE"), os.Getenv("APPLICATION_JSON"))
	err = json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  http.StatusOK,
		"message": "Logged out successfully",
	})
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusInternalServerError, "-", err.Error(), utils.Reset)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c *UserController) Users(w http.ResponseWriter, r *http.Request) {
	err := utils.Environment()
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusInternalServerError, "-", err.Error(), utils.Reset)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method != http.MethodGet {
		utils.LoggerError.Println(utils.Error, http.StatusMethodNotAllowed, "-", os.Getenv("METHOD_NOT_ALLOWED")+utils.Reset)
		http.Error(w, os.Getenv("METHOD_NOT_ALLOWED"), http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != os.Getenv("DEFAULT_API_LINK")+"/users" {
		utils.LoggerError.Println(utils.Error, http.StatusNotFound, "-", os.Getenv("NOT_FOUND")+utils.Reset)
		http.Error(w, os.Getenv("NOT_FOUND"), http.StatusNotFound)
		return
	}

	users, err := c.UserService.GetAllUsers()
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusInternalServerError, "-", err.Error()+utils.Reset)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	utils.LoggerInfo.Println(utils.Info, http.StatusOK, "-", "Users retrieved successfully"+utils.Reset)
	w.Header().Set(os.Getenv("CONTENT_TYPE"), os.Getenv("APPLICATION_JSON"))
	// err = json.NewEncoder(w).Encode(users)
	if len(users) != 0 {
		err = json.NewEncoder(w).Encode(map[string]interface{}{
			"status": http.StatusOK,
			"users":  users,
		})
	} else {
		err = json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  http.StatusNoContent,
			"message": "No users found",
		})
	}
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusInternalServerError, "-", err.Error()+utils.Reset)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c *UserController) GetFollowers(w http.ResponseWriter, r *http.Request) {
	err := utils.Environment()
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusInternalServerError, "-", err.Error(), utils.Reset)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method != http.MethodGet {
		utils.LoggerError.Println(utils.Error, http.StatusMethodNotAllowed, "-", os.Getenv("METHOD_NOT_ALLOWED")+utils.Reset)
		http.Error(w, os.Getenv("METHOD_NOT_ALLOWED"), http.StatusMethodNotAllowed)
		return
	}

	if !strings.HasPrefix(r.URL.Path, os.Getenv("DEFAULT_API_LINK")+"/followers/") {
		utils.LoggerError.Println(utils.Error, http.StatusNotFound, "-", os.Getenv("NOT_FOUND")+utils.Reset)
		http.Error(w, os.Getenv("NOT_FOUND"), http.StatusNotFound)
		return
	}

	userID, err := utils.ExtractIDFromRequest(r)
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusBadRequest, "-", os.Getenv("USER_ID_REQUIRED"), err.Error(), utils.Reset)
		http.Error(w, os.Getenv("USER_ID_REQUIRED"), http.StatusBadRequest)
		return
	}

	followers, err := c.UserService.GetFollowers(userID)
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusInternalServerError, "-", err.Error(), utils.Reset)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set(os.Getenv("CONTENT_TYPE"), os.Getenv("APPLICATION_JSON"))
	w.WriteHeader(http.StatusOK)
	utils.LoggerInfo.Println(utils.Info, http.StatusOK, "-", "Followers get"+utils.Reset)
	if followers != nil {
		err = json.NewEncoder(w).Encode(map[string]interface{}{
			"status":    http.StatusOK,
			"followers": followers,
		})
	} else {
		err = json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  http.StatusNoContent,
			"message": "No followers",
		})
	}
	if err != nil {
		utils.LoggerError.Println(utils.Error, "Followers do not get"+utils.Reset)
		return
	}
}

func (c *UserController) GetFollowings(w http.ResponseWriter, r *http.Request) {
	err := utils.Environment()
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusInternalServerError, "-", err.Error(), utils.Reset)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method != http.MethodGet {
		utils.LoggerError.Println(utils.Error, http.StatusMethodNotAllowed, "-", os.Getenv("METHOD_NOT_ALLOWED")+utils.Reset)
		http.Error(w, os.Getenv("METHOD_NOT_ALLOWED"), http.StatusMethodNotAllowed)
		return
	}

	if !strings.HasPrefix(r.URL.Path, os.Getenv("DEFAULT_API_LINK")+"/followings/") {
		utils.LoggerError.Println(utils.Error, http.StatusNotFound, "-", os.Getenv("NOT_FOUND")+utils.Reset)
		http.Error(w, os.Getenv("NOT_FOUND"), http.StatusNotFound)
		return
	}

	userID, err := utils.ExtractIDFromRequest(r)
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusBadRequest, "-", os.Getenv("USER_ID_REQUIRED"), err.Error(), utils.Reset)
		http.Error(w, os.Getenv("USER_ID_REQUIRED"), http.StatusBadRequest)
		return
	}

	followings, err := c.UserService.GetFollowings(userID)
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusInternalServerError, "-", err.Error(), utils.Reset)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set(os.Getenv("CONTENT_TYPE"), os.Getenv("APPLICATION_JSON"))
	w.WriteHeader(http.StatusOK)
	utils.LoggerInfo.Println(utils.Info, http.StatusOK, "-", "Followings get"+utils.Reset)
	if followings != nil {
		err = json.NewEncoder(w).Encode(map[string]interface{}{
			"status":     http.StatusOK,
			"followings": followings,
		})
	} else {
		err = json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  http.StatusNoContent,
			"message": "No followings",
		})
	}
	if err != nil {
		utils.LoggerError.Println(utils.Error, "Followings do not get"+utils.Reset)
		return
	}
}

func (c *UserController) GetFriends(w http.ResponseWriter, r *http.Request) {
	err := utils.Environment()
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusInternalServerError, "-", err.Error()+utils.Reset)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method != http.MethodGet {
		utils.LoggerError.Println(utils.Error, http.StatusMethodNotAllowed, "-", os.Getenv("METHOD_NOT_ALLOWED")+utils.Reset)
		http.Error(w, os.Getenv("METHOD_NOT_ALLOWED"), http.StatusMethodNotAllowed)
		return
	}

	if !strings.HasPrefix(r.URL.Path, os.Getenv("DEFAULT_API_LINK")+"/friends/") {
		utils.LoggerError.Println(utils.Error, http.StatusNotFound, "-", os.Getenv("NOT_FOUND")+utils.Reset)
		http.Error(w, os.Getenv("NOT_FOUND"), http.StatusNotFound)
		return
	}

	userID, err := utils.ExtractIDFromRequest(r)
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusBadRequest, "-", os.Getenv("USER_ID_REQUIRED"), err.Error(), utils.Reset)
		http.Error(w, os.Getenv("USER_ID_REQUIRED"), http.StatusBadRequest)
		return
	}

	friends, err := c.UserService.GetFriends(userID)
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusInternalServerError, "-", err.Error()+utils.Reset)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set(os.Getenv("CONTENT_TYPE"), os.Getenv("APPLICATION_JSON"))
	w.WriteHeader(http.StatusOK)
	utils.LoggerInfo.Println(utils.Info, http.StatusOK, "-", "Friends get"+utils.Reset)
	if friends != nil {
		err = json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  http.StatusOK,
			"friends": friends,
		})
	} else {
		err = json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  http.StatusNoContent,
			"message": "No friends",
		})
	}
	if err != nil {
		utils.LoggerError.Println(utils.Error, "Friends do not get"+utils.Reset)
		return
	}
}

func (c *UserController) GetFollowerCount(w http.ResponseWriter, r *http.Request) {
	err := utils.Environment()
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusInternalServerError, "-", err.Error(), utils.Reset)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method != http.MethodGet {
		utils.LoggerError.Println(utils.Error, http.StatusMethodNotAllowed, "-", os.Getenv("METHOD_NOT_ALLOWED")+utils.Reset)
		http.Error(w, os.Getenv("METHOD_NOT_ALLOWED"), http.StatusMethodNotAllowed)
		return
	}

	if !strings.HasPrefix(r.URL.Path, os.Getenv("DEFAULT_API_LINK")+"/follower-count/") {
		utils.LoggerError.Println(utils.Error, http.StatusNotFound, "-", os.Getenv("NOT_FOUND")+utils.Reset)
		http.Error(w, os.Getenv("NOT_FOUND"), http.StatusNotFound)
		return
	}

	userID, err := utils.ExtractIDFromRequest(r)
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusBadRequest, "- User ID is required to get follower's number", err.Error(), utils.Reset)
		http.Error(w, os.Getenv("USER_ID_REQUIRED"), http.StatusBadRequest)
		return
	}

	count, err := c.UserService.GetFollowerCount(userID)
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusInternalServerError, "-", err.Error(), utils.Reset)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set(os.Getenv("CONTENT_TYPE"), os.Getenv("APPLICATION_JSON"))
	w.WriteHeader(http.StatusOK)
	utils.LoggerInfo.Println(utils.Info, http.StatusOK, "-", "Follower count get"+utils.Reset)
	err = json.NewEncoder(w).Encode(map[string]interface{}{
		"status": http.StatusOK,
		"count":  count,
	})
	if err != nil {
		utils.LoggerError.Println(utils.Error, "Follower count do not get"+utils.Reset)
		return
	}
}

func (c *UserController) GetFollowingCount(w http.ResponseWriter, r *http.Request) {
	err := utils.Environment()
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusInternalServerError, "-", err.Error(), utils.Reset)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method != http.MethodGet {
		utils.LoggerError.Println(utils.Error, http.StatusMethodNotAllowed, "-", os.Getenv("METHOD_NOT_ALLOWED")+utils.Reset)
		http.Error(w, os.Getenv("METHOD_NOT_ALLOWED"), http.StatusMethodNotAllowed)
		return
	}

	if !strings.HasPrefix(r.URL.Path, os.Getenv("DEFAULT_API_LINK")+"/following-count/") {
		utils.LoggerError.Println(utils.Error, http.StatusNotFound, "-", os.Getenv("NOT_FOUND")+utils.Reset)
		http.Error(w, os.Getenv("NOT_FOUND"), http.StatusNotFound)
		return
	}

	userID, err := utils.ExtractIDFromRequest(r)
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusBadRequest, "- User ID is required", err.Error(), utils.Reset)
		http.Error(w, os.Getenv("USER_ID_REQUIRED"), http.StatusBadRequest)
		return
	}

	count, err := c.UserService.GetFollowingCount(userID)
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusInternalServerError, "-", err.Error(), utils.Reset)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set(os.Getenv("CONTENT_TYPE"), os.Getenv("APPLICATION_JSON"))
	w.WriteHeader(http.StatusOK)
	utils.LoggerInfo.Println(utils.Info, http.StatusOK, "-", "Following count get"+utils.Reset)
	err = json.NewEncoder(w).Encode(map[string]interface{}{
		"status": http.StatusOK,
		"count":  count,
	})
	if err != nil {
		utils.LoggerError.Println(utils.Error, "Following count do not get"+utils.Reset)
		return
	}
}

func (c *UserController) GetFriendCount(w http.ResponseWriter, r *http.Request) {
	err := utils.Environment()
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusInternalServerError, "-", err.Error(), utils.Reset)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method != http.MethodGet {
		utils.LoggerError.Println(utils.Error, http.StatusMethodNotAllowed, "-", os.Getenv("METHOD_NOT_ALLOWED")+utils.Reset)
		http.Error(w, os.Getenv("METHOD_NOT_ALLOWED"), http.StatusMethodNotAllowed)
		return
	}

	if !strings.HasPrefix(r.URL.Path, os.Getenv("DEFAULT_API_LINK")+"/friend-count/") {
		utils.LoggerError.Println(utils.Error, http.StatusNotFound, "-", os.Getenv("NOT_FOUND")+utils.Reset)
		http.Error(w, os.Getenv("NOT_FOUND"), http.StatusNotFound)
		return
	}

	userID, err := utils.ExtractIDFromRequest(r)
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusBadRequest, "- User ID is required", err.Error(), utils.Reset)
		http.Error(w, os.Getenv("USER_ID_REQUIRED"), http.StatusBadRequest)
		return
	}

	count, err := c.UserService.GetFriendsCount(userID)
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusInternalServerError, "-", err.Error(), utils.Reset)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set(os.Getenv("CONTENT_TYPE"), os.Getenv("APPLICATION_JSON"))
	w.WriteHeader(http.StatusOK)
	utils.LoggerInfo.Println(utils.Info, http.StatusOK, "-", "Friend count get"+utils.Reset)
	err = json.NewEncoder(w).Encode(map[string]interface{}{
		"status": http.StatusOK,
		"count":  count,
	})
	if err != nil {
		utils.LoggerError.Println(utils.Error, "Friend count do not get"+utils.Reset)
		return
	}
}

// UsersRoutes Register routes
func (c *UserController) UsersRoutes(routes *http.ServeMux) *http.ServeMux {
	err := utils.Environment()
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusInternalServerError, "-", err.Error()+utils.Reset)
		return routes
	}

	routes.HandleFunc(os.Getenv("DEFAULT_API_LINK")+"/register", c.Register)
	routes.HandleFunc(os.Getenv("DEFAULT_API_LINK")+"/login", c.Login)
	routes.HandleFunc(os.Getenv("DEFAULT_API_LINK")+"/users", c.Users)
	routes.HandleFunc(os.Getenv("DEFAULT_API_LINK")+"/logout", c.Logout)
	routes.HandleFunc(os.Getenv("DEFAULT_API_LINK")+"/is_online", c.IsUserOnline)
	routes.HandleFunc(os.Getenv("DEFAULT_API_LINK")+"/profile/", c.GetProfile)
	routes.HandleFunc(os.Getenv("DEFAULT_API_LINK")+"/profile-update/", c.UpdateProfile)
	routes.HandleFunc(os.Getenv("DEFAULT_API_LINK")+"/followers/", c.GetFollowers)
	routes.HandleFunc(os.Getenv("DEFAULT_API_LINK")+"/followings/", c.GetFollowings)
	routes.HandleFunc(os.Getenv("DEFAULT_API_LINK")+"/follower-count/", c.GetFollowerCount)
	routes.HandleFunc(os.Getenv("DEFAULT_API_LINK")+"/following-count/", c.GetFollowingCount)
	routes.HandleFunc(os.Getenv("DEFAULT_API_LINK")+"/friend-count/", c.GetFriendCount)
	routes.HandleFunc(os.Getenv("DEFAULT_API_LINK")+"/friends/", c.GetFriends)

	return routes
}
