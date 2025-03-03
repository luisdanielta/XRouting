package api

import (
	"time"
	"xrouting/internal/auth"
	"xrouting/internal/ratelimiter"

	"github.com/labstack/echo/v4"
)

type Application struct {
	config        config
	authenticator auth.Authenticator
	rateLimiter   ratelimiter.Limiter
}

type config struct {
	addr        string
	db          dbConfig
	env         string
	apiURL      string
	frontendURL string
	auth        authConfig
	rateLimiter ratelimiter.Config
}

type authConfig struct {
	basic basicConfig
	token tokenConfig
}

type tokenConfig struct {
	secret string
	exp    time.Duration
	iss    string
}

type basicConfig struct {
	user string
	pass string
}

type dbConfig struct {
	addr         string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  string
}

func (app *Application) Mount(e *echo.Echo) {
	e.GET("/health", app.getHealth)

	/* v1 */
	// v1 := e.Group("/api/v1")
}
