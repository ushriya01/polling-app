package main

import (
	"context"
	"log"
	"net/http"
	"poll-app/internal/db"
	"poll-app/internal/server"
	"time"

	"github.com/rs/cors"
)

func main() {
	dbDSN := "postgres://postgres:password@localhost:5432/poll_db?sslmode=disable"
	client, err := db.InitializeDB(context.Background(), dbDSN)
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	defer client.Close()
	ctx := context.WithValue(context.Background(), "client", client)
	srv := server.NewServer(":8080", ctx)
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           int(12 * time.Hour.Seconds()),
	})
	handler := corsHandler.Handler(srv.SetupRoutes())
	log.Printf("Server listening on %s\n", srv.Addr)
	log.Fatal(http.ListenAndServe(srv.Addr, handler))
}
