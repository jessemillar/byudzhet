package controllers

import (
	"net/http"

	"github.com/jessemillar/byudzhet/helpers"
	"github.com/labstack/echo"
)

func (cg *ControllerGroup) Health(c echo.Context) error {
	_, err := helpers.ValidateJWT(c)
	if err != nil {
		return c.String(http.StatusBadRequest, "Whoops")
	}

	return c.String(http.StatusOK, "Uh, we had a slight weapons malfunction, but uh... everything's perfectly all right now. We're fine. We're all fine here now, thank you. How are you?")
}
