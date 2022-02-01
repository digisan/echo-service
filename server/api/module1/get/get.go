package get

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
// -- @Accept json
// @Param name formData string true "Name"
// @Param age formData int true "Age"
// @Success 200 "获取信息成功"
// @Failure 400 "获取信息失败"
// @Router /api/test [get]
func Test(c echo.Context) error {
	num := rand.Intn(100)
	return c.String(http.StatusOK, fmt.Sprintf("GET test a random number: %d", num))
}

func TestSendMsg(c echo.Context) error {
	num := rand.Intn(100)
	ws.SendMsg("id", num)
	return c.String(http.StatusOK, fmt.Sprintf("GET test a random number: %d", num))
}
