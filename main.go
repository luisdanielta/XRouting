package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"
	"xrouting/internal/adapters/db"
	http "xrouting/internal/adapters/http"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/labstack/echo/v4"
	midd "github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

func main() {
	// Logger
	logger := zap.Must(zap.NewProduction()).Sugar()
	defer logger.Sync()

	ctx := context.Background()

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		logger.Errorw("failed to load config", "error", err)
	}

	e := echo.New()                       /* Framework Web */
	dbClient := db.NewDynamoDBClient(cfg) /* DynamoDB */
	logger.Infow("connected to dynamodb")

	e.Static("/", "public")

	/* Middleware */
	e.Use(http.DBMiddleware(dbClient))
	e.Use(midd.CORS())
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			logger.Infow("request", "method", c.Request().Method, "uri", c.Request().RequestURI)
			return next(c)
		}
	})

	/* Server */
	srv := http.NewEchoAdapter(":8000", e)
	logger.Infow("starting server", "addr", srv.Addr())

	/* Mount */
	srv.Mount(dbClient)
	logger.Infow("mounted routes")

	go func() {
		if err := srv.Start(); err != nil {
			logger.Errorw("failed to start server", "error", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("shutting down server")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Errorw("failed to shutdown server", "error", err)
	}
}
