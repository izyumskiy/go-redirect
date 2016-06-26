package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/go-redirect/interfaces/input"
)

func Handler(c *gin.Context) {
	inputInstance := input.Init(c)
	message := inputInstance.Uha_id + " is " + inputInstance.Link_id + " host:" + c.Request.Proto+ "; "+ " header:" + c.Request.UserAgent() + ";"

	c.String(http.StatusOK, message)
}

func redirectScenario(stp string, c *gin.Context) {

}

/*
1. if referer is empty -> go 301 redirect
2. if STEP1
2.1. if not WEBKIT -> META redirect
2.2. save referer to GET
2.3. if SSL -> JS META redirect
2.4. -> 301 redirect
3. if STEP2
3.1. if isset referer with HTTPS -> go 301 redirect
3.2. JS redirect
 */