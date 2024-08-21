package main

import (
	"backend/pkg/db/sqlite"
	"backend/pkg/globale"
	"backend/pkg/middleware"
	"backend/pkg/repository"
	"backend/pkg/service/impl"
	"backend/pkg/utils"
	"backend/pkg/web"
	"errors"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Intialize the logger
	utils.InitLogger()
	utils.Welcome()
	utils.LoggerInfo.Println("Starting the server...")

	// signals channel
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, os.Kill, syscall.SIGINT, syscall.SIGTERM)

	// Goroutine to handle signals
	go func() {
		<-signals
		utils.CleanUp()
		os.Exit(0)
	}()

	// Rotate the log file
	utils.RotateLogFile()

	// Start the server
	err := StartServer(os.Args[1:])
	if err != nil {
		log.Println(err)
		utils.LoggerError.Println(err)
		return
	}
	select {}
}

func StartServer(tab []string) error {
	// Check arguments
	if len(tab) != 0 {
		return errors.New("too many arguments")
	}

	// Check if the .env file exists
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		return errors.New("the .env file does not exist")
	}

	// Read the .env file
	err := utils.Environment()
	if err != nil {
		return err
	}

	globale.DB, err = sqlite.Connect()
	if err != nil {
		return err
	}
	defer globale.DB.Close()

	if err := sqlite.Migrate(globale.DB.GetDB()); err != nil {
		return err
	}

	// Create a new ServerMux
	mux := http.NewServeMux()

	mux = Routes(mux)

	// Welcome handler
	mux.HandleFunc("/swagger", web.HomeController)

	// Add the middleware
	wrappedMux := middleware.LoggingMiddleware(mux)
	wrappedMux = middleware.CORSMiddleware(wrappedMux)
	// wrappedMux = middleware.AuthMiddleware(wrappedMux)
	wrappedMux = middleware.ErrorMiddleware(wrappedMux)

	// Set the server structure
	server := &http.Server{
		Addr:    ":" + os.Getenv("PORT"),
		Handler: wrappedMux,
	}

	// Start the server
	log.Println(utils.Info + "The server is listening at http://localhost:" + os.Getenv("PORT") + utils.Reset)
	utils.LoggerInfo.Println(utils.Info + "The server is listening at http://localhost:" + os.Getenv("PORT") + utils.Reset)
	err = server.ListenAndServe()
	return err
}

func Routes(routes *http.ServeMux) *http.ServeMux {
	// Initializing repositories
	userRepo := repository.NewUserRepoImpl(*globale.DB)
	followRepo := repository.NewFollowRepoImpl(*globale.DB)
	catRepo := repository.NewCategoryRepoImpl(*globale.DB)
	postRepo := repository.NewPostRepoImpl(*globale.DB)
	commentRepo := repository.NewCommentRepoImpl(*globale.DB)

	// Initializing services
	userService := impl.UserServiceImpl{Repository: userRepo}
	followService := impl.FollowServiceImpl{Repository: followRepo}
	catService := impl.CategoryServiceImpl{Repository: catRepo}
	postService := impl.PostServiceImpl{Repository: postRepo}
	commentService := impl.CommentServiceImpl{Repository: commentRepo}

	// Initializing controllers
	userController := web.UserController{UserService: userService}
	followController := web.FollowController{FollowService: followService}
	categoryController := web.CategoryController{CategoryService: catService}
	postController := web.PostController{PostService: postService}
	commentController := web.CommentController{CommentService: commentService}

	// Routes
	routes = userController.UsersRoutes(routes)
	routes = followController.FollowsRoutes(routes)
	routes = categoryController.CategoriesRoutes(routes)
	routes = postController.PostsRoutes(routes)
	routes = commentController.CommentsRoutes(routes)

	return routes
}
