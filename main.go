package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	url := os.Args[1]
	scrapedHTML := scrapeSite(url)
	status := getStatus(scrapedHTML)

	if status == "up" {
		fmt.Printf("✅  %s is up\n", url)
		return
	} else if status == "invalidDomain" {
		fmt.Printf("🔧  %s is not a valid domain\n", url)
		return
	} else {
		fmt.Printf("⚠️  %s is down\n", url)
		return
	}
}

func scrapeSite(url string) []string {
	fullURL := "https://isitup.org/" + url
	response, err := http.Get(fullURL)
	if err != nil {
		print(err)
	}

	defer response.Body.Close()
	if response.StatusCode != 200 {
		log.Fatalf("Status code error: %d %s", response.StatusCode, response.Status)
	}

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	results := make([]string, 1)

	doc.Find("#container").Each(func(i int, s *goquery.Selection) {
		para := s.Find("p").Text()
		results = append(results, para)
	})

	return results
}

func getStatus(scrapeResult []string) string {
	var status string
	for _, value := range scrapeResult {
		if strings.Contains(value, "is up") {
			status = "up"
		} else if strings.Contains(value, "need a valid domain to check") {
			status = "invalidDomain"
		} else {
			status = "down"
		}
	}
	return status
}
