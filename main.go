package main

import (
	"net/http"
	"github.com/redirects/interfaces/handlers"
	"log"
	"fmt"
	"github.com/gin-gonic/gin"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	log.Println("Server is preparing to start")
	
	router := gin.New()
	// Настройки приложения также из конфига
	// Здесь роутинг нужно брать из JSON конфига
	// для каждогу роута задавать имя контроллера
	// далее цикл по конфигу с роутами и подписка на все роуты
    router.GET("/click-:codes", handlers.Handler)
	/*
	/click-XXXXXX-YYYYYY
	/click?
	
	*/

    // By default it serves on :8080 unless a
    // PORT environment variable was defined.
    router.Run(":8080")
}