package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/redirects/interfaces/input"
)

func Handler(c *gin.Context) {
	inputInstance := input.Init(c)
	message := inputInstance.Uha_id + " is " + inputInstance.Link_id
	c.String(http.StatusOK, message)
}