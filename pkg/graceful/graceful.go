package graceful

import (
	"context"
	"log"
	"net/http"
	"time"
)

// Shutdown waits for ctx to be done (e.g. SIGINT), then attempts to
// gracefully shut down srv within timeout, allowing in‑flight requests
// to complete.
func Shutdown(ctx context.Context, srv *http.Server, timeout time.Duration) error {
	// Block until the context is done (Ctrl+C / SIGINT)
	<-ctx.Done()

	log.Println("⚠️  Graceful shutdown initiated")

	// Create a new context with timeout for the shutdown process
	shutdownCtx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// Attempt graceful shutdown
	if err := srv.Shutdown(shutdownCtx); err != nil {
		return err
	}

	log.Println("✅ Server shutdown completed")
	return nil
}
