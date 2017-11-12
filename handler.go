package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SlackResponse represents a response to a slack command
type SlackResponse struct {
	ResponseType string `json:"response_type"`
	Text         string `json:"text"`
}

func getQuote(c *gin.Context) {

	// TODO validate request is comming from slack with a token etc. etc.

	// TODO get randomly a quote from a file (json or whatever else...)
	quote := "Quand ma petite Scandinave a vu que j'la regardait avec un peu d'bave, elle m'a dit: « glabedichlabediglabedichlacken, glambadibediglabedichwetten, glabedichlabegrodibotchiklagen, dabodjekadebotchibotchikouine ». Ça voulait dire : « Voudrais-tu m'aider à visser ma chaise Ikéa »."

	// Build the slack reponse and return through http
	var r SlackResponse
	r.ResponseType = "in_channel"
	r.Text = quote

	c.JSON(http.StatusOK, &r)
}

func getVideoURL(c *gin.Context) {

	// TODO validate request is comming from slack with a token etc. etc.

	// TODO get randomly a video from a file (json or whatever else...)
	videoURL := "https://www.youtube.com/watch?v=E1aK5w7WmNs"

	// Build the slack reponse and return through http
	var r SlackResponse
	r.ResponseType = "in_channel"
	r.Text = videoURL

	c.JSON(http.StatusOK, &r)
}
