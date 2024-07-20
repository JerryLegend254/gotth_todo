package handlers

import (
	"github.com/JerryLegend254/gotth/views"
	"github.com/labstack/echo/v4"
)

func (h *Handler) Todos(c echo.Context) error {
	return views.Todos().Render(c.Request().Context(), c.Response().Writer)
}
