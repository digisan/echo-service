package ws

import (
	"context"
	"fmt"

	lk "github.com/digisan/logkit"
	"github.com/labstack/echo/v4"
	"golang.org/x/net/websocket"
)

// after implementing, register with path in 'api_reg.go'

var (
	mIdMsg      = map[string]chan interface{}{}
	mIdWSCancel = map[string]context.CancelFunc{}
)

func SendMsg(id string, msg interface{}) bool {
	if _, ok := mIdMsg[id]; !ok {
		return false
	}
	mIdMsg[id] <- msg
	return true
}

func BroadCast(msg interface{}) {
	for id := range mIdMsg {
		go func(id string) {
			SendMsg(id, msg)
		}(id)
	}
}

func CloseMsg(id string) bool {
	if _, ok := mIdWSCancel[id]; !ok {
		return false
	}
	mIdWSCancel[id]()
	return true
}

func CloseAllMsg() {
	for id := range mIdWSCancel {
		go func(id string) {
			CloseMsg(id)
		}(id)
	}
}

// Activate WS Msg by GET
func WSMsg(c echo.Context) error {

	id := c.Request().Header.Get("id")
	id = "id" // just for testing ***********************************

	// reg a new message channel
	mIdMsg[id] = make(chan interface{}, 1024)

	// reg message channel closing
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	mIdWSCancel[id] = cancel

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

		done := make(chan struct{})
		go func(ctx context.Context, done chan<- struct{}) {
			defer func() { done <- struct{}{} }()
			for {
				select {
				case msg := <-mIdMsg[id]:
					lk.WarnOnErr("%v", websocket.Message.Send(ws, fmt.Sprintf("From WS Server! --- %v", msg)))
				case <-ctx.Done():
					return
				}
			}
		}(ctx, done)
		<-done

	}).ServeHTTP(c.Response(), c.Request())

	return nil
}
