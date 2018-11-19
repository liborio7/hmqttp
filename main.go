package main

import (
    "github.com/go-chi/chi"
    "github.com/go-chi/chi/middleware"
    "github.com/go-chi/render"
    "hmqttp/clients"
    "hmqttp/handlers"
    "log"
    "net/http"
)


func main() {
    _ = clients.GetMqttClient().Connect()

    r := chi.NewRouter()
    r.Use(middleware.RequestID)
    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)
    r.Use(middleware.URLFormat)
    r.Use(render.SetContentType(render.ContentTypeJSON))

    r.Get("/health", handlers.GetHealthHandler().Get)
    r.Get("/iot", handlers.GetIotHandler().Get)

    if err := http.ListenAndServe(":3000", r); err != nil {
        log.Println("error during router server startup:", err)
        panic("application stopped")
    }
}

/*
func main() {
    // mqtt client
    opts := mqtt.NewClientOptions().AddBroker("tcp://localhost:1883")
    opts.SetClientID("go_client_id")
    //opts.SetUsername("")
    //opts.SetPassword("")
    opts.SetAutoReconnect(true)
    opts.SetOnConnectHandler(func(client mqtt.Client) {
        log.Println("mqtt connection is up")
    })
    opts.SetConnectionLostHandler(func(client mqtt.Client, err error) {
        log.Println("mqtt connection lost", err)
    })

    client := mqtt.NewClient(opts)
    token := client.Connect()
    token.WaitTimeout(3 * time.Second)
    if err := token.Error(); err != nil {
        log.Println("error during mqtt client startup:", err)
        panic("application stopped")
    }

    // http server
    r := chi.NewRouter()
    r.Use(middleware.RequestID)
    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)
    r.Use(middleware.URLFormat)
    r.Use(render.SetContentType(render.ContentTypeJSON))

    r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
        render.JSON(w, r, "check")
    })
    r.Get("/iot", func(w http.ResponseWriter, r *http.Request) {
        client.Publish("/go_topic", 2, false, "go_hello")
        render.NoContent(w, r)
    })

    // start
    if err := http.ListenAndServe(":3000", r); err != nil {
        log.Println("error during router server startup:", err)
        panic("application stopped")
    }
}
*/
