package ent

import (
	"context"
	"entgo.io/ent/dialect"
	"log"
	"tsukuyomi/ent/migrate"
)

var database *Client

func Init() error {
	client, err := Open(dialect.Postgres,
		"host=localhost port=5432 user=postgres dbname=postgres password=postgres sslmode=disable")
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
