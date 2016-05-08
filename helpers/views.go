package helpers

import (
	"net/http"
	"os"

	"github.com/labstack/echo"
)

func Landing(c echo.Context) error {
	return c.Render(http.StatusOK, "landing", os.Getenv("AUTH0_CALLBACK"))
}

func Buckets(c echo.Context) error {
	return c.Render(http.StatusOK, "buckets", os.Getenv("AUTH0_CALLBACK"))
}

func Expenses(c echo.Context) error {
	return c.Render(http.StatusOK, "expenses", os.Getenv("AUTH0_CALLBACK"))
}

func LogExpense(c echo.Context) error {
	return c.Render(http.StatusOK, "log-expense", os.Getenv("AUTH0_CALLBACK"))
}
