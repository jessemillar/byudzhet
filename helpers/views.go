package helpers

import (
	"net/http"
	"os"

	"github.com/labstack/echo"
)

func CheckCookie(c echo.Context) {
	_, err := ValidateJWT(c)
	if err != nil {
		c.Redirect(http.StatusMovedPermanently, "/")
	}
}

func Landing(c echo.Context) error {
	_, err := ValidateJWT(c)
	if err != nil {
		return c.Render(http.StatusOK, "landing", os.Getenv("AUTH0_CALLBACK"))
	}

	c.Redirect(http.StatusMovedPermanently, "/buckets")

	return c.String(http.StatusOK, "Done")
}

func Buckets(c echo.Context) error {
	CheckCookie(c)

	return c.Render(http.StatusOK, "buckets", os.Getenv("AUTH0_CALLBACK"))
}

func Expenses(c echo.Context) error {
	CheckCookie(c)

	return c.Render(http.StatusOK, "expenses", os.Getenv("AUTH0_CALLBACK"))
}

func LogExpense(c echo.Context) error {
	CheckCookie(c)

	return c.Render(http.StatusOK, "log-expense", os.Getenv("AUTH0_CALLBACK"))
}
