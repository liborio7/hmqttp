package iot

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/rs/zerolog/log"
	"hmqttp/httpcli"
	"hmqttp/mqttcli"
	"net/http"
)

type Handler struct {
	chi.Router
	mqtt *mqttcli.Client
}

func NewHandler(mqttClient *mqttcli.Client) *Handler {
	r := &Handler{chi.NewRouter(), mqttClient}

	r.Patch("/publish", r.publish)
	// r.Patch("/subscribe", r.subscribe)

	return r
}

func (h *Handler) publish(resp http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	msg := &Message{}
	if !h.mqtt.IsConnected() {
		_ = render.Render(resp, req, httpcli.ErrServiceUnavailable(mqttcli.ErrNotConnected))
		return
	}
	if err := render.Bind(req, msg); err != nil {
		_ = render.Render(resp, req, httpcli.ErrBadRequest(err))
		return
	}
	log.Ctx(ctx).Info().Msgf("publish message %+v", msg)
	h.mqtt.Publish(msg.Topic, 2, false, msg.Payload)

	render.Status(req, http.StatusOK)
	render.JSON(resp, req, msg)
}
