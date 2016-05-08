package controllers

import (
	"net/http"
	"os"

	"github.com/labstack/echo"
)

func (cg *ControllerGroup) Frontend(c echo.Context) error {
	return c.Render(http.StatusOK, "frontend", os.Getenv("AUTH0_CALLBACK"))
}
