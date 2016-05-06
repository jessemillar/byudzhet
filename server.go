package main

import (
	// Don't forget this first import or nothing will work
	_ "crypto/sha512"
	"os"

	"fmt"

	"github.com/jessemillar/byudzhet/accessors"
	"github.com/jessemillar/byudzhet/controllers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/fasthttp"
	"github.com/labstack/echo/middleware"
)

func main() {
	database := os.Getenv("DATABASE_USER") + ":" + os.Getenv("DATABASE_PASSWORD") + "@tcp(" + os.Getenv("DATABASE_HOST") + ":" + os.Getenv("DATABASE_PORT") + ")/" + os.Getenv("DATABASE_NAME")

	// Construct a new accessor group and connects it to the database
	ag := new(accessors.AccessorGroup)
	ag.Open("mysql", database)

	// Constructs a new controller group and gives it the accessor group
	cg := new(controllers.ControllerGroup)
	cg.Accessors = ag

	port := ":8000"
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())

	e.Get("/health", cg.Health)
	e.Get("/callback", cg.CallbackHandler)

	e.Static("/*", "content")

	fmt.Printf("Byudzhet is listening on %s\n", port)
	e.Run(fasthttp.New(port))
}
