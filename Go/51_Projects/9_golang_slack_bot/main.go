package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("\n\n********************\n\nCommand Events:")
		// Convert the timestamp string to a float64
		timestamp, err := strconv.ParseFloat(event.Event.TimeStamp, 64)
		if err != nil {
			fmt.Println("Error parsing timestamp:", err)
			continue
		}
		fmt.Println("Timestamp:", time.Unix(int64(timestamp), 0).Format(time.RFC3339))
		fmt.Println("Command:", event.Command)
		fmt.Println("Parameters:", event.Parameters)
		fmt.Println("Data:", event.Event.Data)
		fmt.Println("Event:", "\n\tUser:", event.Event.UserID, "\n\tText:", event.Event.Text, "\n\tType:", event.Event.Type)
		fmt.Println() // New line
	}
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file. Error:", err)
	}

	botUserToken := os.Getenv("BOT_USER_TOKEN")
	socketToken := os.Getenv("SOCKET_TOKEN")

	bot := slacker.NewClient(botUserToken, socketToken)

	go printCommandEvents(bot.CommandEvents())

	bot.Command("Ping!", &slacker.CommandDefinition{
		Description: "Ping-Pong!",
		Examples:    []string{"Ping"},
		Handler: func(bc slacker.BotContext, r slacker.Request, w slacker.ResponseWriter) {
			w.Reply("Pong!")
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	fmt.Println("Starting bot...")
	slackErr := bot.Listen(ctx)
	if slackErr != nil {
		log.Fatal("Error listening to Slack events:", slackErr)
	}
}
