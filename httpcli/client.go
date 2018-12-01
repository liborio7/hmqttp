package httpcli

import (
	"context"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/rs/zerolog/log"
	"math/rand"
	"net/http"
	"time"
)

type Client struct {
	*chi.Mux
	ctx context.Context
}

func NewClient(ctx context.Context) *Client {
	r := chi.NewRouter()
	r.Use(requestId)
	r.Use(logging)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	return &Client{Mux: r, ctx: ctx}
}

func (c Client) ListenAndServe(port string) {
	if err := http.ListenAndServe(port, c); err != nil {
		log.Ctx(c.ctx).Panic().Msgf("error during router server startup: %s", err)
	}
}

func requestId(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "rid", fmt.Sprintf("%06d", rand.Intn(999999)))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		rid := ctx.Value("rid")
		logger := log.With().
			Str("rid", rid.(string)).
			Logger()

		t1 := time.Now()
		logger.Info().
			Str("method", r.Method).
			Str("uri", r.RequestURI).
			Msg("--- START ---")
		resp := &Response{ResponseWriter: w}
		next.ServeHTTP(resp, r.WithContext(logger.WithContext(ctx)))
		logger.Info().
			Str("status", fmt.Sprintf("%d", resp.Status())).
			Str("response_time", time.Since(t1).String()).
			Msg("--- END ---")
	})
}
