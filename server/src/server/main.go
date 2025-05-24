package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Priyanshu9898/Strato-Cloud-Assignment/internal"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := ":" + port

	srv := &http.Server{
		Addr:    addr,
		Handler: internal.NewRouter(),
	}

	log.Printf("🚀 Server listening on %s", addr)
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("server error: %v", err)
	}
}
