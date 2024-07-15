package main

import (
	"backend/pkg"
	"backend/pkg/db/sqlite"
	"errors"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"os"
)

func main() {
	// Start the server
	err := StartServer(os.Args[1:])
	if err != nil {
		log.Println(err)
		return
	}
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
	err := pkg.Environment()
	if err != nil {
		return err
	}

	_, err = sqlite.Connect()
	if err != nil {
		return err
	}

	// Create a new ServerMux
	mux := http.NewServeMux()

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

	mux.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/posts" {
			http.NotFound(w, r)
			return
		}

		_, err := w.Write([]byte("Posts here"))
		if err != nil {
			return
		}
	})

	// Add the middleware
	wrappedMux := pkg.LoggingMiddleware(mux)
	wrappedMux = pkg.CORSMiddleware(wrappedMux)
	// wrappedMux = pkg.AuthMiddleware(wrappedMux)
	wrappedMux = pkg.ErrorMiddleware(wrappedMux)

	// Set the server structure
	server := &http.Server{
		Addr:    ":" + os.Getenv("PORT"),
		Handler: wrappedMux,
	}

	// Start the server
	log.Println("The server is listening at http://localhost:" + os.Getenv("PORT"))
	err = server.ListenAndServe()
	return err
}
