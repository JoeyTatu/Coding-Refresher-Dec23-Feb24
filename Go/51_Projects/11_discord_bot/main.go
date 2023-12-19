package main

import (
	"fmt"

	"github.com/joeytatu/discord-bot/bot"
	"github.com/joeytatu/discord-bot/config"
)

func main() {
	err := config.ReadConfig()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	<-make(chan struct{})
	return
}
