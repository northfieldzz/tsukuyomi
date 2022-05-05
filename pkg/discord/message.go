package discord

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"regexp"
)

func OnMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}

	var r *regexp.Regexp
	r = regexp.MustCompile(`^Tsukuyomi$`)
	if r.MatchString(m.Content) {
		_ = makeMessageTsukuyomi(s, m)
		return
	}
	r = regexp.MustCompile(`^youtube$`)
	if r.MatchString(m.Content) {
		_ = makeMessageYoutube(s, m)
		return
	}

	r = regexp.MustCompile(`^please music`)
	if r.MatchString(m.Content) {
		_ = ""
	}
}

func makeMessageTsukuyomi(s *discordgo.Session, m *discordgo.MessageCreate) error {
	nickname := m.Author.Username
	member, err := s.State.Member(m.GuildID, m.Author.ID)
	if err == nil && member.Nick != "" {
		nickname = member.Nick
	}
	_, err = s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("なんや、%s", nickname))
	if err != nil {
		return err
	}
	return nil
}

func makeMessageYoutube(s *discordgo.Session, m *discordgo.MessageCreate) error {
	_, err := s.ChannelMessageSend(m.ChannelID, "https://www.youtube.com/watch?v=SEdMmVeMmqc")
	if err != nil {
		return err
	}
	return nil
}
