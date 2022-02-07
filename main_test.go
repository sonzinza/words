package main

import (
	"bytes"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"gopkg.in/h2non/gock.v1"
)

var (
	response []WordOccurence = []WordOccurence{
		{
			Word:       "sed",
			Occurrence: 249,
		},
		{
			Word:       "in",
			Occurrence: 229,
		},
		{
			Word:       "ut",
			Occurrence: 214,
		},
		{
			Word:       "et",
			Occurrence: 204,
		},
		{
			Word:       "ac",
			Occurrence: 201,
		},
		{
			Word:       "non",
			Occurrence: 197,
		},
		{
			Word:       "eget",
			Occurrence: 191,
		},
		{
			Word:       "quis",
			Occurrence: 171,
		},
		{
			Word:       "id",
			Occurrence: 170,
		},
		{
			Word:       "sit",
			Occurrence: 169,
		},
	}
)

func TestMain(t *testing.T) {
	var r *gin.Engine
	t.Run("Test call successs", func(t *testing.T) {
		r = setUpRouter()
		w := httptest.NewRecorder()
		gock.New("http://localhost:8081").
			Post("/occurrence").
			Reply(200).
			JSON(response)
		req, _ := newfileUploadRequest("/words", "txt_file", "GoLang_Test.txt")

		r.ServeHTTP(w, req)
		assert.Equal(t, 200, w.Code)
	})
}

func newfileUploadRequest(uri string, paramName, path string) (*http.Request, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	fi, err := file.Stat()
	if err != nil {
		return nil, err
	}
	file.Close()

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, fi.Name())
	if err != nil {
		return nil, err
	}
	part.Write(fileContents)
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, _ := http.NewRequest("POST", uri, body)

	req.Header.Add("Content-Type", writer.FormDataContentType())
	return req, nil
}
