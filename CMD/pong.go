package commands


import (
	"github.com/bwmarrin/discordgo"
)


func Pong(session *discordgo.Session, message *discordgo.MessageCreate) {
	session.ChannelMessageSend(message.ChannelID, "Ping!")
}
