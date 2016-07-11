package handlers

import (
	"net/http"

	"github.com/jessemillar/byudzhet/helpers"
	"github.com/labstack/echo"
)

func (handlerGroup *HandlerGroup) GetProjectedIncome(context echo.Context) error {
	token, err := helpers.ValidateJWT(context)
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	response, err := handlerGroup.Accessors.GetProjectedIncome(context, token.Email)
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func (handlerGroup *HandlerGroup) SetProjectedIncome(context echo.Context) error {
	token, err := helpers.ValidateJWT(context)
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	response, err := handlerGroup.Accessors.SetProjectedIncome(context, token.Email)
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func (handlerGroup *HandlerGroup) UpdateProjectedIncome(context echo.Context) error {
	token, err := helpers.ValidateJWT(context)
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	response, err := handlerGroup.Accessors.UpdateProjectedIncome(context, token.Email)
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}
