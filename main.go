package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/yanzay/tbot/v2"
)

func main() {
	
	var (
		// inputs
		token  = os.Getenv("INPUT_TOKEN")
		chat   = os.Getenv("INPUT_CHAT")
		status = os.Getenv("INPUT_STATUS")
		message = os.Getenv("INPUT_MESSAGE")

		// github env
		//workflow = os.Getenv("GITHUB_WORKFLOW")
		repo     = os.Getenv("GITHUB_REPOSITORY")
		//commit   = os.Getenv("GITHUB_SHA")
		
	)

	if token == "" {
		log.Fatal("token input is required")
	}
	if chat == "" {
		log.Fatal("chat input is required")
	}
	if status == "" {
		log.Fatal("status input is required")
	}

	c := tbot.NewClient(token, http.DefaultClient, "https://api.telegram.org")

	link := fmt.Sprintf("https://github.com/%s/pulls", repo)
	msg := fmt.Sprintf(`%s ([%s](%s))`,  message, repo, link)

	_, err := c.SendMessage(chat, msg, tbot.OptParseModeMarkdown)
	if err != nil {
		log.Fatalf("unable to send message: %v", err)
	}
}
