package main

import (
	"io/ioutil"
	"net/http"

	"interview/occurrence"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/words", fetchWordOccurrence)
	router.Run(":8080")
}

func fetchWordOccurrence(c *gin.Context) {
	file, err := c.FormFile("txt_file")
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, err.Error())
	} else {
		content, err := ioutil.ReadFile(file.Filename)
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, err.Error())
		}

		text := string(content)
		wordOccurrence := occurrence.GetOccurence(text)
		c.IndentedJSON(http.StatusOK, wordOccurrence)
	}

}
