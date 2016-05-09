package main

import (
	// Don't forget this first import or nothing will work
	_ "crypto/sha512"
	"html/template"
	"os"

	"fmt"

	"github.com/jessemillar/byudzhet/accessors"
	"github.com/jessemillar/byudzhet/controllers"
	"github.com/jessemillar/byudzhet/helpers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
)

func main() {
	database := os.Getenv("DATABASE_USERNAME") + ":" + os.Getenv("DATABASE_PASSWORD") + "@tcp(" + os.Getenv("DATABASE_HOST") + ":" + os.Getenv("DATABASE_PORT") + ")/" + os.Getenv("DATABASE_NAME")

	// Construct a new accessor group and connects it to the database
	ag := new(accessors.AccessorGroup)
	ag.Open("mysql", database)

	// Constructs a new controller group and gives it the accessor group
	cg := new(controllers.ControllerGroup)
	cg.Accessors = ag

	t := &helpers.Template{
		Templates: template.Must(template.ParseGlob("public/*/*.html")),
	}

	port := ":8000"
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.SetRenderer(t)

	e.Get("/health", cg.Health)
	e.Get("/callback", cg.CallbackHandler)
	e.Get("/user", cg.GetUser)

	e.Post("/expense", cg.LogExpense)

	e.Static("/*", "public")
	e.Get("/", helpers.Landing)
	e.Get("/buckets", helpers.Buckets)
	e.Get("/expenses", helpers.Expenses)
	e.Get("/log", helpers.LogExpense)

	fmt.Printf("Byudzhet is listening on %s\n", port)
	e.Run(standard.New(port))
}
