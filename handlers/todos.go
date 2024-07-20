package handlers

import (
	"net/http"
	"strconv"

	"github.com/JerryLegend254/gotth/views"
	"github.com/labstack/echo/v4"
)

func (h *Handler) Todos(c echo.Context) error {
	todos, err := h.store.GetTodos()
	if err != nil {
		return err
	}

	return views.Todos(todos).Render(c.Request().Context(), c.Response().Writer)
}

func (h *Handler) DeleteTodos(c echo.Context) error {

	id := c.Param("id")
	if id == "" {
		return c.String(http.StatusBadRequest, "id is required")
	}
	convertedId, err := strconv.Atoi(id)

	if err != nil {
		return c.String(http.StatusBadRequest, "id must be a number")
	}

	if err := h.store.DeleteTodos(convertedId); err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}
