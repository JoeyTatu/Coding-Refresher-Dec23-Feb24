package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

type Stock struct {
	Company string
	Price   string
	Change  string
}

func main() {
	ticker := []string{
		"MMM",  // 3M Company
		"INTC", // Intel Corporation
		"AXP",  // American Express Company
		"AAPL", // Apple Inc.
		"BA",   // The Boeing Company
		"CSCO", // Cisco Systems, Inc.
		"GS",   // The Goldman Sachs Group, Inc.
		"JPM",  // JPMorgan Chase & Co.
		"CRM",  // Salesforce.com, Inc.
		"VZ",   // Verizon Communications Inc.
		"DAX",  // DAX 30 (Germany)
		"CAC",  // CAC 40 (France)
		"JCI",  // Jakarta Composite Index (Indonesia)
		"IDX",  // IDX Composite (Indonesia)
		"SMI",  // Swiss Market Index (Switzerland)
	}

	stocks := []Stock{}

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting:", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	c.OnHTML("div#quote-header-info", func(e *colly.HTMLElement) {
		stock := Stock{}

		stock.Company = e.ChildText("h1")
		fmt.Println("Company:", stock.Company)

		stock.Price = e.ChildText("fin-streamer[data-field='regularMarketPrice']")

		stock.Change = e.ChildText("fin-streamer[data-field='regularMarketChangePercent']")

		stocks = append(stocks, stock)
	})

	c.Wait()

	for _, t := range ticker {
		c.Visit("https://finance.yahoo.com/quote/" + t + "/")
	}

	fmt.Println(stocks)

	file, err := os.Create("stocks.csv")
	if err != nil {
		log.Println("Unable to create file:", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	headers := []string{
		"Company",
		"Price",
		"Change %",
	}

	writer.Write(headers)
	for _, stock := range stocks {
		record := []string{
			stock.Company,
			stock.Price,
			stock.Change,
		}
		writer.Write(record)
	}
	defer writer.Flush()

}
