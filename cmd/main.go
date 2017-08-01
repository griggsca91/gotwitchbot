package main

import (
	"bufio"
	"log"
	"net/textproto"
	"twitchbot"
)

func main() {

	bot, err := twitchbot.NewWithConfig("../bot.conf")

	if err != nil {
		log.Fatalf("There was an error: %v", err)
	}

	bot.Connect()

	reader := bufio.NewReader(bot.Conn)
	tp := textproto.NewReader(reader)

	bot.Chat("Hi everyone, I'm bad_hombres bot and I have entered this channel.  I have no commands yet Kappa. Golang RULES!!")
	for {
		line, err := tp.ReadLine()

		if err != nil {
			log.Fatalf("ERROR: %v", err)
		}
		bot.ParseLine(line)
	}

}
