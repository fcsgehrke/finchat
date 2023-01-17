package main

import (
	"log"

	"github.com/fcsgehrke/finchat/pkg/bot"
)

func main() {
	bot, err := bot.NewBot("amqp://guest:guest@localhost:5672/", "stock-request", "stock-response")
	if err != nil {
		log.Fatalf("[ERR] - Couldn't start Bot w/ err: %s", err.Error())
	}

	bot.Start()
}
