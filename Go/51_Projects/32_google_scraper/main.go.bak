package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/joeytatu/google-scraper/vars"
)

var (
	googleDomains = vars.GoogleDomains
	userAgents    = vars.UserAgents
)

type SearchResult struct {
	Rank        int
	URL         string
	Title       string
	Description string
}

func randomUserAgent() string {
	source := rand.NewSource(time.Now().UnixNano())
	randomGenerator := rand.New(source)

	randNum := randomGenerator.Int() % len(userAgents)
	return userAgents[randNum]
}

func GoogleScrape(searchTerm, countryCode, languageCode string, proxyString interface{}, pages, count, backoff int) ([]SearchResult, error) {
	results := []SearchResult{}
	resultCounter := 0

	googlePages, err := buildGoogleURLs(searchTerm, countryCode, languageCode, pages, count)
	if err != nil {
		return nil, err
	}

	for _, page := range googlePages {
		res, err := scrapeClientResponse(page, proxyString)
		if err != nil {
			return nil, err
		}

		data, err := googleResultParsing(res, resultCounter)
		if err != nil {
			return nil, err
		}
		resultCounter += len(data)

		results = append(results, data...)

		time.Sleep(time.Duration(backoff) * time.Second)
	}
	return results, nil
}

func scrapeClientResponse(searchURL string, proxyString interface{}) (*http.Response, error) {
	client := getScrapeClient(proxyString)

	req, err := http.NewRequest("GET", searchURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", randomUserAgent())

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("scraper received a non-200 status code: %d", res.StatusCode)
	}

	return res, nil
}

func getScrapeClient(proxyString interface{}) *http.Client {
	switch v := proxyString.(type) {
	case string:
		proxyURL, err := url.Parse(v)
		if err != nil {
			log.Println("Error parsing proxy URL:", err)
			return &http.Client{}
		}

		return &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxyURL),
			},
		}
	default:
		return &http.Client{}
	}
}

func googleResultParsing(response *http.Response, rank int) ([]SearchResult, error) {
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil, err
	}

	results := []SearchResult{}
	selection := doc.Find("div.g")
	rank++

	for i := range selection.Nodes {
		item := selection.Eq(i)

		linkTag := item.Find("a")
		link, exists := linkTag.Attr("href")
		if !exists {
			continue // Skip results without a link
		}

		titleTag := item.Find("h3.r")
		title := titleTag.Text()

		descTag := item.Find("span.st")
		desc := descTag.Text()

		link = strings.Trim(link, " ")

		if link != "" && link != "#" && !strings.HasPrefix(link, "/") {
			result := SearchResult{
				rank,
				link,
				title,
				desc,
			}
			results = append(results, result)
			rank++
		}
	}

	return results, nil
}

func buildGoogleURLs(searchTerm, countryCode, languageCode string, pages, count int) ([]string, error) {
	toScrape := []string{}

	searchTerm = strings.Trim(searchTerm, " ")
	searchTerm = strings.Replace(searchTerm, " ", "+", -1)

	if googleBase, found := googleDomains[countryCode]; found {
		for i := 0; i < pages; i++ {
			start := i * count
			scrapeURL := fmt.Sprintf("%s%s&num=%d&hl=%s&start=%d&filter=%d", googleBase, searchTerm, count, languageCode, start, 0)
			toScrape = append(toScrape, scrapeURL)
		}
	} else {
		err := fmt.Errorf("country (%s) is not supported", countryCode)
		return nil, err
	}
	return toScrape, nil
}

func main() {
	// Gets the search term from the user
	fmt.Print("Enter search term:\n> ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	searchTerm := scanner.Text()

	// Loops and prints country code
	fmt.Println("\nAvailable country codes:")
	for code := range googleDomains {
		fmt.Println(code)
	}

	// Gets country code from the user
	fmt.Print("\nEnter country code:\n> ")
	scanner.Scan()
	countryCode := scanner.Text()

	result, err := GoogleScrape(searchTerm, countryCode, "en", nil, 1, 30, 10)
	if err != nil {
		log.Fatal("Error scraping Google:", err)
	} else {
		for _, res := range result {
			fmt.Println(res)
		}
	}
}
