package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file. Error:", err)
	}

	apiToken := os.Getenv("OPENWEATHERAPP_API")

}
