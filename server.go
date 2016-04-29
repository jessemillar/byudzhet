package main

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/fasthttp"
	"github.com/labstack/echo/middleware"
)

func main() {
	port := ":8000"
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())

	e.Get("/health", controllers.Health)

	fmt.Printf("AV API is listening on %s\n", port)
	e.Run(fasthttp.New(port))
}
