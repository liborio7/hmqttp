package main

import (
	"context"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"hmqttp/httpcli"
	"hmqttp/iot"
	"hmqttp/mqttcli"
	"os"
	"time"
)

func main() {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	output.PartsOrder = []string{"time", "level", "rid", "caller", "message"}
	output.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("%s >", i)
	}
	log.Logger = zerolog.New(output).With().Timestamp().Caller().Logger()

	ctx := context.Background()
	logger := log.With().Str("rid", "main").Logger()
	ctx = logger.WithContext(ctx)

	mqtt := mqttcli.NewClient(ctx, &mqttcli.ClientOpt{
		Address: "tcp://localhost:1883",
		Id:      "hmqttp_client",
	})
	_ = mqtt.Connect()
	ticker := time.NewTicker(5 * time.Second)
	go func() {
		for range ticker.C {
			if !mqtt.IsConnected() {
				_ = mqtt.Connect()
			}
		}
	}()

	http := httpcli.NewClient(ctx)
	http.Mount("/iot", iot.NewHandler(mqtt))
	http.ListenAndServe(":3000")
}
