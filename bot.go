package twitchbot

import (
	"fmt"
	"log"
	"net"
	"time"
	"io/ioutil"
	"encoding/json"
	"regexp"
)

type Bot struct {
	Server  string
	Port    string
	Nick    string
	Channel string
	Conn    net.Conn
	Pass    string
	GiphyAPI string

}

type Message struct {
	raw     string
	channel string
	user    string
}


// Creates a new bot with the settings defined in a configuration file at the filePath location.
func NewWithConfig(filePath string) (*Bot, error) {

	bytes, err := ioutil.ReadFile(filePath)

	if err != nil {
		return nil, err
	}
	
	var dat map[string]interface{}

	if err = json.Unmarshal(bytes, &dat); err != nil {
		return nil, err
	}

	oauth :=  "oauth:" + dat["oauth"].(string)
	channel := dat["channel"].(string)
	nick := dat["nick"].(string)
	giphyAPI := dat["giphy_api"].(string)

	return &Bot{
		Server:  "irc.twitch.tv",
		Port:    "6667",
		Nick:    nick,
		Channel: channel,
		Pass:    oauth,
		GiphyAPI: giphyAPI,
	}, nil

}

// SendMsg sends an unformated message to the IRC connection.
func (bot *Bot) SendMsg(message string) {
	fmt.Fprintln(bot.Conn, message)
	log.Print(message)
}

// Chat sends a PRIVMSG to the channel the bot is connected to.
func (bot *Bot) Chat(message string) {
	msg := fmt.Sprintf("PRIVMSG #%s :%s", bot.Channel, message)
	bot.SendMsg(msg)
}

// ParseLine parses everyline that the bot encounters.  This is used to execute commands that begin with !
func (bot *Bot) ParseLine(message string) {

	re := regexp.MustCompile("PRIVMSG #(.*) :!(.*)")
	results := re.FindAllStringSubmatch(message, -1)	

	if len(results) != 1 {
		return
	}

	switch results[0][2] {
		case "dankmeme":
			bot.CommandDankMeme()			
	}

}

// Connect connects the bot to the IRC channel defined in the config file
func (bot *Bot) Connect() {

	var err error
	log.Printf("Attempting to connect to server...\n")
	bot.Conn, err = net.Dial("tcp", bot.Server+":"+bot.Port)
	if err != nil {
		log.Printf("Unable to connect to Twitch IRC server! Reconnecting in 10 seconds...\n")
		time.Sleep(10 * time.Second)
		bot.Connect()
	}

	log.Printf("Connected to IRC server %s\n", bot.Server)

	bot.SendMsg(fmt.Sprintf("USER %s 8 * :%s", bot.Nick, bot.Nick))
	bot.SendMsg(fmt.Sprintf("PASS %s", bot.Pass))
	bot.SendMsg(fmt.Sprintf("NICK %s", bot.Nick))
	bot.SendMsg(fmt.Sprintf("JOIN #%s", bot.Channel))
	bot.SendMsg("CAP REQ :twitch.tv/membership")
	bot.SendMsg("CAP REQ :twitch.tv/tags")
	bot.SendMsg("CAP REQ :twitch.tv/commands")
}
