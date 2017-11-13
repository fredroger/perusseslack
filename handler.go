package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var lastVideoUpdateTime time.Time
var videoList []string

// SlackResponse represents a response to a slack command
type SlackResponse struct {
	ResponseType string `json:"response_type"`
	Text         string `json:"text"`
}

// Perusse gives access to an array of quotes
type Perusse struct {
	Quotes []Quote `json:"quotes"`
}

// Quote represents a quote
type Quote struct {
	Name string `json:"name"`
	Text string `json:"text"`
}

func getQuote(c *gin.Context) {

	// TODO validate request is comming from slack with a token etc. etc.
	quote := "Not quote found...ishhhhhh"

	file, err := ioutil.ReadFile("./perusse.json")
	if err == nil {
		var perusse Perusse
		json.Unmarshal(file, &perusse)

		// Pick a random number between 0 and maxResults
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)

		randomIndex := r1.Intn(len(perusse.Quotes))

		quote = perusse.Quotes[randomIndex].Text

	} else {
		fmt.Printf("File error: %v\n", err)
	}

	// Build the slack reponse and return through http
	var r SlackResponse
	r.ResponseType = "in_channel"
	r.Text = quote

	c.JSON(http.StatusOK, &r)
}

func getVideoURL(c *gin.Context) {

	// TODO validate request is comming from slack with a token etc. etc.

	// Update video list each 24 hours at most
	if (time.Since(lastVideoUpdateTime).Hours() >= 24) || (len(videoList) == 0) {

		lastVideoUpdateTime = time.Now()
		response := searchListByKeyword("snippet", 50, "album du peuple extrait", "video")
		videoList = videoList[:0] // empty list
		for _, item := range response.Items {
			videoList = append(videoList, item.Id.VideoId)
		}
	}

	// Pick a random number between 0 and maxResults
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	randomIndex := r1.Intn(len(videoList))

	// TODO get randomly a video from a file (json or whatever else...)
	videoURL := "https://www.youtube.com/watch?v=" + videoList[randomIndex]

	// Build the slack reponse and return through http
	var r SlackResponse
	r.ResponseType = "in_channel"
	r.Text = videoURL

	c.JSON(http.StatusOK, &r)
}
