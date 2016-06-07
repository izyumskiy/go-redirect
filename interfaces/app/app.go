package app


var appInstance *Application

// GetApplication возвращает экземпляр Application
func GetApplication() *Application {
    if appInstance == nil {
        appInstance = new(Application)

        // Init code
        appInstance.Config = loadConfig("config.ini")
        appInstance.Doc = make(AbstractPage)
        appInstance.routes = make(MapRoutes)
        // ...
    }

    return appInstance
}

// Routes устанавливает обработчики запросов в соответствии с URL'ами
func (app *Application) Routes(r MapRoutes) {
    app.routes = r
}

func (app *Application) Run() {
    r := mux.NewRouter()
    r.StrictSlash(true)

    for url, ctrl := range app.routes {
        r.HandleFunc(url, obs(ctrl))
    }

    http.Handle("/", r)
    listen := fmt.Sprintf("%s:%d", app.Config.Net.Listen_host, app.Config.Net.Listen_port)

    log.Println("Server is started on", listen)
    if err := http.ListenAndServe(listen, nil); err != nil {
        log.Println(err)
    }
}
