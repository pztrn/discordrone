package main

import (
	// stdlib
	"log"
	"os"

	// local
	"go.dev.pztrn.name/discordrone/env"
	"go.dev.pztrn.name/discordrone/message"
)

func main() {
	env.ParseEnv()

	if env.Data.Debug {
		log.Printf("Starting Discordrone with parameters: %+v\n", env.Data)
	}

	err := message.Process()
	if err != nil {
		log.Fatalln("Error appeared:", err)
		os.Exit(1)
	}
}
