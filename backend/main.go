package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"tsukuyomi/ent"
	"tsukuyomi/server"
)

func main() {
	fmt.Println("Running application")
	// Logger setting

	// Initialize Database
	if err := ent.Init(); err != nil {
		fmt.Println("aaa")
		panic(err)
	}

	// Initialize Router
	if err := server.Init(); err != nil {
		panic(err)
	}
	fmt.Println("Stopping application")
}
