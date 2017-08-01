package twitchbot

import (
	"log"
)


type Bot struct {
	server  string
	port    string
	nick    string
	channel string
	conn    net.Conn
}

type Message struct {
  raw string
  channel string
  user string
}


func NewMessage(

func (bot *Bot) ConsoleInput() {
	// _ := bufio.NewReader(os.Stdin)

}

func (bot *Bot) Connect() {

	var err error
	log.Printf("Attempting to connect to server...\n")
	bot.conn, err = net.Dial("tcp", bot.server+":"+bot.port)
	if err != nil {
		log.Printf("Unable to connect to Twitch IRC server! Reconnecting in 10 seconds...\n")
		time.Sleep(10 * time.Second)
		bot.Connect()
	}

	log.Printf("Connected to IRC server %s\n", bot.server)
}

func (bot *Bot) Chat(message string) {
	log.Fprintf(bot.conn, message+"\r\n")
	log.Printf(message + "\r\n")
}