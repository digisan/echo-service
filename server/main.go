package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/digisan/echo-service/server/docs" // once `swag init`, comment it out
	"github.com/digisan/echo-service/server/ws"
	gio "github.com/digisan/gotk/io"
	lk "github.com/digisan/logkit"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/postfinance/single"
	echoSwagger "github.com/swaggo/echo-swagger"
)

var (
	fHttp2 = false
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server.
// @termsOfService
// @contact.name API Support
// @contact.url
// @contact.email
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host 127.0.0.1:1323
// @BasePath
func main() {

	http2Ptr := flag.Bool("http2", false, "http2 mode?")
	flag.Parse()
	fHttp2 = *http2Ptr

	// Only One Instance
	const dir = "./tmp-locker"
	gio.MustCreateDir(dir)
	one, err := single.New("echo-service", single.WithLockPath(dir))
	lk.FailOnErr("%v", err)
	lk.FailOnErr("%v", one.Lock())
	defer func() {
		lk.FailOnErr("%v", one.Unlock())
		os.RemoveAll(dir)
		fmt.Println("Server Exited Successfully")
	}()

	// Start Service
	done := make(chan string)
	echoHost(done)
	fmt.Println(<-done)
}

func waitShutdown(e *echo.Echo) {
	go func() {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
		<-sig // got ctrl+c
		lk.Log("Got Ctrl+C")

		// other clean-up before shutting down
		ws.BroadCast("backend service shutting down...") // testing
		ws.CloseAllMsg()

		// shutdown echo
		lk.FailOnErr("%v", e.Shutdown(ctx)) // close echo at e.Shutdown
	}()
}

func echoHost(done chan<- string) {
	go func() {
		defer func() { done <- "Echo Shutdown Successfully" }()

		e := echo.New()
		defer e.Close()

		// Middleware
		e.Use(middleware.Logger())
		e.Use(middleware.Recover())
		e.Use(middleware.BodyLimit("2G"))
		// CORS
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowCredentials: true,
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
			AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		}))

		hookPathHandler(e) // hook each path-handler
		hookStatic(e)      // host static file/folder
		waitShutdown(e)    // waiting for shutdown

		// host swagger
		// http://localhost:1323/swagger/index.html
		e.GET("/swagger/*", echoSwagger.WrapHandler)

		// running...
		var err error
		if fHttp2 {
			err = e.StartTLS(":1323", "./cert/public.pem", "./cert/private.pem")
		} else {
			err = e.Start(":1323")
		}
		lk.FailOnErrWhen(err != http.ErrServerClosed, "%v", err)
	}()
}
