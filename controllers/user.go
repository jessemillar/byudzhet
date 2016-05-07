package controllers

import (
	"net/http"

	"github.com/jessemillar/byudzhet/helpers"
	"github.com/labstack/echo"
)

func (cg *ControllerGroup) GetUser(c echo.Context) error {
	user, err := helpers.ValidateJWT(c)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	response, err := cg.Accessors.GetUser(user.Email)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, response)
}
