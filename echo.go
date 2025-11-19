package main

import (
    "net/http"
    "net/http/httputil"
    "net/url"

    "github.com/labstack/echo/v4"
)

func main() {
    e := echo.New()

    // Define backend servers
    tmsTarget, _ := url.Parse("http://127.0.0.1:8081") // TMS backend
    gsTarget, _ := url.Parse("http://127.0.0.1:8082")  // GS backend

    // Create reverse proxies
    tmsProxy := httputil.NewSingleHostReverseProxy(tmsTarget)
    gsProxy := httputil.NewSingleHostReverseProxy(gsTarget)

    // Route /tms/* to TMS backend
    e.Any("/tms/*", func(c echo.Context) error {
        c.Request().Header.Set("X-Real-IP", c.RealIP())
        c.Request().Header.Set("X-Forwarded-For", c.RealIP())
        c.Request().Header.Set("X-Forwarded-Proto", c.Scheme())
        tmsProxy.ServeHTTP(c.Response(), c.Request())
        return nil
    })

    // Route /gs/* to GS backend
    e.Any("/gs/*", func(c echo.Context) error {
        c.Request().Header.Set("X-Real-IP", c.RealIP())
        c.Request().Header.Set("X-Forwarded-For", c.RealIP())
        c.Request().Header.Set("X-Forwarded-Proto", c.Scheme())
        gsProxy.ServeHTTP(c.Response(), c.Request())
        return nil
    })

    // Start Echo on port 80
    e.Logger.Fatal(e.Start(":80"))
}
