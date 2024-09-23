package commands

import (
	"github.com/bwmarrin/discordgo"
	"discordBot/gogo/eventHandler"
)

func init() {
	events.RegisterCommand("pong", Pong)
}

func Pong(session *discordgo.Session, message *discordgo.MessageCreate, args []string) {
	session.ChannelMessageSend(message.ChannelID, "Ping!")
}

