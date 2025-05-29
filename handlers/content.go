package handlers

import (
	"net/http"

	"github.com/aminoxix/highliSense-server/helpers"
	"github.com/aminoxix/highliSense-server/interfaces"
	library "github.com/aminoxix/highliSense-server/libs"
	"github.com/gin-gonic/gin"
)

func PostContent(c *gin.Context) {
    var content interfaces.CONTENT
    if err := c.ShouldBindJSON(&content); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	if content.TYPE == "link" {
		// scrape, filter & format string to pass through Gemini
		content.TEXT = helpers.Scraper(content.LINK, &content)
	} else if content.TYPE == "extension" {
		// directly pass to Gemini with highlighted text
		content.TEXT = library.Gemini(content.CONTEXT, &content.HIGHLIGHTER)
	}
    c.JSON(http.StatusCreated, content)
}