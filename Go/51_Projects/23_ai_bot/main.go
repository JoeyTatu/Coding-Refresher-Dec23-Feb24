package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Krognol/go-wolfram"
	"github.com/joho/godotenv"
	"github.com/shomali11/slacker"
	"github.com/tidwall/gjson"
	witai "github.com/wit-ai/wit-go"
)

func printCommandEvents(ch <-chan *slacker.CommandEvent) { // ch = analyticsChannel
	for event := range ch {
		fmt.Println("Command Events")
		fmt.Println("-- Timestamp:", event.Timestamp)
		fmt.Println("-- Command:", event.Command)
		fmt.Println("-- Parameters:", event.Parameters)
		fmt.Println("-- Event:", event.Event)
		fmt.Println()
	}
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file. Error:", err)
	}

	slackBotToken := os.Getenv("SLACK_BOT_TOKEN")
	slackAppToken := os.Getenv("SLACK_APP_TOKEN")
	witAiToken := os.Getenv("WIT_AI_TOKEN")
	wolframAppId := os.Getenv("WOLFRAM_APP_ID")

	bot := slacker.NewClient(slackBotToken, slackAppToken)
	client := witai.NewClient(witAiToken)
	wolframClient := &wolfram.Client{
		AppID: wolframAppId,
	}

	go printCommandEvents(bot.CommandEvents())

	bot.Command("<message>", &slacker.CommandDefinition{
		Description: "Send a question to Wolfram",
		// Example: "Who is the president of Ireland?",
		Handler: func(bc slacker.BotContext, r slacker.Request, w slacker.ResponseWriter) {
			query := r.Param("message")

			msg, err := client.Parse(&witai.MessageRequest{
				Query: query,
			})
			if err != nil {
				log.Fatal("Error:", err)
			}

			data, err := json.MarshalIndent(msg, "", "    ")
			if err != nil {
				log.Fatal("Error:", err)
			}

			rough := string(data[:])
			value := gjson.Get(rough, "entities.wolfram_search_query.0.value")
			answer := value.String()

			resp, err := wolframClient.GetSpokentAnswerQuery(answer, wolfram.Metric, 1000)
			if err != nil {
				log.Fatal("An error occured with the Wolfram Client - GetSpokentAnswerQuery")
			}

			if strings.Contains(answer, "<") {
				resp = ":)"
			}

			fmt.Println("Value: ", value, "\nResposne:", resp)
			w.Reply(resp)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = bot.Listen(ctx)
	if err != nil {
		log.Fatal("Error:", err)
	}
}
