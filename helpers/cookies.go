package helpers

import (
	"time"

	"github.com/labstack/echo"
)

func MakeCookie(c echo.Context, name string, value string) {
	cookie := new(echo.Cookie)
	cookie.SetName(name)
	cookie.SetValue(value)
	cookie.SetExpires(time.Now().Add(24 * 7 * time.Hour))
	c.SetCookie(cookie)
}
