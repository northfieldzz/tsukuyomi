package ent

import (
	"context"
	"entgo.io/ent/dialect"
	"fmt"
	"log"
	"os"
	"tsukuyomi/ent/migrate"
)

var database *Client

func Init() error {
	client, err := Open(dialect.Postgres, fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_NAME"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_SSLMODE"),
	))
	database = client
	if err != nil {
		return err
	}
	for i := 0; i < 20; i++ {
		if err != nil {
			log.Printf("fialed connection database: %v", err)
			log.Printf("try connect database: %d / 20", i)
		} else {
			break
		}
	}
	defer func(client *Client) {
		err := client.Close()
		if err != nil {
			log.Fatalf("fialed instance database: %v", err)
		}
	}(client)
	ctx := context.Background()
	if err := client.Schema.Create(ctx, migrate.WithForeignKeys(false)); err != nil {
		log.Fatalf("failed create schema: %v", err)
	}
	return nil
}

func GetClient() *Client {
	return database
}
