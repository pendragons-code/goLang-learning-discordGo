package commands

import (
	"discordBot/gogo/eventHandler"
	"io"
	"log"
	"net/http"
	"encoding/json"
	"github.com/bwmarrin/discordgo"
)

func init() {
	events.RegisterCommand("dog", randomDog)
}



// response structure
type DogResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func randomDog(session *discordgo.Session, message *discordgo.MessageCreate, args []string) {
	

	// DOGGOOOOOO
	resp, err := http.Get("https://dog.ceo/api/breeds/image/random")
	if err != nil {
		log.Fatalln(err)
	}



	// response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	// Parse the JSON response
	var dogResponse DogResponse
	err = json.Unmarshal(body, &dogResponse)
	if err != nil {
		log.Fatalln(err)
	}

	// Check if the status is successful
	if dogResponse.Status == "success" {
		session.ChannelMessageSend(message.ChannelID, dogResponse.Message) // Send the dog image URL
	} else {
		session.ChannelMessageSend(message.ChannelID, "Sorry, I couldn't fetch a dog image!")
	}
}

