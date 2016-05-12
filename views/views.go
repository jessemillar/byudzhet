package views

import (
	"net/http"
	"os"

	"github.com/jessemillar/byudzhet/helpers"
	"github.com/labstack/echo"
)

func Login(c echo.Context) error {
	_, err := helpers.ValidateJWT(c)
	if err != nil {
		return c.Render(http.StatusOK, "login", os.Getenv("AUTH0_CALLBACK"))
	}

	// c.Redirect(http.StatusMovedPermanently, "/buckets")

	return c.String(http.StatusOK, "Done")
}

func Buckets(c echo.Context) error {
	// helpers.CheckCookie(c)

	return c.Render(http.StatusOK, "buckets", os.Getenv("AUTH0_CALLBACK"))
}

func MakeBucket(c echo.Context) error {
	helpers.CheckCookie(c)

	return c.Render(http.StatusOK, "make-bucket", os.Getenv("AUTH0_CALLBACK"))
}

func Expenses(c echo.Context) error {
	helpers.CheckCookie(c)

	return c.Render(http.StatusOK, "expenses", os.Getenv("AUTH0_CALLBACK"))
}

func LogExpense(c echo.Context) error {
	helpers.CheckCookie(c)

	return c.Render(http.StatusOK, "log-expense", os.Getenv("AUTH0_CALLBACK"))
}

func Income(c echo.Context) error {
	helpers.CheckCookie(c)

	return c.Render(http.StatusOK, "income", os.Getenv("AUTH0_CALLBACK"))
}

func LogIncome(c echo.Context) error {
	helpers.CheckCookie(c)

	return c.Render(http.StatusOK, "log-income", os.Getenv("AUTH0_CALLBACK"))
}

func Settings(c echo.Context) error {
	helpers.CheckCookie(c)

	return c.Render(http.StatusOK, "settings", os.Getenv("AUTH0_CALLBACK"))
}

func Share(c echo.Context) error {
	helpers.CheckCookie(c)

	return c.Render(http.StatusOK, "share", os.Getenv("AUTH0_CALLBACK"))
}
