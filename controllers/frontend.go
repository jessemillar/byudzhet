package controllers

import (
	"net/http"
	"os"

	"github.com/labstack/echo"
)

func (cg *ControllerGroup) Index(c echo.Context) error {
	return c.Render(http.StatusOK, "index", os.Getenv("AUTH0_CALLBACK"))
}

func (cg *ControllerGroup) Buckets(c echo.Context) error {
	return c.Render(http.StatusOK, "buckets", os.Getenv("AUTH0_CALLBACK"))
}
