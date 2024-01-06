package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

type Quiz struct {
	Question string
	Answer   string
}

func getQuestions(fileName string) ([]Quiz, error) {
	fileObj, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("Error opening %s.csv file: %s", fileName, err.Error())
	}
	defer fileObj.Close()

	csvReader := csv.NewReader(fileObj)
	cLines, err := csvReader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("Error reading data in %s.csv file: %s", fileName, err.Error())
	}
	return parseQuiz(cLines), nil
}

func parseQuiz(lines [][]string) []Quiz {
	request := make([]Quiz, len(lines))
	for i := 0; i < len(lines); i++ {
		request[i] = Quiz{
			Question: lines[i][0],
			Answer:   lines[i][1],
		}
	}
	return request
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func main() {
	fileName := flag.String("f", "quiz.csv", "Path of csv file")
	timer := flag.Int("t", 30, "Timer of the Quiz")
	flag.Parse()

	quiz, err := getQuestions(*fileName) // Removed the '&' before fileName
	if err != nil {
		log.Fatalf("Unable to obtain quiz: %s", err.Error()) // Changed log.Fatal to log.Fatalf
	}

	correctAnswers := 0
	timerObj := time.NewTimer(time.Duration(*timer) * time.Second)

	answerChannel := make(chan string)

quizLoop:
	for i, q := range quiz {
		var ans string
		fmt.Printf("Question %d: %s=", i+1, q.Question)

		go func() {
			fmt.Scanf("%s", &ans)
			answerChannel <- ans
		}()

		select {
		case <-timerObj.C:
			fmt.Println()
			break quizLoop

		case iAns := <-answerChannel:
			if iAns == q.Answer {
				correctAnswers++
			}
			if i == len(quiz)-1 {
				close(answerChannel)
			}
		}
	}
	fmt.Printf("You got %d out of %d correct\n", correctAnswers, len(quiz))
	fmt.Println("Press Enter to exit")
	<-answerChannel
}
