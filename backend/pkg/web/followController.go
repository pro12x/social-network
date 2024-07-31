package web

import (
	"backend/pkg/service/impl"
	"backend/pkg/utils"
	"encoding/json"
	"net/http"
	"os"
)

type FollowController struct {
	FollowService impl.FollowServiceImpl
}

func (c *FollowController) FollowUser(w http.ResponseWriter, r *http.Request) {
	err := utils.Environment()
	if err != nil {
		utils.Logger.Println(http.StatusInternalServerError, "-", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method != http.MethodPost {
		utils.Logger.Println(http.StatusMethodNotAllowed, "-", os.Getenv("METHOD_NOT_ALLOWED"))
		http.Error(w, os.Getenv("METHOD_NOT_ALLOWED"), http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != os.Getenv("DEFAULT_API_LINK")+"/follow" {
		utils.Logger.Println(http.StatusNotFound, "-", os.Getenv("NOT_FOUND"))
		http.Error(w, os.Getenv("NOT_FOUND"), http.StatusNotFound)
		return
	}

	var follow struct {
		FollowerID uint `json:"follower_id" db:"follower_id"`
		FolloweeID uint `json:"followee_id" db:"followee_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&follow); err != nil {
		utils.Logger.Println(http.StatusBadRequest, "-", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if follow.FollowerID == follow.FolloweeID {
		utils.Logger.Println(http.StatusBadRequest, "-", "Cannot follow yourself")
		http.Error(w, "Cannot follow yourself", http.StatusBadRequest)
		return
	}

	err = c.FollowService.FollowUser(follow.FollowerID, follow.FolloweeID)
	if err != nil {
		utils.Logger.Println(http.StatusInternalServerError, "-", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	utils.Logger.Println(http.StatusOK, "-", "Follow request sent")
	_, err = w.Write([]byte("Follow request sent"))
	if err != nil {
		utils.Logger.Println("Follow request do not send")
		return
	}
}

func (c *FollowController) UnfollowUser(w http.ResponseWriter, r *http.Request) {
	err := utils.Environment()
	if err != nil {
		utils.Logger.Println(http.StatusInternalServerError, "-", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method != http.MethodDelete {
		utils.Logger.Println(http.StatusMethodNotAllowed, "-", os.Getenv("METHOD_NOT_ALLOWED"))
		http.Error(w, os.Getenv("METHOD_NOT_ALLOWED"), http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != os.Getenv("DEFAULT_API_LINK")+"/unfollow" {
		utils.Logger.Println(http.StatusNotFound, "-", os.Getenv("NOT_FOUND"))
		http.Error(w, os.Getenv("NOT_FOUND"), http.StatusNotFound)
		return
	}

	var follow struct {
		FollowerID uint `json:"follower_id" db:"follower_id"`
		FolloweeID uint `json:"followee_id" db:"followee_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&follow); err != nil {
		utils.Logger.Println(http.StatusBadRequest, "-", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if follow.FollowerID == follow.FolloweeID {
		utils.Logger.Println(http.StatusBadRequest, "-", "Cannot unfollow yourself")
		http.Error(w, "Cannot unfollow yourself", http.StatusBadRequest)
		return
	}

	err = c.FollowService.UnfollowUser(follow.FollowerID, follow.FolloweeID)
	if err != nil {
		utils.Logger.Println(http.StatusInternalServerError, "-", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	utils.Logger.Println(http.StatusOK, "-", "Unfollow request sent")
	_, err = w.Write([]byte("Unfollow request sent"))
	if err != nil {
		utils.Logger.Println("Unfollow request do not send")
		return
	}
}

func (c *FollowController) AcceptFollowRequest(w http.ResponseWriter, r *http.Request) {
	err := utils.Environment()
	if err != nil {
		utils.Logger.Println(http.StatusInternalServerError, "-", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method != http.MethodPut {
		utils.Logger.Println(http.StatusMethodNotAllowed, "-", os.Getenv("METHOD_NOT_ALLOWED"))
		http.Error(w, os.Getenv("METHOD_NOT_ALLOWED"), http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != os.Getenv("DEFAULT_API_LINK")+"/accept" {
		utils.Logger.Println(http.StatusNotFound, "-", os.Getenv("NOT_FOUND"))
		http.Error(w, os.Getenv("NOT_FOUND"), http.StatusNotFound)
		return
	}

	var follow struct {
		ID uint `json:"id" db:"id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&follow); err != nil {
		utils.Logger.Println(http.StatusBadRequest, "-", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.FollowService.AcceptFollowRequest(follow.ID)
	if err != nil {
		utils.Logger.Println(http.StatusInternalServerError, "-", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	utils.Logger.Println(http.StatusOK, "-", "Follow request accepted")
	_, err = w.Write([]byte("Follow request accepted"))
	if err != nil {
		utils.Logger.Println("Follow request do not accepted")
		return
	}
}

func (c *FollowController) DeclineFollowRequest(w http.ResponseWriter, r *http.Request) {
	err := utils.Environment()
	if err != nil {
		utils.Logger.Println(http.StatusInternalServerError, "-", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method != http.MethodDelete {
		utils.Logger.Println(http.StatusMethodNotAllowed, "-", os.Getenv("METHOD_NOT_ALLOWED"))
		http.Error(w, os.Getenv("METHOD_NOT_ALLOWED"), http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != os.Getenv("DEFAULT_API_LINK")+"/decline" {
		utils.Logger.Println(http.StatusNotFound, "-", os.Getenv("NOT_FOUND"))
		http.Error(w, os.Getenv("NOT_FOUND"), http.StatusNotFound)
		return
	}

	var follow struct {
		ID uint `json:"id" db:"id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&follow); err != nil {
		utils.Logger.Println(http.StatusBadRequest, "-", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.FollowService.DeclineFollowRequest(follow.ID)
	if err != nil {
		utils.Logger.Println(http.StatusInternalServerError, "-", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	utils.Logger.Println(http.StatusOK, "-", "Follow request declined")
	_, err = w.Write([]byte("Follow request declined"))
	if err != nil {
		utils.Logger.Println("Follow request do not declined")
		return
	}
}

func (c *FollowController) FollowsRoutes(routes *http.ServeMux) *http.ServeMux {
	err := utils.Environment()
	if err != nil {
		utils.Logger.Println(http.StatusInternalServerError, "-", err.Error())
		return routes
	}

	routes.HandleFunc(os.Getenv("DEFAULT_API_LINK")+"/follow", c.FollowUser)
	routes.HandleFunc(os.Getenv("DEFAULT_API_LINK")+"/unfollow", c.UnfollowUser)
	routes.HandleFunc(os.Getenv("DEFAULT_API_LINK")+"/accept", c.AcceptFollowRequest)
	routes.HandleFunc(os.Getenv("DEFAULT_API_LINK")+"/decline", c.DeclineFollowRequest)

	return routes
}
