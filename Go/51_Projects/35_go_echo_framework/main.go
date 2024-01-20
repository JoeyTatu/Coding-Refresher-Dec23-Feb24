package main

import (
	"github.com/joeytatu/go-echo-framework/cmd/api/handlers"
	"github.com/labstack/echo"
)

func main() {
	echo := echo.New()
	echo.GET("/health-check", handlers.HealthCheckHandler)
	echo.GET("/posts", handlers.PostIndexHandler)
	echo.GET("/post/:id", handlers.PostSingleHandler)

	echo.Logger.Fatal(echo.Start(":1323"))
}