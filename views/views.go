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

	context.Redirect(http.StatusMovedPermanently, "/frontend")

	return context.String(http.StatusOK, "Done")
}

func Frontend(context echo.Context) error {
	helpers.CheckCookie(context)

	return context.Render(http.StatusOK, "frontend", os.Getenv("AUTH0_CALLBACK"))
}
