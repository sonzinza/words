package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WordOccurence struct {
	Word       string `json:"word"`
	Occurrence int    `json:"occurrence"`
}

func main() {
	router := setUpRouter()
	router.Run(":8080")
}

func setUpRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/words", fetchWordOccurrence)
	return r
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

		postBody, _ := json.Marshal(map[string]string{
			"text": string(content),
		})
		// TODO: read file by chunk and process each chunk to optimize
		responseBody := bytes.NewBuffer(postBody)
		resp, err := http.Post("http://localhost:8081/occurrence", "application/json", responseBody)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		var wordsChunk []WordOccurence
		if err := json.Unmarshal(body, &wordsChunk); err != nil {
			log.Fatal(err)
		}
		c.IndentedJSON(http.StatusOK, wordsChunk)
	}
}
