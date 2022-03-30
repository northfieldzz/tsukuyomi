package discord

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
)

func onMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	switch {
	case strings.HasPrefix(m.Content, fmt.Sprintf("%s %s", BotName, HelloWorld)):
		sendMessage(s, m.ChannelID, "Hello World!!")
		//case strings.HasPrefix(m.Content, fmt.Sprintf("%s %s", BotName, ChannelVoiceJoin)):
		//	//今いるサーバーのチャンネル情報の一覧を喋らせる処理を書いておきますね
		//	c, err := s.State.Channel(m.ChannelID) //チャンネル取得
		//	if err != nil {
		//		log.Fatalf("error")
		//	}
		//	guildChannels, _ := s.GuildChannels(c.GuildID)
		//	var sendText string
		//	for _, a := range guildChannels {
		//		sendText += fmt.Sprintf("%vチャンネルの%v(IDは%v)\n", a.Type, a.Name, a.ID)
		//	}
		//	sendMessage(s, c, sendText) // チャンネルの名前、ID、タイプ(通話orテキスト)をBOTが話す
		//
		//	//VOICE CHANNEL IDには、botを参加させたい通話チャンネルのIDを代入してください
		//	//コメントアウトされた上記の処理を使うことでチャンネルIDを確認できます
		//	cvsession, _ = s.ChannelVoiceJoin(c)
		//case strings.HasPrefix(m.Content, fmt.Sprintf("%s %s", BotName, ChannelVoiceLeave)):
		//	vcsession.Disconnect()
	}
}

func onVoiceReceived(cv *discordgo.VoiceConnection, vs *discordgo.VoiceSpeakingUpdate) {
	log.Print("おれのおれのおれの話をきけ!")
}

func sendMessage(s *discordgo.Session, channelID string, msg string) {
	_, err := s.ChannelMessageSend(channelID, msg)
	if err != nil {
		log.Print("failed sending message")
	}
}
