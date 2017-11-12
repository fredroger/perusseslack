package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getQuote(c *gin.Context) {
	quote := "Quand ma petite Scandinave a vu que j'la regardait avec un peu d'bave, elle m'a dit: « glabedichlabediglabedichlacken, glambadibediglabedichwetten, glabedichlabegrodibotchiklagen, dabodjekadebotchibotchikouine ». Ça voulait dire : « Voudrais-tu m'aider à visser ma chaise Ikéa »."
	c.String(http.StatusOK, quote)
}
