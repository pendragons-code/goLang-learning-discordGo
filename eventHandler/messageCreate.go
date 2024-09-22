package events

import (
	"github.com/bwmarrin/discordgo"
)

func MessageCreate(session *discordgo.Session, message *discordgo.MessageCreate) {
	
	// I used message and session here because i came from discord.js
	// this actually helped me and stopped me from doing mental gymnastics
	if message.Author.ID == session.State.User.ID {
		return
	}

	if message.Content == "ping" {
		session.ChannelMessageSend(message.ChannelID, "Pong!")
	}
}
