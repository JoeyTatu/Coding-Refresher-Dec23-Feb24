package weatherwarnings

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type Warning struct {
	Title       string
	Description string
	Valid       string
	Issued      string
}

func getDayAfterTomorrow(day string) (string, error) {
	day = strings.ToLower(day)

	currentTime := time.Now()

	dayAfterTomorrow := currentTime.Add(48 * time.Hour)
	nameOfDay := dayAfterTomorrow.Weekday().String()
	nameOfDay = strings.ToLower(nameOfDay)

	return nameOfDay, nil
}

func GetWarnings(day ...string) ([]Warning, error) {
	var (
		inputDay  string
		nameOfDay string
	)
	
	if len(day) > 0 {
		if day[0] != "today" && day[0] != "tommorow" && day[0] != "dayAfterTomorrow" {
			err := fmt.Errorf("invalid day name")
			return nil, err
		}
		inputDay = day[0]
	}

	if inputDay == "" {
		inputDay = "today"
	}

	if inputDay == "dayAfterTomorrow" {
		result, err := getDayAfterTomorrow(inputDay)
		if err != nil {
			return nil, err
		}
		nameOfDay = result
	}

	// Use nameOfDay in the rest of the code
	url := "https://www.met.ie/warnings/" + nameOfDay

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil, err
	}

	warnings := []Warning{}

	doc.Find("h2").Each(func(i int, s *goquery.Selection) {
		title := strings.TrimSpace(s.Text())

		if title == "Search" || title == "Connect" || title == "Download our apps" {
			return
		}

		warning := Warning{}

		warning.Title = title

		statusSpan := s.Next().Find("span.sr-only")
		if statusSpan != nil {

			warning.Description = strings.TrimSpace(strings.ReplaceAll(statusSpan.Text(), "<br />", "\n"))
		}

		strongElement := s.Next().Find("strong.sr-only")
		if strongElement != nil {

			warning.Description = strings.TrimSpace(strings.ReplaceAll(strongElement.Text(), "<br />", "\n"))
		}

		s.Next().Find("p").Each(func(j int, p *goquery.Selection) {

			text := strings.TrimSpace(strings.ReplaceAll(p.Text(), "<br />", "\n"))
			switch {
			case strings.Contains(text, "Valid"):
				warning.Valid = strings.TrimPrefix(text, "Valid: ")
			case strings.Contains(text, "Issued"):
				warning.Issued = strings.TrimPrefix(text, "Issued: ")
			}
		})

		warnings = append(warnings, warning)
	})

	// Testing:
	// var result string
	// for _, w := range warnings {
	// 	result += fmt.Sprintf("Title: %s\n", w.Title)
	// 	result += fmt.Sprintf("Description: %s\n", w.Description)
	// 	result += fmt.Sprintf("Valid: %s\n", w.Valid)
	// 	result += fmt.Sprintf("Issued: %s\n", w.Issued)
	// 	result += "-----------\n"
	// }

	return warnings, nil
}
