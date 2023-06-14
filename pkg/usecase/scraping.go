package usecase

import (
	"log"
	"regexp"
	"strings"

	"web-scraper/pkg/entity"
	"web-scraper/pkg/repo"

	"github.com/go-playground/validator/v10"
	"github.com/gocolly/colly"
)

var validate = validator.New()

func OnScrapping(url string, c *colly.Collector) string {
	var data entity.DataBaseColumn

	// Set up a callback for when a visited page is parsed
	c.OnHTML("div#listbox a.xlink", func(e *colly.HTMLElement) {
		href := e.Attr("href")
		detail := e.ChildText("span.xinfo")

		category := strings.Split(detail, ",")[0]
		title := e.ChildText("span.xtitle")
		cover_img := e.ChildAttr("img.ximg", "src")

		// Create a new collector for crawling the href
		subC := c.Clone()
		subC.OnHTML("span.xcr", func(e *colly.HTMLElement) {
			author := e.ChildText("a")
			data.CreatedBy = author
		})
		subC.OnHTML("div.xdate time[itemprop=datePublished]", func(e *colly.HTMLElement) {
			datetime := e.Attr("datetime")
			createdDate := datetime[:10]
			data.CreatedDate = createdDate
		})

		re := regexp.MustCompile(`\((ตอบ\s*(\d+))\)`)
		match := re.FindStringSubmatch(detail)
		var replyTotalNumber string
		if len(match) > 2 {
			replyTotalNumber = match[2]
		} else {
			replyTotalNumber = "0"
		}
		subC.OnHTML("div#maincontent", func(e *colly.HTMLElement) {
			content := e.Text
			data.Content = content
		})

		subC.Visit(e.Request.AbsoluteURL(href))

		data = entity.DataBaseColumn{
			WebboardCategory: category,
			Topic:            title,
			CoverImage:       cover_img,
			ReplyTotalNumber: replyTotalNumber,
			CreatedBy:        data.CreatedBy,
			CreatedDate:      data.CreatedDate,
			Content:          data.Content,
		}

		// Validate the data struct
		err := validate.Struct(data)
		if err != nil {
			// Handle validation error
			log.Println("Validation error:", err)
			return
		}
		// Proceed with the validated data
		// log.Println("Data is valid")

		// log.Println(data)
		repo.InsertScrapeToMongoAtlas(data)
	})

	// Set up a callback for when the crawling is finished
	c.OnScraped(func(r *colly.Response) {
		log.Println("Crawling finished!")
	})

	return "Scrapping done !"
}
