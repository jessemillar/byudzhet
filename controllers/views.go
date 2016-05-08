package controllers

import (
	"net/http"
	"os"

	"github.com/labstack/echo"
)

func (cg *ControllerGroup) Landing(c echo.Context) error {
	return c.Render(http.StatusOK, "landing", os.Getenv("AUTH0_CALLBACK"))
}

func (cg *ControllerGroup) Buckets(c echo.Context) error {
	return c.Render(http.StatusOK, "buckets", os.Getenv("AUTH0_CALLBACK"))
}

func (cg *ControllerGroup) Expense(c echo.Context) error {
	return c.Render(http.StatusOK, "expense", os.Getenv("AUTH0_CALLBACK"))
}
