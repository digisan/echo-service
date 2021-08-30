package get

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/net/websocket"
)

// after implementing, register with path in 'api_reg.go'

func Test(c echo.Context) error {
	return c.String(http.StatusOK, fmt.Sprintf("GET test a random number: %d", rand.Intn(100)))
}

func WSHello(c echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()

		// Read
		msg := ""
		err := websocket.Message.Receive(ws, &msg)
		if err != nil {
			c.Logger().Error(err)
			return
		}
		fmt.Printf("%s\n", msg)

		idx := 0
		for {
			// Write
			err = websocket.Message.Send(ws, fmt.Sprintf("Hello, Client! --- %d", idx))
			if err != nil {
				c.Logger().Error(err)
				return
			}
			time.Sleep(20 * time.Millisecond)
			idx++
		}
		
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}
