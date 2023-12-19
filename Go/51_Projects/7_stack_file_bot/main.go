package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func main() {
	err := godotenv.Load(".tokens")
	if err != nil {
		fmt.Print("Error loading .tokens file. Err:", err)
	}

	token := os.Getenv("SLACK_BOT_TOKEN")
	channelId := os.Getenv("CHANNEL_ID")

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <file_path>")
		os.Exit(1)
	}

	filePath := os.Args[1]

	api := slack.New(token)
	channelArr := []string{channelId}
	fileArr := []string{filePath}

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File:     fileArr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Print("Error uploading file. Err:", err)
			return
		}
		fmt.Printf("Name: %s, URL: %s\n", file.Name, file.Permalink)
	}
}
