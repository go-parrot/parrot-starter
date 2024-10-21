package server

import (
	"github.com/go-parrot/parrot-starter/internal/routers"
	"github.com/go-parrot/parrot/pkg/app"
	"github.com/go-parrot/parrot/pkg/transport/http"
)

// NewHTTPServer creates a HTTP server
func NewHTTPServer(c *app.Config) *http.Server {
	router := routers.NewRouter()

	srv := http.NewServer(
		http.WithAddress(c.HTTP.Addr),
		http.WithReadTimeout(c.HTTP.ReadTimeout),
		http.WithWriteTimeout(c.HTTP.WriteTimeout),
	)

	srv.Handler = router

	return srv
}
