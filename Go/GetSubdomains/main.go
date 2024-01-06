package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func getSubdomains(targetDomain string) ([]string, error) {
	var subdomains []string

	httpSubdomains, err := performHttpRequest(targetDomain)
	if err != nil {
		return nil, err
	}

	subdomains = append(subdomains, httpSubdomains...)
	subdomains = removeDuplicates(subdomains)

	return subdomains, nil
}

func performHttpRequest(targetDomain string) ([]string, error) {
	var subdomains []string

	resp, err := http.Get("http://" + targetDomain)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	var traverse func(*html.Node)
	traverse = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					subdomain := extractSubDomain(attr.Val)
					if subdomain != "" {
						subdomains = append(subdomains, subdomain)
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			traverse(c)
		}
	}

	traverse(doc)
	return subdomains, nil
}

func extractSubDomain(urlString string) string {
	parsedURL, err := url.Parse(urlString)
	if err != nil {
		log.Println("Error:", err.Error())
		return ""
	}

	hostParts := strings.Split(parsedURL.Hostname(), ".")
	if len(hostParts) > 1 {
		return hostParts[0]
	}
	return ""
}

func removeDuplicates(slice []string) []string {
	encountered := make(map[string]bool)
	var result []string

	for _, value := range slice {
		if !encountered[value] {
			encountered[value] = true
			result = append(result, value)
		}
	}

	return result
}

func main() {
	fmt.Print("Enter the domain name:\n> ")
	var targetDomain string
	fmt.Scanln(&targetDomain)

	subdomains, err := getSubdomains(targetDomain)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println("Subdomains for", targetDomain, ":")
	for _, subdomain := range subdomains {
		fmt.Println(subdomain)
	}
}
