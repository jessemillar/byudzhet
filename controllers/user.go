package controllers

import (
	"net/http"

	"github.com/jessemillar/byudzhet/helpers"
	"github.com/labstack/echo"
)

func (cg *ControllerGroup) GetUserByID(context echo.Context) error {
	_, err := helpers.ValidateJWT(context)
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	response, err := cg.Accessors.GetUserByID(context.Param("id"))
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func (cg *ControllerGroup) GetUserByEmail(context echo.Context) error {
	token, err := helpers.ValidateJWT(context)
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	email := ""

	if len(context.Param("email")) > 0 {
		email = context.Param("email")
	} else {
		email = token.Email
	}

	response, err := cg.Accessors.GetUserByEmail(email)
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func (cg *ControllerGroup) MakeUser(context echo.Context) error {
	token, err := helpers.ValidateJWT(context)
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	response, err := cg.Accessors.MakeUser(token.Email)
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}
