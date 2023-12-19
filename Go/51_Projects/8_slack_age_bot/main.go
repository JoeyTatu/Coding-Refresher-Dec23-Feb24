package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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
		fmt.Println("Error loading .env file. Err:", err)
	}

	botToken := os.Getenv("SLACK_BOT_TOKEN")
	appToken := os.Getenv("SLACK_APP_TOKEN")

	// log.Println("Bot Token:", botToken)
	// log.Println("App Token:", appToken)

	bot := slacker.NewClient(botToken, appToken)

	go printCommandEvents(bot.CommandEvents())

	bot.Command("DOB <day> <month> <year>", &slacker.CommandDefinition{
		Description: "Date of birth calculator",
		Examples:    []string{"DOB 01 Jan 2000"},
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			// Extract day, month, and year parameters
			day := request.Param("day")
			month := request.Param("month")
			year := request.Param("year")

			// Convert parameters to integers
			if len(day) == 1 {
				day = "0" + day
			}
			dayInt, err := strconv.Atoi(day)
			if err != nil {
				fmt.Println("Error:", err)
				response.Reply("Invalid day format. Please provide a valid number.")
				return
			}

			// Convert month to uppercase and parse as a time.Month
			month = strings.ToUpper(month)
			monthInt, err := parseMonth(month)
			if err != nil {
				fmt.Println("Error:", err)
				response.Reply("Invalid month format. Please provide the first three letters of the month.")
				return
			}

			yearInt, err := strconv.Atoi(year)
			if err != nil {
				fmt.Println("Error:", err)
				response.Reply("Invalid year format. Please provide a valid number.")
				return
			}

			// Calculate age based on DOB
			currentTime := time.Now()
			age := currentTime.Year() - yearInt
			if currentTime.YearDay() < time.Date(currentTime.Year(), monthInt, dayInt, 0, 0, 0, 0, time.UTC).YearDay() {
				age--
			}

			// Respond with the calculated age
			month = strings.ToLower(month)
			month = strings.ToUpper(month[:1]) + strings.ToLower(month[1:])
			res := fmt.Sprintf("DOB entered: %s %s %s. Age is %d", day, month, year, age)
			response.Reply(res)
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

// parseMonth converts the first three letters of the month to a time.Month
func parseMonth(month string) (time.Month, error) {
	switch month {
	case "JAN":
		return time.January, nil
	case "FEB":
		return time.February, nil
	case "MAR":
		return time.March, nil
	case "APR":
		return time.April, nil
	case "MAY":
		return time.May, nil
	case "JUN":
		return time.June, nil
	case "JUL":
		return time.July, nil
	case "AUG":
		return time.August, nil
	case "SEP":
		return time.September, nil
	case "OCT":
		return time.October, nil
	case "NOV":
		return time.November, nil
	case "DEC":
		return time.December, nil
	default:
		return 0, fmt.Errorf("unknown month: %s", month)
	}
}
