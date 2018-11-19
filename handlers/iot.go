package handlers

import (
    "hmqttp/clients"
    "net/http"
    "sync"
)

var iotHandler *IotHandler
var iotHandlerOnce sync.Once

type IotHandler struct {
    mqttClient *clients.MqttClient
}

func GetIotHandler() *IotHandler {
    iotHandlerOnce.Do(func() {
        iotHandler = &IotHandler{
            mqttClient: clients.GetMqttClient(),
        }
    })
    return iotHandler
}

func (ih *IotHandler) Get(http.ResponseWriter, *http.Request) {
    ih.mqttClient.Publish("/go_topic", 2, false, "go_hello")
}
