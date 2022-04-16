package discord

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)

func OnMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}
	nick := m.Author.Username
	member, err := s.State.Member(m.GuildID, m.Author.ID)
	if err == nil && member.Nick != "" {
		nick = member.Nick
	}
	fmt.Println("< " + m.Content + " by " + nick)

	if m.Content == "ああ言えば" {
		_, err := s.ChannelMessageSend(m.ChannelID, "こう言う")
		if err != nil {
			return
		}
		fmt.Println("> こう言う")
	}
	if strings.Contains(m.Content, "www") {
		_, err2 := s.ChannelMessageSend(m.ChannelID, "lol")
		if err2 != nil {
			return
		}
		fmt.Println("> lol")
	}
}
