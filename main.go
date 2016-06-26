package main

import (
	"net/http"
	"github.com/go-redirect/app/routes"
	"log"
	"fmt"
)


func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}


func main() {
	log.Println("Server is preparing to start")

	server := routes.GetServer(":8080")
	server.InitRoutes()
	server.Run()
}