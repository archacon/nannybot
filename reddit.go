package main

import "github.com/turnage/graw/reddit"
import "fmt"

func initBot() {
	bot, err := reddit.NewBotFromAgentFile("./nannybot.agent", 0)
	if err != nil {
		fmt.Println("Bot unable to initialize: ", err, bot)
	}
	harvest, err := bot.Listing("/r/dankmemes", "")
	if err != nil {
       fmt.Println("Failed to fetch /r/dankmemes: ", err)
       return
	}

	for _, post := range harvest.Posts[:5] {
	   fmt.Printf("[%s] posted [%s] %s\n", post.Author, post.Title, post.URL)
	}
}
