package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joeytatu/warn/weatherwarnings"
)

func writeWarningsToFile(outputFile *os.File, warn []weatherwarnings.Warning) {
	for _, w := range warn {
		outputFile.WriteString("Title: " + w.Title + "\n")
		outputFile.WriteString("Description: " + w.Description + "\n")
		outputFile.WriteString("Valid: " + w.Valid + "\n")
		outputFile.WriteString("Issued: " + w.Issued + "\n")

		outputFile.WriteString("\n#################################\n\n")
	}
}

func main() {
	var day string

	fmt.Print("Enter the day: \"today\", \"tomorrow\", \"dayAfterTomorrow\":\n> ")
	_, err := fmt.Scanf("%s", &day)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	warn, err := weatherwarnings.GetWarnings(day)
	if err != nil {
		log.Fatal(err)
	}

	// Open a file for writing
	file, err := os.Create(day + "warnings.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writeWarningsToFile(file, warn)

	fmt.Println("Output has been written to warnings.txt")
}
