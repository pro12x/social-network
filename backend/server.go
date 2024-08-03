package main

import (
	"backend/pkg/db/sqlite"
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

	db, err := sqlite.Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	if err := sqlite.Migrate(db.GetDB()); err != nil {
		return err
	}

	// Create a new ServerMux
	mux := http.NewServeMux()

	// Initializing repositories
	userRepo := repository.NewUserRepoImpl(*db)
	followRepo := repository.NewFollowRepoImpl(*db)

	// Initializing services
	userService := impl.UserServiceImpl{Repository: userRepo}
	followService := impl.FollowServiceImpl{Repository: followRepo}

	// Initializing controllers
	userController := web.UserController{UserService: userService}
	followController := web.FollowController{FollowService: followService}

	// Routes
	mux = userController.UsersRoutes(mux)
	mux = followController.FollowsRoutes(mux)

	// Create a new handler
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		_, err := w.Write([]byte("Hello Janel"))
		if err != nil {
			return
		}
	})

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
