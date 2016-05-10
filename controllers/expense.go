package controllers

import (
	"net/http"

	"github.com/jessemillar/byudzhet/helpers"
	"github.com/labstack/echo"
)

func (cg *ControllerGroup) LogExpense(c echo.Context) error {
	token, err := helpers.ValidateJWT(c)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	response, err := cg.Accessors.LogExpense(c, token.Email)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, response)
}

func (cg *ControllerGroup) GetExpense(c echo.Context) error {
	token, err := helpers.ValidateJWT(c)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	response, err := cg.Accessors.GetExpense(c, token.Email)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, response)
}
