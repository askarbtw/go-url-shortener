package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/askarbtw/url-shortener-golang/config"
	"github.com/askarbtw/url-shortener-golang/controllers"
	"github.com/askarbtw/url-shortener-golang/repositories"
	"github.com/askarbtw/url-shortener-golang/services"
	"github.com/gorilla/mux"
)

// CORS middleware
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

func main() {
	// Load configuration
	conf := config.LoadConfig()

	// Connect to database
	db := config.ConnectDB(conf)
	defer db.Close()

	// Create repository
	urlRepository := repositories.NewURLRepository(db)

	// Create service
	urlService := services.NewURLService(urlRepository)

	// Create controller
	urlController := controllers.NewURLController(urlService, conf.BaseURL)

	// Create router
	router := mux.NewRouter()

	// Apply CORS middleware
	router.Use(corsMiddleware)

	// API routes
	router.HandleFunc("/shorten", urlController.CreateURL).Methods("POST")
	router.HandleFunc("/shorten/{shortCode}", urlController.GetURL).Methods("GET")
	router.HandleFunc("/shorten/{shortCode}", urlController.UpdateURL).Methods("PUT")
	router.HandleFunc("/shorten/{shortCode}", urlController.DeleteURL).Methods("DELETE")
	router.HandleFunc("/shorten/{shortCode}/stats", urlController.GetURLStats).Methods("GET")

	// Redirect route
	router.HandleFunc("/r/{shortCode}", urlController.RedirectURL).Methods("GET")

	// Start server
	srv := &http.Server{
		Addr:    ":" + conf.Port,
		Handler: router,
	}

	// Start the server in a goroutine
	go func() {
		fmt.Printf("Server is running on port %s...\n", conf.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on %s: %v\n", conf.Port, err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited properly")
}
