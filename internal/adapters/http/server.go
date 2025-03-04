package http

import (
	"context"
	"net/http"
	"time"
	"xrouting/cmd/api"
	"xrouting/internal/adapters/db"
	"xrouting/internal/domain/repositories"
	"xrouting/internal/ports"

	"github.com/labstack/echo/v4"
)

type Server interface {
	Start() error
	Shutdown(ctx context.Context) error
	Mount(db *db.DynamoDBClient)
	Addr() string
}

type EchoAdapter struct {
	echo   *echo.Echo
	server *http.Server
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
	}
}

func (ea *EchoAdapter) Start() error {

	return ea.server.ListenAndServe()
}

func (ea *EchoAdapter) Shutdown(ctx context.Context) error {
	return ea.server.Shutdown(ctx)
}

func (ea *EchoAdapter) Addr() string {
	return ea.server.Addr
}

func (ea *EchoAdapter) Mount(db *db.DynamoDBClient) {

	/* domain */
	userRepo := repositories.NewUserRepository(db)
	userService := ports.NewUserService(userRepo)
	userHandler := api.NewUserHandler(userService)

	commentRepo := repositories.NewCommentRepository(db)
	commentService := ports.NewCommentService(commentRepo)
	commentHandler := api.NewCommentHandler(commentService)

	componentRepo := repositories.NewComponentRepository(db)
	componentService := ports.NewComponentService(componentRepo)
	componentHandler := api.NewComponentHandler(componentService)

	e := ea.echo
	e.GET("/health", api.GetHealth)

	v1 := e.Group("/api/v1")

	v1.POST("/user", userHandler.RegisterUser)
	v1.GET("/user/:id", userHandler.FindUser)
	v1.DELETE("/user/:id", userHandler.RemoveUser)

	v1.POST("/comment", commentHandler.RegisterComment)
	v1.GET("/comment/:id", commentHandler.FindComment)
	v1.DELETE("/comment/:id", commentHandler.RemoveComment)

	v1.GET("/components", componentHandler.Components)
	v1.POST("/component", componentHandler.RegisterComponent)
	v1.GET("/component/:id", componentHandler.FindComponent)
	v1.DELETE("/component/:id", componentHandler.RemoveComponent)
}
