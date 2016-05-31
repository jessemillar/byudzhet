package views

import (
	"net/http"
	"os"

	"github.com/jessemillar/byudzhet/helpers"
	"github.com/labstack/echo"
)

func Login(context echo.Context) error {
	_, err := helpers.ValidateJWT(context)
	if err != nil {
		return context.Render(http.StatusOK, "login", os.Getenv("AUTH0_CALLBACK"))
	}

	context.Redirect(http.StatusMovedPermanently, "/buckets")

	return context.String(http.StatusOK, "Done")
}

func Buckets(context echo.Context) error {
	helpers.CheckCookie(context)

	return context.Render(http.StatusOK, "buckets", os.Getenv("AUTH0_CALLBACK"))
}

func MakeBucket(context echo.Context) error {
	helpers.CheckCookie(context)

	return context.Render(http.StatusOK, "make-bucket", os.Getenv("AUTH0_CALLBACK"))
}

func Expenses(context echo.Context) error {
	helpers.CheckCookie(context)

	return context.Render(http.StatusOK, "expenses", os.Getenv("AUTH0_CALLBACK"))
}

func LogExpense(context echo.Context) error {
	helpers.CheckCookie(context)

	return context.Render(http.StatusOK, "log-expense", os.Getenv("AUTH0_CALLBACK"))
}

func Income(context echo.Context) error {
	helpers.CheckCookie(context)

	return context.Render(http.StatusOK, "income", os.Getenv("AUTH0_CALLBACK"))
}

func LogIncome(context echo.Context) error {
	helpers.CheckCookie(context)

	return context.Render(http.StatusOK, "log-income", os.Getenv("AUTH0_CALLBACK"))
}

func History(context echo.Context) error {
	helpers.CheckCookie(context)

	return context.Render(http.StatusOK, "history", os.Getenv("AUTH0_CALLBACK"))
}

func Settings(context echo.Context) error {
	helpers.CheckCookie(context)

	return context.Render(http.StatusOK, "settings", os.Getenv("AUTH0_CALLBACK"))
}
