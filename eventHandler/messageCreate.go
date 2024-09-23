package events

import (
	"log"
	"os"
	"strings"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)



// this has got to be about the most goofy thing I have ever done.
type messageCommands func(*discordgo.Session, *discordgo.MessageCreate)
var commandMap = make(map[string]messageCommands)



// command init
func RegisterCommand(name string, command messageCommands) {
	commandMap[name] = command
}



// messageCreate
func MessageCreate(session *discordgo.Session, message *discordgo.MessageCreate) {
	
	// environment variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// PREFIX
	PREFIX := os.Getenv("PREFIX")
	

	

	// I am not gonna combine these 2 since they are not related and it helps me learn easier.

	// I used message and session here because i came from discord.js
	// this actually helped me and stopped me from doing mental gymnastics
	if message.Author.ID == session.State.User.ID {
		return
	}

	content := message.Content
	// step 1 it treats the content as an array (a string is an array of characters)
	// find the length of the array
	// it slices from 0, and finds the array length
	// after slicing, if the slice of characters do not match the prefix, it returns
	
	// why is the syntax similar to python here lmao
	if len(content) <= len(PREFIX) || content[:len(PREFIX)] != PREFIX {
		return
	}

	// this forces a message with only the prefix to return
	content = content[len(PREFIX):]
	if len(content) < 1 {
		return
	}

	args := strings.Fields(content)
	name := strings.ToLower(args[0])

	// check if command exists
	if cmd, exists := commandMap[name]; exists {
		cmd(session, message)
	} else {
		session.ChannelMessageSend(message.ChannelID, "Unknown command: "+name)
	}
}
