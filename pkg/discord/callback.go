package discord

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"syscall"
)

var session *discordgo.Session

func Init() error {
	var err error
	token := fmt.Sprintf("Bot %s", os.Getenv("DISCORD_TOKEN"))
	session, err = discordgo.New(token)
	if err != nil {
		return err
	}
	session.AddHandler(OnMessageCreate)
	err = session.Open()
	if err != nil {
		return err
	}
	return nil
}

func Close() {
	stopBot := make(chan os.Signal, 1)
	signal.Notify(stopBot, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-stopBot
	err := session.Close()
	if err != nil {
		return
	}
}
