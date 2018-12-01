package mqttcli

import (
	"context"
	"github.com/eclipse/paho.mqtt.golang"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"time"
)

type ClientOpt struct {
	Address string
	Id      string
}

type Client struct {
	mqtt.Client
	ctx context.Context
}

var ErrOnConnection = errors.New("can not establish mqtt client connection")
var ErrNotConnected = errors.New("mqtt client is not connected")

func NewClient(ctx context.Context, o *ClientOpt) *Client {
	opts := mqtt.NewClientOptions().AddBroker(o.Address)
	opts.SetClientID(o.Id)
	// opts.SetUsername("")
	// opts.SetPassword("")
	opts.SetAutoReconnect(true)
	opts.SetOnConnectHandler(func(client mqtt.Client) {
		log.Ctx(ctx).Info().Msg("mqtt connection is up")
	})
	opts.SetConnectionLostHandler(func(client mqtt.Client, err error) {
		client.Disconnect(5000)
		log.Ctx(ctx).Error().Msgf("mqtt connection lost: %s", err)
	})
	client := mqtt.NewClient(opts)

	return &Client{Client: client, ctx: ctx}
}

func (c *Client) Connect() error {
	token := c.Client.Connect()
	token.WaitTimeout(3 * time.Second)
	if err := token.Error(); err != nil {
		log.Ctx(c.ctx).Error().Msgf("error on mqtt client connection: %s", err)
		return ErrOnConnection
	}
	return nil
}
