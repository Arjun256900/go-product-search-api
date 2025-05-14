package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/arjun256900/go-product-search-api/internal/api/handlers"
	"github.com/arjun256900/go-product-search-api/internal/search"
	"github.com/arjun256900/go-product-search-api/pkg/graceful"
	"github.com/go-chi/chi/v5"
)

func main() {
	// Initializing the search Service (this generates & indexes products)
	const numProducts = 1000000
	svc, err := search.NewService(numProducts)
	if err != nil {
		log.Fatalf("could not create search service: %v", err)
	}

	// Creating router and mounting handlers
	r := chi.NewRouter()

	// Handlers from the handler folder
	r.Get("/health", handlers.HealthHandler)
	r.Get("/search", handlers.SearchHandler(svc))

	// Set up HTTP server
	httpSrv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	// Listen for interrupt signal (Ctrl + C)
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Start server in goroutine
	go func() {
		log.Println("🚀 Server listening on :8080")
		if err := httpSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server error: %v", err)
		}
	}()

	// Graceful shutdown when signal arrives
	if err := graceful.Shutdown(ctx, httpSrv, 5*time.Second); err != nil {
		log.Fatalf("graceful shutdown failed: %v", err)
	}

	log.Println("🛑 Server shut down cleanly")
}
