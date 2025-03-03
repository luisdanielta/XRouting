package http

import (
	"context"
	"net/http"
	"time"
	"xrouting/cmd/api"

	"github.com/labstack/echo/v4"
)

type Server interface {
	Start() error
	Shutdown(ctx context.Context) error
}

type EchoAdapter struct {
	echo   *echo.Echo
	server *http.Server
	app    *api.Application
}

func NewEchoAdapter(addr string, echoInstance *echo.Echo) Server {
	return &EchoAdapter{
		echo: echoInstance,
		server: &http.Server{
			Addr:         addr,
			Handler:      echoInstance,
			ReadTimeout:  30 * time.Second,
			WriteTimeout: 30 * time.Second,
			IdleTimeout:  60 * time.Second,
		},
		app: &api.Application{},
	}
}

func (ea *EchoAdapter) Start() error {
	ea.app.Mount(ea.echo)
	return ea.server.ListenAndServe()
}

func (ea *EchoAdapter) Shutdown(ctx context.Context) error {
	return ea.server.Shutdown(ctx)
}

func (ea *EchoAdapter) Addr() string {
	return ea.server.Addr
}
