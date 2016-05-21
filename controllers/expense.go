package controllers

import (
	"net/http"

	"github.com/jessemillar/byudzhet/helpers"
	"github.com/labstack/echo"
)

func (cg *ControllerGroup) LogExpense(context echo.Context) error {
	token, err := helpers.ValidateJWT(context)
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	response, err := cg.Accessors.LogExpense(context, token.Email)
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func (cg *ControllerGroup) GetExpense(context echo.Context) error {
	token, err := helpers.ValidateJWT(context)
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	response, err := cg.Accessors.GetExpense(context, token.Email)
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}
