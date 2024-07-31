package middleware

import (
	"backend/pkg/session"
	"backend/pkg/utils"
	"log"
	"net/http"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request from", r.RemoteAddr, "to", r.URL)
		utils.Logger.Println("Request from", r.RemoteAddr, "to", r.URL)

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

func ErrorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Recover from panics
		defer func() {
			if err := recover(); err != nil {
				log.Println(err)
				utils.Logger.Println(err)
				http.Error(w, "Internal server error", http.StatusInternalServerError)
			}
		}()

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

/*func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check the Authentication
		if r.Header.Get("Authorization") != "Bearer token" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		if !isAuthenticated(r) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Do something
		next.ServeHTTP(w, r)
	})
}*/

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := session.GetSessionTokenFromRequest(r)
		if err != nil || token == "" {
			utils.Logger.Println(http.StatusUnauthorized, "Unauthorized")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		_, err = session.GetSession(token)
		if err != nil {
			utils.Logger.Println(http.StatusUnauthorized, "Unauthorized")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set the headers
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Check the request method
		if r.Method == "OPTIONS" {
			return
		}

		// Do something
		next.ServeHTTP(w, r)
	})
}

func isAuthenticated(r *http.Request) bool {
	return r.Header.Get("Authorization") == "Bearer token"
}
