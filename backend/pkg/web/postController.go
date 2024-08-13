package web

import (
	"backend/pkg/entity"
	"backend/pkg/globale"
	"backend/pkg/repository"
	"backend/pkg/service/impl"
	"backend/pkg/utils"
	"encoding/json"
	"net/http"
	"os"
	"strings"
)

type PostController struct {
	PostService impl.PostServiceImpl
}

func (p *PostController) CreatePost(w http.ResponseWriter, r *http.Request) {
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

	if r.URL.Path != os.Getenv("DEFAULT_API_LINK")+"/post" {
		utils.LoggerError.Println(utils.Error, http.StatusNotFound, "-", os.Getenv("NOT_FOUND")+utils.Reset)
		http.Error(w, os.Getenv("NOT_FOUND"), http.StatusNotFound)
		return
	}

	post := new(entity.Post)
	if err := json.NewDecoder(r.Body).Decode(post); err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusBadRequest, "-", err.Error(), utils.Reset)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !utils.CheckPost(post) {
		utils.LoggerError.Println(utils.Error, http.StatusBadRequest, "-", os.Getenv("BAD_REQUEST")+utils.Reset)
		http.Error(w, os.Getenv("BAD_REQUEST"), http.StatusBadRequest)
		return
	}

	err = p.PostService.CreatePost(post)
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusInternalServerError, "-", err.Error(), utils.Reset)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	utils.LoggerInfo.Println(utils.Info, http.StatusCreated, "-", "Post created"+utils.Reset)
	_, err = w.Write([]byte("Post created"))
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusInternalServerError, "-", err.Error(), utils.Reset)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (p *PostController) UpdatePost(w http.ResponseWriter, r *http.Request) {
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

	if !strings.HasPrefix(r.URL.Path, os.Getenv("DEFAULT_API_LINK")+"/update-post/") {
		utils.LoggerError.Println(utils.Error, http.StatusNotFound, "-", os.Getenv("NOT_FOUND")+utils.Reset)
		http.Error(w, os.Getenv("NOT_FOUND"), http.StatusNotFound)
		return
	}

	id, err := utils.ExtractIDFromRequest(r)
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusBadRequest, "-", err.Error(), utils.Reset)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	limit, err := p.PostService.CountAllPosts()
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusInternalServerError, "-", err.Error(), utils.Reset)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if id > limit {
		utils.LoggerError.Println(utils.Error, http.StatusNotFound, "-", os.Getenv("NOT_FOUND")+utils.Reset)
		http.Error(w, os.Getenv("NOT_FOUND"), http.StatusNotFound)
		return
	}

	post := new(entity.Post)
	if err := json.NewDecoder(r.Body).Decode(post); err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusBadRequest, "-", err.Error(), utils.Reset)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = p.PostService.UpdatePost(post, id)
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusInternalServerError, "-", err.Error(), utils.Reset)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	utils.LoggerInfo.Println(utils.Info, http.StatusOK, "-", "Post updated"+utils.Reset)
	_, err = w.Write([]byte("Post updated"))
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusInternalServerError, "-", err.Error(), utils.Reset)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (p *PostController) DeletePost(w http.ResponseWriter, r *http.Request) {
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

	if !strings.HasPrefix(r.URL.Path, os.Getenv("DEFAULT_API_LINK")+"/delete-post/") {
		utils.LoggerError.Println(utils.Error, http.StatusNotFound, "-", os.Getenv("NOT_FOUND")+utils.Reset)
		http.Error(w, os.Getenv("NOT_FOUND"), http.StatusNotFound)
		return
	}

	id, err := utils.ExtractIDFromRequest(r)
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusBadRequest, "-", err.Error(), utils.Reset)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	limit, err := p.PostService.CountAllPosts()
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusInternalServerError, "-", err.Error(), utils.Reset)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if id > limit {
		utils.LoggerError.Println(utils.Error, http.StatusNotFound, "-", os.Getenv("NOT_FOUND")+utils.Reset)
		http.Error(w, os.Getenv("NOT_FOUND"), http.StatusNotFound)
		return
	}

	err = p.PostService.DeletePost(id)
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusInternalServerError, "-", err.Error(), utils.Reset)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	utils.LoggerInfo.Println(utils.Info, http.StatusOK, "-", "Post deleted"+utils.Reset)
	_, err = w.Write([]byte("Post deleted"))
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusInternalServerError, "-", err.Error(), utils.Reset)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (p *PostController) GetPost(w http.ResponseWriter, r *http.Request) {
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

	if !strings.HasPrefix(r.URL.Path, os.Getenv("DEFAULT_API_LINK")+"/get-post/") {
		utils.LoggerError.Println(utils.Error, http.StatusNotFound, "-", os.Getenv("NOT_FOUND")+utils.Reset)
		http.Error(w, os.Getenv("NOT_FOUND"), http.StatusNotFound)
		return
	}

	id, err := utils.ExtractIDFromRequest(r)
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusBadRequest, "-", err.Error(), utils.Reset)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	limit, err := p.PostService.CountAllPosts()
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusInternalServerError, "-", err.Error(), utils.Reset)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if id > limit {
		utils.LoggerError.Println(utils.Error, http.StatusNotFound, "-", os.Getenv("NOT_FOUND")+utils.Reset)
		http.Error(w, os.Getenv("NOT_FOUND"), http.StatusNotFound)
		return
	}

	post, err := p.PostService.FindPostByID(id)
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusInternalServerError, "-", err.Error(), utils.Reset)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if post == nil {
		utils.LoggerError.Println(utils.Error, http.StatusNotFound, "-", os.Getenv("NOT_FOUND")+utils.Reset)
		http.Error(w, os.Getenv("NOT_FOUND"), http.StatusNotFound)
		return
	}

	w.Header().Set(os.Getenv("CONTENT_TYPE"), os.Getenv("APPLICATION_JSON"))
	w.WriteHeader(http.StatusOK)
	utils.LoggerInfo.Println(utils.Info, http.StatusOK, "-", "Post found"+utils.Reset)
	err = json.NewEncoder(w).Encode(map[string]interface{}{
		"status": http.StatusOK,
		"post":   post,
	})
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusInternalServerError, "-", err.Error(), utils.Reset)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (p *PostController) GetAllPosts(w http.ResponseWriter, r *http.Request) {
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

	if r.URL.Path != os.Getenv("DEFAULT_API_LINK")+"/posts" {
		utils.LoggerError.Println(utils.Error, http.StatusNotFound, "-", os.Getenv("NOT_FOUND")+utils.Reset)
		http.Error(w, os.Getenv("NOT_FOUND"), http.StatusNotFound)
		return
	}

	posts, err := p.PostService.FindAllPosts()
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusInternalServerError, "-", err.Error(), utils.Reset)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set(os.Getenv("CONTENT_TYPE"), os.Getenv("APPLICATION_JSON"))
	w.WriteHeader(http.StatusOK)
	utils.LoggerInfo.Println(utils.Info, http.StatusOK, "-", "Posts found"+utils.Reset)
	if posts != nil {
		err = json.NewEncoder(w).Encode(map[string]interface{}{
			"status": http.StatusOK,
			"posts":  posts,
		})
	} else {
		err = json.NewEncoder(w).Encode(map[string]interface{}{
			"status": http.StatusOK,
			"posts":  "No posts found",
		})
	}
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusInternalServerError, "-", err.Error(), utils.Reset)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (p *PostController) CountAllPosts(w http.ResponseWriter, r *http.Request) {
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

	if r.URL.Path != os.Getenv("DEFAULT_API_LINK")+"/count-posts" {
		utils.LoggerError.Println(utils.Error, http.StatusNotFound, "-", os.Getenv("NOT_FOUND")+utils.Reset)
		http.Error(w, os.Getenv("NOT_FOUND"), http.StatusNotFound)
		return
	}

	count, err := p.PostService.CountAllPosts()
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusInternalServerError, "-", err.Error(), utils.Reset)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set(os.Getenv("CONTENT_TYPE"), os.Getenv("APPLICATION_JSON"))
	w.WriteHeader(http.StatusOK)
	utils.LoggerInfo.Println(utils.Info, http.StatusOK, "-", "Posts counted"+utils.Reset)
	err = json.NewEncoder(w).Encode(map[string]interface{}{
		"status": http.StatusOK,
		"count":  count,
	})
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusInternalServerError, "-", err.Error(), utils.Reset)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (p *PostController) GetPostByUser(w http.ResponseWriter, r *http.Request) {
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

	if !strings.HasPrefix(r.URL.Path, os.Getenv("DEFAULT_API_LINK")+"/user-posts/") {
		utils.LoggerError.Println(utils.Error, http.StatusNotFound, "-", os.Getenv("NOT_FOUND")+utils.Reset)
		http.Error(w, os.Getenv("NOT_FOUND"), http.StatusNotFound)
		return
	}

	id, privacy, err := utils.ExtractIdAndPrivacyFromRequest(r)
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusBadRequest, "-", err.Error(), utils.Reset)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userRepo := repository.NewUserRepoImpl(*globale.DB)
	userService := impl.UserServiceImpl{Repository: userRepo}

	limit, err := userService.CountUsers()
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusInternalServerError, "-", err.Error(), utils.Reset)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if id > limit {
		utils.LoggerError.Println(utils.Error, http.StatusNotFound, "-", os.Getenv("NOT_FOUND")+utils.Reset)
		http.Error(w, os.Getenv("NOT_FOUND"), http.StatusNotFound)
		return
	}

	posts, err := p.PostService.FindPostsByUserID(id, privacy)
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusInternalServerError, "-", err.Error(), utils.Reset)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set(os.Getenv("CONTENT_TYPE"), os.Getenv("APPLICATION_JSON"))
	w.WriteHeader(http.StatusOK)
	utils.LoggerInfo.Println(utils.Info, http.StatusOK, "-", "Posts found"+utils.Reset)
	if posts != nil {
		err = json.NewEncoder(w).Encode(map[string]interface{}{
			"status": http.StatusOK,
			"posts":  posts,
		})
	} else {
		err = json.NewEncoder(w).Encode(map[string]interface{}{
			"status": http.StatusOK,
			"posts":  "No posts found",
		})
	}
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusInternalServerError, "-", err.Error(), utils.Reset)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (p *PostController) GetPostByCategory(w http.ResponseWriter, r *http.Request) {
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

	if !strings.HasPrefix(r.URL.Path, os.Getenv("DEFAULT_API_LINK")+"/category-posts/") {
		utils.LoggerError.Println(utils.Error, http.StatusNotFound, "-", os.Getenv("NOT_FOUND")+utils.Reset)
		http.Error(w, os.Getenv("NOT_FOUND"), http.StatusNotFound)
		return
	}

	id, err := utils.ExtractIDFromRequest(r)
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusBadRequest, "-", err.Error(), utils.Reset)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	catRepo := repository.NewCategoryRepoImpl(*globale.DB)
	catService := impl.CategoryServiceImpl{Repository: catRepo}

	limit, err := catService.CountAllCategories()
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusInternalServerError, "-", err.Error(), utils.Reset)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if id > limit {
		utils.LoggerError.Println(utils.Error, http.StatusNotFound, "-", os.Getenv("NOT_FOUND")+utils.Reset)
		http.Error(w, os.Getenv("NOT_FOUND"), http.StatusNotFound)
		return
	}

	posts, err := p.PostService.FindPostByCategory(id)
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusInternalServerError, "-", err.Error(), utils.Reset)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set(os.Getenv("CONTENT_TYPE"), os.Getenv("APPLICATION_JSON"))
	w.WriteHeader(http.StatusOK)
	utils.LoggerInfo.Println(utils.Info, http.StatusOK, "-", "Posts found"+utils.Reset)
	if posts != nil {
		err = json.NewEncoder(w).Encode(map[string]interface{}{
			"status": http.StatusOK,
			"posts":  posts,
		})
	} else {
		err = json.NewEncoder(w).Encode(map[string]interface{}{
			"status": http.StatusOK,
			"posts":  "No posts found",
		})
	}
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusInternalServerError, "-", err.Error(), utils.Reset)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (p *PostController) PostsRoutes(routes *http.ServeMux) *http.ServeMux {
	err := utils.Environment()
	if err != nil {
		utils.LoggerError.Println(utils.Error, http.StatusInternalServerError, "-", err.Error(), utils.Reset)
		return routes
	}

	routes.HandleFunc(os.Getenv("DEFAULT_API_LINK")+"/post", p.CreatePost)
	routes.HandleFunc(os.Getenv("DEFAULT_API_LINK")+"/update-post/", p.UpdatePost)
	routes.HandleFunc(os.Getenv("DEFAULT_API_LINK")+"/delete-post/", p.DeletePost)
	routes.HandleFunc(os.Getenv("DEFAULT_API_LINK")+"/get-post/", p.GetPost)
	routes.HandleFunc(os.Getenv("DEFAULT_API_LINK")+"/posts", p.GetAllPosts)
	routes.HandleFunc(os.Getenv("DEFAULT_API_LINK")+"/count-posts", p.CountAllPosts)
	routes.HandleFunc(os.Getenv("DEFAULT_API_LINK")+"/user-posts/", p.GetPostByUser)
	routes.HandleFunc(os.Getenv("DEFAULT_API_LINK")+"/category-posts/", p.GetPostByCategory)

	return routes
}
