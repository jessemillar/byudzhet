package helpers

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func MakeCookie(c echo.Context, name string, value string) {
	cookie := new(echo.Cookie)
	cookie.SetName(name)
	cookie.SetValue(value)
	cookie.SetExpires(time.Now().Add(24 * 365 * time.Hour)) // Make the cookie good for a year
	c.SetCookie(cookie)
}

func CheckCookie(c echo.Context) {
	_, err := ValidateJWT(c)
	if err != nil {
		// c.Redirect(http.StatusMovedPermanently, "/")
	}
}
