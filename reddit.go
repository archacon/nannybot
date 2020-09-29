package main

import "github.com/turnage/graw/reddit"
import "fmt"
import "time"
import "math/rand"

var (
	PostList         []reddit.Post
	UPVOTE_THRESHOLD = int32(20)
)

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

	for _, post := range harvest.Posts[:] {
		fmt.Printf("[%s] posted [%s] %s\n", post.Author, post.Title, post.URL)
	}
}

func getRandomMeme() reddit.Post{
	if(len(PostList) == 0) {
		
	}
    sauce := rand.NewSource(time.Now().UnixNano())
	r := rand.New(sauce)
	return PostList[r.Intn(len(PostList)-1)]
}

func RefreshMemeList() {
	bot, err := reddit.NewBotFromAgentFile("./nannybot.agent", 0)
	if err != nil {
		fmt.Println("Bot unable to initialize: ", err, bot)
	}
	getFreshMemes(bot, []string{
		"/r/dankmemes",
		"/r/wholesomemes",
		"/r/me_irl",
	})
}

func getFreshMemes(bot reddit.Bot, threadList []string) {
	PostList = []reddit.Post{}
	for _, thread := range threadList {
		harvest, err := bot.Listing(thread, "")
		if err != nil {
			fmt.Printf("Failed to fetch %s: %w", thread, err)
			return
		}
		PostList = append(PostList, refinePostList(harvest.Posts, UPVOTE_THRESHOLD)...)
	}
}

func refinePostList(posts []*reddit.Post, threshold int32) (refinedPostList []reddit.Post) {
	for _, post := range posts {
		if post.Ups > threshold {
			refinedPostList = append(refinedPostList, *post)
		}
	}
	return
}
