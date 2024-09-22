package main

import (
	"discordBot/gogo/eventHandler"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	
	// environment variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	TOKEN := os.Getenv("TOKEN")


	
	// Discord Client
	dg, err := discordgo.New("Bot " + TOKEN)
	if err != nil {
		fmt.Println("error creating instance,", err)
		return
	}
	dg.Identify.Intents = discordgo.IntentsGuildMessages



	// events
	messageCreate := events.MessageCreate
	dg.AddHandler(messageCreate)



	// initial connection
	err = dg.Open()
	if err != nil {
		fmt.Println("connection init failed,", err)
		return
	}



	// graceful Killing
	fmt.Println("Bot is now up")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
	dg.Close()
}
