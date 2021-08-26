package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	gio "github.com/digisan/gotk/io"
	lk "github.com/digisan/logkit"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/postfinance/single"
)

var (
	failOnErr = lk.FailOnErr
)

func main() {
	// Only One Instance
	const dir = "./tmp-locker"
	gio.MustCreateDir(dir)
	one, err := single.New("echo-service", single.WithLockPath(dir))
	failOnErr("%v", err)
	failOnErr("%v", one.Lock())
	defer func() {
		failOnErr("%v", one.Unlock())
		os.RemoveAll(dir)
		fmt.Println("Program Exited")
	}()

	// Start Service
	done := make(chan string)
	hostHTTP(done)
	fmt.Println(<-done)
}

func waitShutdown(e *echo.Echo) {
	go func() {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
		<-sig // got ctrl+c
		fmt.Println("\nGot Ctrl(c)")
		failOnErr("%v", e.Shutdown(ctx)) // close http at e.Shutdown
	}()
}

func hostHTTP(done chan<- string) {
	go func() {
		defer func() { done <- "HTTP Shutdown Successfully" }()

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

		if err := e.Start(fmt.Sprintf(":%d", 1545)); err != nil {
			fmt.Println(err)
		}
	}()
}
