package post

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// after implementing, register with path in 'api_reg.go'

func DrawImage(c echo.Context) error {
	return c.String(http.StatusOK, "test")
}
