package web

import (
	"backend/pkg/service/impl"
	"backend/pkg/utils"
	"encoding/json"
	"net/http"
	"os"
	"strings"
)

type FollowController struct {
	FollowService impl.FollowServiceImpl
}

func (c *FollowController) FollowUser(w http.ResponseWriter, r *http.Request) {
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

	if r.URL.Path != os.Getenv("DEFAULT_API_LINK")+"/follow" {
		utils.LoggerError.Println(utils.Error, http.StatusNotFound, "-", os.Getenv("NOT_FOUND")+utils.Reset)
		http.Error(w, os.Getenv("NOT_FOUND"), http.StatusNotFound)
		return
	}

	var follow struct {
		FollowerID uint `json:"follower_id" db:"follower_id"`
		FolloweeID uint `json:"followee_id" db:"followee_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&follow); err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusBadRequest, "-", err.Error(), utils.Reset)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if follow.FollowerID == follow.FolloweeID {
		utils.LoggerError.Println(utils.Error, http.StatusBadRequest, "-", "Cannot follow yourself"+utils.Reset)
		http.Error(w, "Cannot follow yourself", http.StatusBadRequest)
		return
	}

	err = c.FollowService.FollowUser(follow.FollowerID, follow.FolloweeID)
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusInternalServerError, "-", err.Error(), utils.Reset)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	utils.LoggerInfo.Println(utils.Info, http.StatusOK, "-", "Follow request sent"+utils.Reset)
	_, err = w.Write([]byte("Follow request sent"))
	if err != nil {
		utils.LoggerError.Println(utils.Error, "Follow request do not send"+utils.Reset)
		return
	}
}

func (c *FollowController) UnfollowUser(w http.ResponseWriter, r *http.Request) {
	err := utils.Environment()
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusInternalServerError, "-", err.Error(), utils.Reset)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method != http.MethodDelete {
		utils.LoggerError.Println(utils.Error, http.StatusMethodNotAllowed, "-", os.Getenv("METHOD_NOT_ALLOWED")+utils.Reset)
		http.Error(w, os.Getenv("METHOD_NOT_ALLOWED"), http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != os.Getenv("DEFAULT_API_LINK")+"/unfollow" {
		utils.LoggerError.Println(utils.Error, http.StatusNotFound, "-", os.Getenv("NOT_FOUND")+utils.Reset)
		http.Error(w, os.Getenv("NOT_FOUND"), http.StatusNotFound)
		return
	}

	var follow struct {
		FollowerID uint `json:"follower_id" db:"follower_id"`
		FolloweeID uint `json:"followee_id" db:"followee_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&follow); err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusBadRequest, "-", err.Error(), utils.Reset)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if follow.FollowerID == follow.FolloweeID {
		utils.LoggerInfo.Println(utils.Info, http.StatusBadRequest, "-", "Cannot unfollow yourself"+utils.Reset)
		http.Error(w, "Cannot unfollow yourself", http.StatusBadRequest)
		return
	}

	err = c.FollowService.UnfollowUser(follow.FollowerID, follow.FolloweeID)
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusInternalServerError, "-", err.Error(), utils.Reset)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	utils.LoggerInfo.Println(http.StatusOK, "-", "Unfollow request sent"+utils.Reset)
	_, err = w.Write([]byte("Unfollow request sent"))
	if err != nil {
		utils.LoggerError.Println(utils.Error, "Unfollow request do not send"+utils.Reset)
		return
	}
}

func (c *FollowController) AcceptFollowRequest(w http.ResponseWriter, r *http.Request) {
	err := utils.Environment()
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusInternalServerError, "-", err.Error()+utils.Reset)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method != http.MethodPut {
		utils.LoggerError.Println(utils.Error, http.StatusMethodNotAllowed, "-", os.Getenv("METHOD_NOT_ALLOWED")+utils.Reset)
		http.Error(w, os.Getenv("METHOD_NOT_ALLOWED"), http.StatusMethodNotAllowed)
		return
	}

	if !strings.HasPrefix(r.URL.Path, os.Getenv("DEFAULT_API_LINK")+"/accept/") {
		utils.LoggerError.Println(utils.Error, http.StatusNotFound, "-", os.Getenv("NOT_FOUND")+utils.Reset)
		http.Error(w, os.Getenv("NOT_FOUND"), http.StatusNotFound)
		return
	}

	id, err := utils.ExtractIDFromRequest(r)
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusBadRequest, "-", "Follow ID is required"+utils.Reset)
		http.Error(w, "Follow ID is required", http.StatusBadRequest)
		return
	}

	limit, err := c.FollowService.CountAllFollows()
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusInternalServerError, "-", err.Error()+utils.Reset)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if id > limit {
		utils.LoggerError.Println(utils.Error, http.StatusNotFound, "-", os.Getenv("NOT_FOUND")+utils.Reset)
		http.Error(w, os.Getenv("NOT_FOUND"), http.StatusNotFound)
		return
	}

	/*if r.URL.Path != os.Getenv("DEFAULT_API_LINK")+"/accept" {
		utils.LoggerError.Println(utils.Error, http.StatusNotFound, "-", os.Getenv("NOT_FOUND"))
		http.Error(w, os.Getenv("NOT_FOUND"), http.StatusNotFound)
		return
	}

	var follow struct {
		ID uint `json:"id" db:"id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&follow); err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusBadRequest, "-", err.Error(), utils.Reset)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}*/

	//err = c.FollowService.AcceptFollowRequest(follow.ID)
	err = c.FollowService.AcceptFollowRequest(id)
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusInternalServerError, "-", err.Error(), utils.Reset)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	utils.LoggerInfo.Println(http.StatusOK, "-", "Follow request accepted"+utils.Reset)
	_, err = w.Write([]byte("Follow request accepted"))
	if err != nil {
		utils.LoggerError.Println(utils.Error, "Follow request do not accepted"+utils.Reset)
		return
	}
}

func (c *FollowController) DeclineFollowRequest(w http.ResponseWriter, r *http.Request) {
	err := utils.Environment()
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusInternalServerError, "-", err.Error(), utils.Reset)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method != http.MethodDelete {
		utils.LoggerError.Println(utils.Error, http.StatusMethodNotAllowed, "-", os.Getenv("METHOD_NOT_ALLOWED")+utils.Reset)
		http.Error(w, os.Getenv("METHOD_NOT_ALLOWED"), http.StatusMethodNotAllowed)
		return
	}

	if !strings.HasPrefix(r.URL.Path, os.Getenv("DEFAULT_API_LINK")+"/decline/") {
		utils.LoggerError.Println(utils.Error, http.StatusNotFound, "-", os.Getenv("NOT_FOUND")+utils.Reset)
		http.Error(w, os.Getenv("NOT_FOUND"), http.StatusNotFound)
		return
	}

	id, err := utils.ExtractIDFromRequest(r)
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusBadRequest, "-", "Follow request ID is required"+utils.Reset)
		http.Error(w, "Follow request ID is required", http.StatusBadRequest)
		return
	}

	limit, err := c.FollowService.CountAllFollows()
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusInternalServerError, "-", err.Error()+utils.Reset)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if id > limit {
		utils.LoggerError.Println(utils.Error, http.StatusNotFound, "-", os.Getenv("NOT_FOUND")+utils.Reset)
		http.Error(w, os.Getenv("NOT_FOUND"), http.StatusNotFound)
		return
	}

	/*if r.URL.Path != os.Getenv("DEFAULT_API_LINK")+"/decline" {
		utils.LoggerError.Println(utils.Error, http.StatusNotFound, "-", os.Getenv("NOT_FOUND")+utils.Reset)
		http.Error(w, os.Getenv("NOT_FOUND"), http.StatusNotFound)
		return
	}

	var follow struct {
		ID uint `json:"id" db:"id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&follow); err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusBadRequest, "-", err.Error(), utils.Reset)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}*/

	// err = c.FollowService.DeclineFollowRequest(follow.ID)
	err = c.FollowService.DeclineFollowRequest(id)
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusInternalServerError, "-", err.Error(), utils.Reset)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	utils.LoggerInfo.Println(http.StatusOK, "-", "Follow request declined"+utils.Reset)
	_, err = w.Write([]byte("Follow request declined"))
	if err != nil {
		utils.LoggerError.Println(utils.Error, "Follow request do not declined"+utils.Reset)
		return
	}
}

func (c *FollowController) GetPendingFollowRequest(w http.ResponseWriter, r *http.Request) {
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

	if !strings.HasPrefix(r.URL.Path, os.Getenv("DEFAULT_API_LINK")+"/pending/") {
		utils.LoggerError.Println(utils.Error, http.StatusNotFound, "-", os.Getenv("NOT_FOUND")+utils.Reset)
		http.Error(w, os.Getenv("NOT_FOUND"), http.StatusNotFound)
		return
	}

	userID, err := utils.ExtractIDFromRequest(r)
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusBadRequest, "- User ID is required", err.Error(), utils.Reset)
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	follows, err := c.FollowService.GetPendingFollowRequest(userID)
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusInternalServerError, "-", err.Error(), utils.Reset)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	utils.LoggerInfo.Println(utils.Info, http.StatusOK, "-", "Follow requests get"+utils.Reset)
	if follows != nil {
		err = json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  http.StatusOK,
			"follows": follows,
		})
	} else {
		err = json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  http.StatusNoContent,
			"message": "No follow requests",
		})
	}
	// err = json.NewEncoder(w).Encode(follows)
	/*err = json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  http.StatusOK,
		"follows": follows,
	})*/
	if err != nil {
		utils.LoggerError.Println(utils.Error, "Follow requests do not get"+utils.Reset)
		return
	}
}

func (c *FollowController) GetFollowers(w http.ResponseWriter, r *http.Request) {
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
		utils.LoggerError.Println(utils.Error, http.StatusBadRequest, "- User ID is required", err.Error(), utils.Reset)
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	followers, err := c.FollowService.GetFollowers(userID)
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusInternalServerError, "-", err.Error(), utils.Reset)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
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
	/*err = json.NewEncoder(w).Encode(map[string]interface{}{
		"status":    http.StatusOK,
		"followers": followers,
	})*/
	if err != nil {
		utils.LoggerError.Println(utils.Error, "Followers do not get"+utils.Reset)
		return
	}
}

func (c *FollowController) GetFollowings(w http.ResponseWriter, r *http.Request) {
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
		utils.LoggerError.Println(utils.Error, http.StatusBadRequest, "- User ID is required", err.Error(), utils.Reset)
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	followings, err := c.FollowService.GetFollowings(userID)
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusInternalServerError, "-", err.Error(), utils.Reset)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
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
	/*err = json.NewEncoder(w).Encode(map[string]interface{}{
		"status":     http.StatusOK,
		"followings": followings,
	})*/
	if err != nil {
		utils.LoggerError.Println(utils.Error, "Followings do not get"+utils.Reset)
		return
	}
}

func (c *FollowController) GetFollowerCount(w http.ResponseWriter, r *http.Request) {
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
		utils.LoggerError.Println(utils.Error, http.StatusBadRequest, "- User ID is required", err.Error(), utils.Reset)
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	count, err := c.FollowService.GetFollowerCount(userID)
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusInternalServerError, "-", err.Error(), utils.Reset)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
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

func (c *FollowController) GetFollowingCount(w http.ResponseWriter, r *http.Request) {
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
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	count, err := c.FollowService.GetFollowingCount(userID)
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusInternalServerError, "-", err.Error(), utils.Reset)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
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

func (c *FollowController) CountAllFollows(w http.ResponseWriter, r *http.Request) {
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

	if r.URL.Path != os.Getenv("DEFAULT_API_LINK")+"/follow-count" {
		utils.LoggerError.Println(utils.Error, http.StatusNotFound, "-", os.Getenv("NOT_FOUND")+utils.Reset)
		http.Error(w, os.Getenv("NOT_FOUND"), http.StatusNotFound)
		return
	}

	count, err := c.FollowService.CountAllFollows()
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusInternalServerError, "-", err.Error(), utils.Reset)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	utils.LoggerInfo.Println(utils.Info, http.StatusOK, "-", "Follow count get"+utils.Reset)
	err = json.NewEncoder(w).Encode(map[string]interface{}{
		"status": http.StatusOK,
		"count":  count,
	})
	if err != nil {
		utils.LoggerError.Println(utils.Error, "Follow count do not get"+utils.Reset)
		return
	}
}

func (c *FollowController) FollowsRoutes(routes *http.ServeMux) *http.ServeMux {
	err := utils.Environment()
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusInternalServerError, "-", err.Error(), utils.Reset)
		return routes
	}

	routes.HandleFunc(os.Getenv("DEFAULT_API_LINK")+"/follow", c.FollowUser)
	routes.HandleFunc(os.Getenv("DEFAULT_API_LINK")+"/unfollow", c.UnfollowUser)
	routes.HandleFunc(os.Getenv("DEFAULT_API_LINK")+"/accept/", c.AcceptFollowRequest)
	routes.HandleFunc(os.Getenv("DEFAULT_API_LINK")+"/decline/", c.DeclineFollowRequest)
	routes.HandleFunc(os.Getenv("DEFAULT_API_LINK")+"/pending/", c.GetPendingFollowRequest)
	routes.HandleFunc(os.Getenv("DEFAULT_API_LINK")+"/followers/", c.GetFollowers)
	routes.HandleFunc(os.Getenv("DEFAULT_API_LINK")+"/followings/", c.GetFollowings)
	routes.HandleFunc(os.Getenv("DEFAULT_API_LINK")+"/follower-count/", c.GetFollowerCount)
	routes.HandleFunc(os.Getenv("DEFAULT_API_LINK")+"/following-count/", c.GetFollowingCount)
	routes.HandleFunc(os.Getenv("DEFAULT_API_LINK")+"/follow-count", c.CountAllFollows)

	return routes
}
