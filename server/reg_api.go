package main

import (
	"github.com/digisan/echo-service/server/api/module1/get"
	"github.com/digisan/echo-service/server/api/module2/post"
	"github.com/digisan/echo-service/server/ws"
	"github.com/labstack/echo/v4"
)

// path: handler
var mGET = map[string]echo.HandlerFunc{

	// web socket for message
	"/ws/msg": ws.WSMsg,

	// normal get api
	"/api/module1/test":    get.Test,
	"/api/module1/testmsg": get.TestSendMsg,
}

var mPOST = map[string]echo.HandlerFunc{
	"/api/test": post.DrawImage,
}

var mPUT = map[string]echo.HandlerFunc{
	"/api/test": nil,
}

var mDELETE = map[string]echo.HandlerFunc{
	"/api/test": nil,
}

var mPATCH = map[string]echo.HandlerFunc{
	"/api/test": nil,
}

// ---------------------------------------- //

func hookPathHandler(e *echo.Echo) {
	methods := []string{"GET", "POST", "PUT", "DELETE", "PATCH"}

	var mRegAPIs = map[string]map[string]echo.HandlerFunc{
		"GET":    mGET,
		"POST":   mPOST,
		"PUT":    mPUT,
		"DELETE": mDELETE,
		"PATCH":  mPATCH,
		// others...
	}

	type echoRoute func(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	mRegMethod := map[string]echoRoute{
		"GET":    e.GET,
		"POST":   e.POST,
		"PUT":    e.PUT,
		"DELETE": e.DELETE,
		"PATCH":  e.PATCH,
		// others...
	}

	for _, m := range methods {
		mAPI := mRegAPIs[m]
		method := mRegMethod[m]
		for path, handler := range mAPI {
			if handler == nil {
				continue
			}
			method(path, handler)
		}
	}
}
