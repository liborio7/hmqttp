package clients

import (
    "github.com/eclipse/paho.mqtt.golang"
    "log"
    "sync"
    "time"
)

var mqttClient *MqttClient
var mqttClientOnce sync.Once

type MqttClient struct {
    sync.Once
    mqtt.Client
}

func GetMqttClient() *MqttClient {
    mqttClientOnce.Do(func() {
        mqttClient = &MqttClient{}
    })
    return mqttClient
}

func (mc *MqttClient) Connect() error {
    opts := mqtt.NewClientOptions().AddBroker("tcp://localhost:1883")
    opts.SetClientID("go_client")
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
        return err
    }
    return nil
}
