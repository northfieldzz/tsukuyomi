package main

import (
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"os"
	"tsukuyomi/ent"
	"tsukuyomi/internal/tsukuyomi/server"
	"tsukuyomi/pkg/discord"
	"tsukuyomi/pkg/log"
)

func main() {
	// Logger setting
	if err := log.Init(); err != nil {
		panic(err)
	}
	logger := log.GetLogger()
	logger.Debug("Running application")

	if err := godotenv.Load(fmt.Sprintf("./configs/.env.%s", os.Getenv("ENV_MODE"))); err != nil {
		logger.Panic("Failed load environment")
	}

	// Initialize Database
	if err := ent.Init(); err != nil {
		logger.Panic("Failed initialized database")
	}
	client := ent.GetClient()
	defer ent.Close(client)

	// Initialize Discord client
	if err := discord.Init(); err != nil {
		logger.Panic("Failed initialized discord")
	}
	defer discord.Close()

	// Initialize Server
	if err := server.Init(); err != nil {
		logger.Panic("Failed initialized server")
	}
	logger.Debug("Stopping application")
}
