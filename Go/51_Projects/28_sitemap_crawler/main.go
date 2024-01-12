package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var userAgents = []string{
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Safari/14.0.3",
	"Mozilla/5.0 (Linux; Android 11; Pixel 4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Mobile Safari/537.36",
	"Mozilla/5.0 (Windows NT 6.1; WOW64; Trident/7.0; AS; rv:11.0) like Gecko",
	"Mozilla/5.0 (X11; Linux x86_64; rv:89.0) Gecko/20100101 Firefox/89.0",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36 Edg/91.0.864.48",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.1 Safari/605.1.15",
	"Mozilla/5.0 (Linux; Android 11; SM-G991B) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Mobile Safari/537.36",
	"Mozilla/5.0 (Windows NT 6.1; WOW64; rv:89.0) Gecko/20100101 Firefox/89.0",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36 OPR/76.0.4017.123",
	"Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:89.0) Gecko/20100101 Firefox/89.0",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36 Edg/91.0.864.48",
}

type SEOData struct {
	URL             string
	Title           string
	H1Tag           string
	MetaDescription string
	StatusCode      int
}

type DefaultParser struct{}

type Parser interface {
	getSEOData(response *http.Response) (SEOData, error)
}

func randomUserAgent() string {
	source := rand.NewSource(time.Now().UnixNano())
	randomGenerator := rand.New(source)

	randNum := randomGenerator.Int() % len(userAgents)
	return userAgents[randNum]
}

func isSiteMap(urls []string) ([]string, []string) {
	siteMapFiles := []string{}
	pages := []string{}

	for _, page := range urls {
		foundSiteMap := strings.Contains(page, "xml")
		if foundSiteMap == true {
			fmt.Println("found SiteMap", page)
			siteMapFiles = append(siteMapFiles, page)
		} else {
			pages = append(pages, page)
		}
	}
	return siteMapFiles, pages
}

func scrapeSiteMap(url string, parser Parser, concurrency int) []SEOData {
	results := extractSiteMapURLs(url)
	res := scrapeURLs(results, parser, concurrency)
	return res
}

func extractSiteMapURLs(startURL string) []string {
	worklist := make(chan []string)
	toCrawl := []string{}

	var n int
	n++

	go func() { worklist <- []string{startURL} }()

	for ; n > 0; n++ {
		list := <-worklist
		for _, link := range list {
			n++
			go func(link string) {
				response, err := makeRequest(link)
				if err != nil {
					log.Printf("Error retrieving URL: %s", err)
					return
				}

				urls, err := extractURLs(response)
				if err != nil {
					log.Printf("Error extracting document from response: %s", err)
					return
				}

				siteMapFiles, pages := isSiteMap(urls)
				if siteMapFiles != nil {
					worklist <- siteMapFiles
				}

				for _, page := range pages {
					toCrawl = append(toCrawl, page)
				}
			}(link)
		}
	}
	return toCrawl
}

func makeRequest(url string) (*http.Response, error) {
	client := http.Client{
		Timeout: 10 * time.Second,
	}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("User-Agent", randomUserAgent())

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func scrapeURLs(urls []string, parser Parser, concurrency int) []SEOData {
	tokens := make(chan struct{}, concurrency)
	var n int
	worklist := make(chan []string)
	results := []SEOData{}

	go func() { worklist <- urls }()
	for ; n > 0; n-- {
		list := <-worklist
		for _, url := range list {
			if url != "" {
				n++
				go func(url string, token chan struct{}) {
					log.Printf("Requesting URL: %s", url)
					result, err := scrapePage(url, parser)
					if err != nil {
						log.Printf("Error, URL: %s", url)
					} else {
						results = append(results, result)
					}
					worklist <- []string{}
				}(url, tokens)
			}
		}
	}
	return results
}

func extractURLs(response *http.Response) ([]string, error) {
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil, err
	}

	results := []string{}
	sel := document.Find("loc")
	for i := range sel.Nodes {
		location := sel.Eq(i)
		result := location.Text()
		results = append(results, result)
	}

	return results, nil
}

func scrapePage(url string, parser Parser) (SEOData, error) {
	res, err := crawlPage(url)
	if err != nil {
		return SEOData{}, err
	}

	data, err := parser.getSEOData(res)
	if err != nil {
		return SEOData{}, err
	}

	return data, nil
}

func crawlPage(url string) (*http.Response, error) {
	response, err := makeRequest(url)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (defaultParser DefaultParser) getSEOData(response *http.Response) (SEOData, error) {
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return SEOData{}, err
	}

	result := SEOData{}

	result.URL = response.Request.URL.String()
	result.Title = document.Find("title").First().Text()
	result.H1Tag = document.Find("h1").First().Text()
	result.MetaDescription, _ = document.Find("meta[name^=description]").Attr("content")
	result.StatusCode = response.StatusCode

	return result, nil
}

func main() {
	parser := DefaultParser{}

	fmt.Print("Enter the domain (e.g., 'quicksprout.com'):\n> ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	url := scanner.Text()

	results := scrapeSiteMap("https://www."+url+"/sitemap.xml", parser, 10)
	for _, result := range results {
		fmt.Println(result)
	}
}
