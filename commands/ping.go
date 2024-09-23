package commands

import (
	"github.com/bwmarrin/discordgo"
	"discordBot/gogo/eventHandler"
)

func init() {
	events.RegisterCommand("ping", Ping)
}

func Ping(session *discordgo.Session, message *discordgo.MessageCreate, args []string) {
	session.ChannelMessageSend(message.ChannelID, "Pong!")
}

