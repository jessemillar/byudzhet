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
	router := echo.New()
	router.Pre(middleware.RemoveTrailingSlash())
	router.SetRenderer(t)

	router.Get("/callback", cg.CallbackHandler)

	router.Get("/health", health.Check)

	router.Get("/api/user/id/:id", cg.GetUserByID)
	router.Get("/api/user/email/:email", cg.GetUserByEmail)
	router.Get("/api/expense", cg.GetExpense)
	router.Get("/api/bucket", cg.GetBucket)
	router.Get("/api/bucket/:name", cg.GetBucketByName)
	router.Get("/api/income", cg.GetIncome)
	router.Get("/api/projected", cg.GetProjectedIncome)

	router.Post("/api/user", cg.MakeUser)
	router.Post("/api/expense", cg.LogExpense)
	router.Post("/api/bucket", cg.MakeBucket)
	router.Post("/api/income", cg.LogIncome)
	router.Post("/api/projected", cg.SetProjectedIncome)

	router.Put("/api/projected", cg.UpdateProjectedIncome)

	// Views
	router.Static("/*", "public")
	router.Get("/", views.Login)
	router.Get("/buckets", views.Buckets)
	router.Get("/buckets/make", views.MakeBucket)
	router.Get("/expenses", views.Expenses)
	router.Get("/expenses/log", views.LogExpense)
	router.Get("/income", views.Income)
	router.Get("/income/log", views.LogIncome)
	router.Get("/history", views.History)
	router.Get("/settings", views.Settings)

	fmt.Printf("Byudzhet is listening on %s\n", port)
	router.Run(standard.New(port))
}
