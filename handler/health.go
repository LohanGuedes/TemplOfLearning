package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HandleHealthCheck(c echo.Context) error {
	status := map[string]interface{}{
		"name":    "tol_api",
		"version": "0.0.1",
		"status":  "health",
	}
	return c.JSON(http.StatusOK, status)
}
