package handlers

import (
	"net/http"

	"github.com/jessemillar/byudzhet/helpers"
	"github.com/labstack/echo"
)

func (handlerGroup *HandlerGroup) GetUserByID(context echo.Context) error {
	_, err := helpers.ValidateJWT(context)
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	response, err := handlerGroup.Accessors.GetUserByID(context.Param("id"))
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func (handlerGroup *HandlerGroup) GetUserByEmail(context echo.Context) error {
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

	response, err := handlerGroup.Accessors.GetUserByEmail(email)
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func (handlerGroup *HandlerGroup) MakeUser(context echo.Context) error {
	token, err := helpers.ValidateJWT(context)
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	response, err := handlerGroup.Accessors.MakeUser(token.Email)
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}
