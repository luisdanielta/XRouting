package api

import "github.com/labstack/echo/v4"

func GetHealth(c echo.Context) error {
	return c.JSON(200, map[string]string{"status": "ok"})
}
