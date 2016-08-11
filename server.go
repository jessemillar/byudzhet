package main

import (
	// Don't forget this first import or nothing will work
	_ "crypto/sha512"
	"html/template"
	"log"
	"os"

	"github.com/jessemillar/byudzhet/accessors"
	"github.com/jessemillar/byudzhet/handlers"
	"github.com/jessemillar/byudzhet/helpers"
	"github.com/jessemillar/byudzhet/views"
	"github.com/jessemillar/health"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
)

func main() {
	database := os.Getenv("DATABASE_USERNAME") + ":" + os.Getenv("DATABASE_PASSWORD") + "@tcp(" + os.Getenv("DATABASE_HOST") + ":" + os.Getenv("DATABASE_PORT") + ")/" + os.Getenv("DATABASE_NAME")

	// Construct a new accessor group and connects it to the database
	accessorGroup := new(accessors.AccessorGroup)
	accessorGroup.Open("mysql", database)

	// Constructs a new controller group and gives it the accessor group
	handlerGroup := new(handlers.HandlerGroup)
	handlerGroup.Accessors = accessorGroup

	templateEngine := &helpers.Template{
		Templates: template.Must(template.ParseGlob("public/*/*.html")),
	}

	port := ":8000"
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.SetRenderer(templateEngine)

	e.Get("/callback", handlerGroup.CallbackHandler)

	e.Get("/health", health.Check)

	e.Get("/api/user/id/:id", handlerGroup.GetUserByID)
	e.Get("/api/user/email/:email", handlerGroup.GetUserByEmail)
	e.Get("/api/expense", handlerGroup.GetExpense)
	e.Get("/api/bucket", handlerGroup.GetBucket)
	e.Get("/api/bucket/:name", handlerGroup.GetBucketByName)
	e.Get("/api/income", handlerGroup.GetIncome)
	e.Get("/api/projected", handlerGroup.GetProjectedIncome)

	e.Post("/api/user", handlerGroup.MakeUser)
	e.Post("/api/expense", handlerGroup.LogExpense)
	e.Post("/api/bucket", handlerGroup.MakeBucket)
	e.Post("/api/income", handlerGroup.LogIncome)
	e.Post("/api/projected", handlerGroup.SetProjectedIncome)

	e.Put("/api/projected", handlerGroup.UpdateProjectedIncome)

	// Views
	e.Static("/*", "public")
	e.Get("/", views.Login)
	e.Get("/frontend", views.Frontend)

	log.Println("Byudzhet is listening on " + port)
	e.Run(standard.New(port))
}
