package main

import (
	"github.com/digisan/echo-service/server/get"
	"github.com/digisan/echo-service/server/post"
	"github.com/labstack/echo/v4"
)

var mGetAPI = map[string]echo.HandlerFunc{
	"/api/test": get.Test,
}

var mPostAPI = map[string]echo.HandlerFunc{
	"/api/test": post.DrawImage,
}

var mPutAPI = map[string]echo.HandlerFunc{
	"/api/test": nil,
}

var mDeleteAPI = map[string]echo.HandlerFunc{
	"/api/test": nil,
}

var mPatchAPI = map[string]echo.HandlerFunc{
	"/api/test": nil,
}

// ---------------------------------------- //

func hookPathHandler(e *echo.Echo) {
	methods := []string{"GET", "POST", "PUT", "DELETE", "PATCH"}

	var mRegAPIs = map[string]map[string]echo.HandlerFunc{
		"GET":    mGetAPI,
		"POST":   mPostAPI,
		"PUT":    mPutAPI,
		"DELETE": mDeleteAPI,
		"PATCH":  mPatchAPI,
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
