package main

import "github.com/turnage/graw/reddit"
import "fmt"

func initBot() {
	bot, err := reddit.NewBotFromAgentFile("reminderbot.agent", 0)
	harvest, err := bot.Listing("/r/golang", "")
	if err != nil {
       fmt.Println("Failed to fetch /r/golang: ", err)
       return
	}

	for _, post := range harvest.Posts[:5] {
       fmt.Printf("[%s] posted [%s]\n", post.Author, post.Title)
	}
}
