package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Main struct {
	Result Result `json:"result"`
}

type Result struct {
	Emails string `json:"emails"`
}

func main() {
	var domain string

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		os.Exit(1)
	}

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		fmt.Println("API_KEY not found in .env file.")
		os.Exit(1)
	}

	fmt.Print("Enter the domain:\n> ")
	fmt.Scan(&domain)

	url := "https://api.apilayer.com/whois/query?domain=" + domain

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		os.Exit(1)
	}
	req.Header.Set("apikey", apiKey)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		os.Exit(1)
	}
	defer res.Body.Close()

	if res.StatusCode == 400 {
		fmt.Println("Domain not found or not supported!")
		os.Exit(1)
	}

	if res.StatusCode != http.StatusOK {
		if res.StatusCode == 400 {
			fmt.Println("Domain not found or not supported!")
			os.Exit(1)
		} else {
			fmt.Printf("Error: Request failed with status %d\n", res.StatusCode)
			os.Exit(1)
		}
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		os.Exit(1)
	}

	var main Main
	if err := json.Unmarshal(body, &main); err != nil {
		fmt.Println("Error decoding JSON:", err)
		os.Exit(1)
	}

	if main.Result.Emails == "" {
		fmt.Println("No abuse email found!")
	} else {
		// Access the Emails field and print it
		fmt.Printf("Email: %v\n", main.Result.Emails)
	}
}
