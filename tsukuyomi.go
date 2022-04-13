package main

import (
	_ "github.com/lib/pq"
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

	// Initialize Database
	if err := ent.Init(); err != nil {
		logger.Panic("Failed initialized database")
		panic(err)
	}
	client := ent.GetClient()
	defer ent.Close(client)

	// Initialize Server
	if err := server.Init(); err != nil {
		logger.Panic("Failed initialized server")
		panic(err)
	}
	logger.Debug("Stopping application")
}
