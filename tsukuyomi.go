package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"tsukuyomi/ent"
	"tsukuyomi/server"
)

func main() {
	// Logger setting
	logger, _ := zap.NewDevelopment()
	logger.Info("Running application")

	// Initialize Database
	if err := ent.Init(); err != nil {
		fmt.Println("aaa")
		panic(err)
	}

	// Initialize Router
	if err := server.Init(); err != nil {
		panic(err)
	}
	logger.Info("Stopping application")
}
