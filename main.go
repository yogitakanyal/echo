package main

import (
	"net/http/httputil"
	"net/url"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Proxy for /gs
	e.Any("/gs/*", func(c echo.Context) error {
		proxy := httputil.NewSingleHostReverseProxy(&url.URL{Scheme: "http", Host: "194.0.0.111"})
		proxy.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	// Proxy for /tms
	e.Any("/tms/*", func(c echo.Context) error {
		proxy := httputil.NewSingleHostReverseProxy(&url.URL{Scheme: "http", Host: "195.0.0.222"})
		proxy.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	e.Logger.Fatal(e.Start(":8080"))

}
