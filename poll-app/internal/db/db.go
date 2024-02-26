package db

import (
	"context"
	"log"
	"poll-app/ent"
	"poll-app/ent/migrate"

	_ "github.com/lib/pq"
)

// InitializeDB initializes the database connection and sets up ORM
func InitializeDB(ctx context.Context, dsn string) (*ent.Client, error) {
	client, err := ent.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	if err := client.Schema.Create(ctx, migrate.WithDropIndex(true), migrate.WithDropColumn(true)); err != nil {
		return nil, err
	}
	log.Println("Database schema has been created")
	return client, nil
}
