package ent

import (
	"context"
	"entgo.io/ent/dialect"
	"fmt"
	"os"
	"tsukuyomi/ent/migrate"
	"tsukuyomi/log"
)

var database *Client

func Init() error {
	var err error
	logger := log.GetLogger()
	database, err = Open(dialect.Postgres, fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_NAME"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_SSLMODE"),
	))
	if err != nil {
		logger.Error(fmt.Sprintf("failed connect database: %v", err))
		return err
	}
	err = createSchema(database)
	if err != nil {
		return err
	}
	logger.Info("Finish initialize database")
	return nil
}

func Close(client *Client) {
	logger := log.GetLogger()
	err := client.Close()
	if err != nil {
		logger.Error(fmt.Sprintf("failed disconnect database: %v", err))
	}
	logger.Info("Closed Database connection")
}

func createSchema(client *Client) error {
	logger := log.GetLogger()
	if err := client.Schema.Create(context.Background(), migrate.WithForeignKeys(false)); err != nil {
		logger.Info(fmt.Sprintf("failed create schema: %v", err))
		return err
	}
	logger.Info("Success create schema")
	return nil
}

func GetClient() *Client {
	return database
}
