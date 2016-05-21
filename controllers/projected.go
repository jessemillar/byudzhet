package controllers

import (
	"net/http"

	"github.com/jessemillar/byudzhet/helpers"
	"github.com/labstack/echo"
)

func (cg *ControllerGroup) GetProjectedIncome(context echo.Context) error {
	token, err := helpers.ValidateJWT(context)
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	response, err := cg.Accessors.GetProjectedIncome(context, token.Email)
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func (cg *ControllerGroup) SetProjectedIncome(context echo.Context) error {
	token, err := helpers.ValidateJWT(context)
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	response, err := cg.Accessors.SetProjectedIncome(context, token.Email)
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func (cg *ControllerGroup) UpdateProjectedIncome(context echo.Context) error {
	token, err := helpers.ValidateJWT(context)
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	response, err := cg.Accessors.UpdateProjectedIncome(context, token.Email)
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}
