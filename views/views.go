package views

import (
	"net/http"
	"os"

	"github.com/jessemillar/byudzhet/helpers"
	"github.com/labstack/echo"
)

func CheckCookie(c echo.Context) {
	_, err := helpers.ValidateJWT(c)
	if err != nil {
		c.Redirect(http.StatusMovedPermanently, "/")
	}
}

func Login(c echo.Context) error {
	_, err := helpers.ValidateJWT(c)
	if err != nil {
		return c.Render(http.StatusOK, "login", os.Getenv("AUTH0_CALLBACK"))
	}

	c.Redirect(http.StatusMovedPermanently, "/bucket")

	return c.String(http.StatusOK, "Done")
}

func Bucket(c echo.Context) error {
	CheckCookie(c)

	return c.Render(http.StatusOK, "bucket", os.Getenv("AUTH0_CALLBACK"))
}

func Expense(c echo.Context) error {
	CheckCookie(c)

	return c.Render(http.StatusOK, "expense", os.Getenv("AUTH0_CALLBACK"))
}

func MakeBucket(c echo.Context) error {
	CheckCookie(c)

	return c.Render(http.StatusOK, "make-bucket", os.Getenv("AUTH0_CALLBACK"))
}

func LogExpense(c echo.Context) error {
	CheckCookie(c)

	return c.Render(http.StatusOK, "log-expense", os.Getenv("AUTH0_CALLBACK"))
}
