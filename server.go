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
	"github.com/jessemillar/byudzhet/views"
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

	e.Get("/callback", cg.CallbackHandler)

	e.Get("/api/health", cg.Health)
	e.Get("/api/user", cg.GetUser)
	e.Get("/api/expense", cg.GetExpenses)

	e.Post("/api/expense", cg.LogExpense)
	e.Post("/api/bucket", cg.MakeBucket)

	// Views
	e.Static("/*", "public")
	e.Get("/", views.Login)
	e.Get("/buckets", views.Bucket)
	e.Get("/buckets/make", views.MakeBucket)
	e.Get("/expenses", views.Expense)
	e.Get("/expenses/log", views.LogExpense)

	fmt.Printf("Byudzhet is listening on %s\n", port)
	e.Run(standard.New(port))
}
