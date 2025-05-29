package helpers

import (
	"regexp"
	"strings"

	"github.com/aminoxix/highliSense-server/interfaces"
	library "github.com/aminoxix/highliSense-server/libs"
	"github.com/aminoxix/highliSense-server/utils"
	"github.com/go-rod/rod"
)


func Scraper(link string, content *interfaces.CONTENT) string {
	var contentArray []string

	// crawl through website URL to get the content
    page := rod.New().MustConnect().MustPage(link) // access through link

	// of paragraph tag
    paragraphs := page.MustWaitStable().MustElements("p")

	// loop over paragraphs elements
	for _, paragraph := range paragraphs {
		// to get it's inner-text
		text := paragraph.MustEval(`() => this.innerText`).String()
		// & remove the new line by replacing it with empty string
		str := regexp.MustCompile(`\n`)
		filteredString := str.ReplaceAllString(text, "")
		contentArray = append(contentArray, filteredString)
	}
	// more filtering to loop over contentArray & remove empty elements
	filteredContentArray := utils.FilterEmptyStrings(contentArray)

	// formatting the final filtered element's text by adding  new line
	content.CONTEXT = strings.Join(filteredContentArray, "\n ")

	// pass the formatted string to Gemini to process further...
	return library.Gemini(content.CONTEXT, &content.HIGHLIGHTER)
}
