package main

import (
	m1 "github.com/digisan/echo-service/server/api/module1"
	m2 "github.com/digisan/echo-service/server/api/module2"
	"github.com/digisan/echo-service/server/ws"
	"github.com/labstack/echo/v4"
)

// path: handler
var mGET = map[string]echo.HandlerFunc{

	// web socket for message
	"/ws/msg": ws.WSMsg,

	// module1 api
	"/api/module1/test":    m1.Test,
	"/api/module1/testmsg": m1.TestSendMsg,
}

var mPOST = map[string]echo.HandlerFunc{

	// module2 api
	"/api/module2/test": m2.TestPost,
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
