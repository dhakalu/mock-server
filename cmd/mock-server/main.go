package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	port := os.Getenv("MOCK_SERVER_PORT")
	if port == "" {
		port = "8090"
	}
	run(port)
}

func run(port string) {

	interruptionContext, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	server := http.Server{
		Addr:    ":" + port,
		Handler: handler(),
	}
	go func() {
		slog.Info("Server started", "port", port)
		if err := server.ListenAndServe(); err != nil {
			slog.Error("Server error", "error", err)
		}
	}()

	<-interruptionContext.Done()
	shutdownContext, cancelFunc := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancelFunc()
	err := server.Shutdown(shutdownContext)
	if err != nil {
		slog.Error("Server shutdown error", "error", err)
	}
	slog.Info("Server stopped")

}

func handler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Apply CORS headers
		cors(w, r)

		// Handle preflight OPTIONS request
		if r.Method == "OPTIONS" {
			return
		}

		path := r.URL.Path

		// read the path in mock-responses folder
		filePath := "mock-responses" + path + ".json"
		data, err := os.ReadFile(filePath)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	})
}

func cors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusNoContent)
	}
}
