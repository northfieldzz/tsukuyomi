package main

import (
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"os"
	"tsukuyomi/controllers/discord"
	"tsukuyomi/ent"
	"tsukuyomi/log"
	"tsukuyomi/server"
)

func main() {
	// Logger setting
	if err := log.Init(); err != nil {
		panic(err)
	}
	logger := log.GetLogger()
	logger.Debug("Running application")

	if err := godotenv.Load(fmt.Sprintf("./.env.%s", os.Getenv("ENV_MODE"))); err != nil {
		logger.Error("Failed load environment")
		panic(err)
	}

	// Initialize Database
	if err := ent.Init(); err != nil {
		logger.Error("Failed initialized database")
		panic(err)
	}
	client := ent.GetClient()
	defer ent.Close(client)

	if err := discord.Init(); err != nil {
		logger.Error("Failed initialized discord")
		panic(err)
	}
	defer discord.Close()

	// Initialize Server
	if err := server.Init(); err != nil {
		logger.Error("Failed initialized server")
		panic(err)
	}
	logger.Debug("Stopping application")
}
