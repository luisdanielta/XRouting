package http

import (
	"xrouting/internal/adapters/db"

	"github.com/labstack/echo/v4"
)

func DBMiddleware(db *db.DynamoDBClient) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("db", db)
			return next(c)
		}
	}
}
