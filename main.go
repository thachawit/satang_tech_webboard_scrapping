package main

import (
	"log"
	"web-scraper/pkg/config"
	"web-scraper/pkg/usecase"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()
	// Set the entry point for crawling
	scrape_url := "https://board.postjung.com/index.php?page=0"

	//connect database
	config.ConnectDB()

	//scrapping starts here
	usecase.OnScrapping(scrape_url, c)

	err := c.Visit(scrape_url)
	if err != nil {
		log.Fatal(err)
	}

}
