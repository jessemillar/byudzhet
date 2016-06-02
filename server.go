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
	"github.com/jessemillar/health"
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

	e.Get("/health", health.Check)

	e.Get("/api/user/id/:id", cg.GetUserByID)
	e.Get("/api/user/email/:email", cg.GetUserByEmail)
	e.Get("/api/expense", cg.GetExpense)
	e.Get("/api/bucket", cg.GetBucket)
	e.Get("/api/bucket/:name", cg.GetBucketByName)
	e.Get("/api/income", cg.GetIncome)
	e.Get("/api/projected", cg.GetProjectedIncome)

	e.Post("/api/user", cg.MakeUser)
	e.Post("/api/expense", cg.LogExpense)
	e.Post("/api/bucket", cg.MakeBucket)
	e.Post("/api/income", cg.LogIncome)
	e.Post("/api/projected", cg.SetProjectedIncome)

	e.Put("/api/projected", cg.UpdateProjectedIncome)

	// Views
	e.Static("/*", "public")
	e.Get("/", views.Login)
	e.Get("/frontend", views.Frontend)

	fmt.Printf("Byudzhet is listening on %s\n", port)
	e.Run(standard.New(port))
}
