package discord

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var (
	token      = fmt.Sprintf("Bot %s", os.Getenv("DISCORD_TOKEN"))
	BotName    = "Tsukuyomi"
	HelloWorld = "!helloworld"
)

func GetSession() (*discordgo.Session, error) {
	session, err := discordgo.New(token)
	if err != nil {
		return nil, err
	}

	return session, nil
}

func init() {
	session, err := GetSession()
	if err != nil {
		fmt.Println(err)
	}

	session.AddHandler(onMessageCreate)
	err = session.Open()
	if err != nil {
		// TODO:
		fmt.Println(err)
	}

	defer func(discord *discordgo.Session) {
		err := discord.Close()
		if err != nil {
			log.Println("failed close discord session")
		}
	}(session)
	stopBot := make(chan os.Signal, 1)
	signal.Notify(stopBot, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-stopBot
}
