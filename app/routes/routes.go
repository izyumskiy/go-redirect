package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redirect/interfaces/handlers"
)

type ServerApp struct {
	Router  *gin.Engine
	Port string
}

var server *ServerApp

func GetServer(port string) *ServerApp {
	//if server == nil {
		router := gin.New()
		server := new(ServerApp)
		server.Router = router
		server.Port = port
	//}
	return server;
}

func (server *ServerApp) InitRoutes() {
	server.Router.GET("/click-:codes", handlers.Handler)
	//server.Router.GET("/click-:codes", handlers.Handler)
}

func (server *ServerApp) Run() {
	server.Router.Run(server.Port)
}

