package main

import (
	"bufio"
	"log"
	"net/textproto"
	"twitchbot"
)

const (
	oauth = ""
	nick  = "bad_hombres_bot"
)

func main() {
	bot := &twitchbot.Bot{
		Server:  "irc.twitch.tv",
		Port:    "6667",
		Nick:    nick,
		Channel: "03f001",
		Conn:    nil,
		Pass:    "oauth:" + oauth,
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
