package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
)

func New(ctx context.Context, h http.Handler) error {
	// TODO handle context cancel with signals

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	s := http.NewServeMux()
	s.HandleFunc("/", h.ServeHTTP)

	log.Printf("Starting service http://localhost:%s", port)

	server := &http.Server{Addr: fmt.Sprintf(":%s", port), Handler: s}
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatalln(err)
		}
	}()

	<-ctx.Done()
	return server.Shutdown(ctx)
}
