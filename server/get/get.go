package get

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// after implementing, register with path in 'api_reg.go'

func Test(c echo.Context) error {
	return c.String(http.StatusOK, "test")
}
