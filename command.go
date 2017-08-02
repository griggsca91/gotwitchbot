package twitchbot

import (
	"io/ioutil"
	"net/http"
	"log"
	"encoding/json"
)

func (bot *Bot) CommandDankMeme() {
	log.Println("this is what we want")
	resp, err := http.Get("https://api.giphy.com/v1/gifs/random?api_key=" + bot.GiphyAPI + "&tag=dank%20meme")
	if err != nil {
		log.Printf("Error in getting dank meme: %v", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error in reading response for dank meme: %v", err)
		return
	}

	var dat map[string]interface{}
	err = json.Unmarshal(body, &dat)
	if err != nil {
		log.Printf("Error in unmarshling dank meme: %v", err)
		return
	}
	data := dat["data"].(map[string]interface{})
	dankMemeURL := data["image_url"].(string)
	bot.Chat(dankMemeURL)
}
