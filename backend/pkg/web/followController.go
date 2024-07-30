package web

import (
	"backend/pkg/service/impl"
	"backend/pkg/utils"
	"log"
	"net/http"
	"os"
	"strconv"
)

type FollowController struct {
	FollowService impl.FollowServiceImpl
}

func (c *FollowController) FollowUser(w http.ResponseWriter, r *http.Request) {
	err := utils.Environment()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, os.Getenv("METHOD_NOT_ALLOWED"), http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != os.Getenv("DEFAULT_API_LINK")+"/follow" {
		http.Error(w, os.Getenv("NOT_FOUND"), http.StatusNotFound)
		return
	}

	/*
		var userDTO dto.UserDTO
			if err := json.NewDecoder(r.Body).Decode(&userDTO); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
	*/

	followerID, err := strconv.Atoi(r.FormValue("follower_id"))
	if err != nil {
		http.Error(w, "Invalid follower ID", http.StatusBadRequest)
		return
	}

	followeeID, err := strconv.Atoi(r.FormValue("followee_id"))
	if err != nil {
		http.Error(w, "Invalid followed ID", http.StatusBadRequest)
		return
	}

	if followerID == followeeID {
		http.Error(w, "Cannot follow yourself", http.StatusBadRequest)
		return
	}

	err = c.FollowService.FollowUser(uint(followerID), uint(followeeID))
	if err != nil {
		http.Error(w, "Error creating follow request", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("Follow request sent"))
	if err != nil {
		log.Println("Follow request sent")
		return
	}
}

func (c *FollowController) FollowsRoutes(routes *http.ServeMux) *http.ServeMux {
	err := utils.Environment()
	if err != nil {
		log.Println(err)
		return routes
	}

	routes.HandleFunc(os.Getenv("DEFAULT_API_LINK")+"/follow", c.FollowUser)

	return routes
}
