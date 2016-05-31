package helpers

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func MakeCookie(context echo.Context, name string, value string) {
	cookie := new(echo.Cookie)
	cookie.SetName(name)
	cookie.SetValue(value)
	cookie.SetExpires(time.Now().Add(24 * 365 * time.Hour)) // Make the cookie good for a year
	context.SetCookie(cookie)
}

func CheckCookie(context echo.Context) {
	_, err := ValidateJWT(context)
	if err != nil {
		context.Redirect(http.StatusMovedPermanently, "/")
	}
}
