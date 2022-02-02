package module1

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/digisan/echo-service/server/ws"
	"github.com/labstack/echo/v4"
)

// after implementing, register with path in 'api_reg.go'

// @Title Test
// @Description Get Random Number
// @Success 200 "OK"
// @Router /api/module1/test [get]
func Test(c echo.Context) error {
	num := rand.Intn(100)
	return c.String(http.StatusOK, fmt.Sprintf("GET test a random number: %d", num))
}

// @Title Test WS Message
// @Description Get Random Number
// @Success 200 "OK"
// @Router /api/module1/testmsg [get]
func TestSendMsg(c echo.Context) error {
	num := rand.Intn(100)
	ws.SendMsg("id", num)
	return c.String(http.StatusOK, fmt.Sprintf("GET test a random number: %d", num))
}
