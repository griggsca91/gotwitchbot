package main

import (
	"fmt"
	"log"
	//	"os"
	"bufio"
	"net/textproto"
	"twitchbot"
)

const (
	oauth = ""
	nick  = ""
)

func main() {
	bot := &twitchbot.Bot{
		server:  "irc.twitch.tv",
		port:    "6667",
		nick:    nick,
		channel: "lck1",
		conn:    nil,
	}

	pass := "oauth:" + oauth

	bot.Connect()

	bot.Chat(fmt.Sprintf("USER %s 8 * :%s", bot.nick, bot.nick))
	bot.Chat(fmt.Sprintf("PASS %s", pass))
	bot.Chat(fmt.Sprintf("NICK %s", bot.nick))
	bot.Chat(fmt.Sprintf("JOIN #%s", bot.channel))
	bot.Chat("CAP REQ :twitch.tv/membership")
	bot.Chat("CAP REQ :twitch.tv/tags")
	bot.Chat("CAP REQ :twitch.tv/commands")

	reader := bufio.NewReader(bot.conn)
	tp := textproto.NewReader(reader)

	for {
		line, err := tp.ReadLine()

		if err != nil {
			log.Printf("ERROR")
			break
		}
		log.Println(line)

	}

}
