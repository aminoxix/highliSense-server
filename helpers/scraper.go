package helpers

import (
	"regexp"
	"strings"

	library "github.com/aminoxix/highliSense-server/libs"
	"github.com/aminoxix/highliSense-server/utils"
	"github.com/go-rod/rod"
)

func Scraper() string {
	var array []string
    page := rod.New().MustConnect().MustPage("https://github.com/github/gh-ost/pull/731") // access through link
    paragraphs := page.MustWaitStable().MustElements("p")
	for _, paragraph := range paragraphs {
		content:= paragraph.MustEval(`() => this.innerText`).String()
		str := regexp.MustCompile(`\n`)
		result := str.ReplaceAllString(content, "")
		array = append(array, result)
	}
	filteredArray := utils.FilterEmptyStrings(array)
	return library.Gemini(strings.Join(filteredArray, "\n "))
}
