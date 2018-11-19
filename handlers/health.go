package handlers

import (
    _ "github.com/facebookgo/inject"
    "github.com/go-chi/render"
    "net/http"
    "sync"
)

var healthHandler *HealthHandler
var healthHandlerOnce sync.Once

type HealthHandler struct {
}

func GetHealthHandler() *HealthHandler {
    healthHandlerOnce.Do(func() {
        healthHandler = &HealthHandler{}
    })
    return healthHandler
}

func (HealthHandler) Get(w http.ResponseWriter, r *http.Request) {
    render.JSON(w, r, "check")
}

