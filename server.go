package main

import (
	// Don't forget this first import or nothing will work
	_ "crypto/sha512"

	"fmt"

	"github.com/jessemillar/byudzhet/controllers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/fasthttp"
	"github.com/labstack/echo/middleware"
)

func main() {
	port := ":8000"
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())

	e.Get("/health", controllers.Health)
	e.Get("/callback", controllers.CallbackHandler)

	fmt.Printf("AV API is listening on %s\n", port)
	e.Run(fasthttp.New(port))
}
