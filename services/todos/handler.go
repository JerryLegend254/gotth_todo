package todos

import (
	"log"
	"net/http"
	"strconv"

	"github.com/JerryLegend254/gotth/components"
	"github.com/JerryLegend254/gotth/types"
	"github.com/JerryLegend254/gotth/views"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	store types.TodoStore
}

func NewHandler(s types.TodoStore) *Handler {
	return &Handler{store: s}
}

func (h *Handler) Todos(c echo.Context) error {
	isAddingTodo := c.QueryParam("isAddingTodo") == "true"
	todos, err := h.store.GetTodos()
	if err != nil {
		return err
	}

	return views.Todos(todos, isAddingTodo).Render(c.Request().Context(), c.Response().Writer)
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

func (h *Handler) GetStoreBySearchQuery(c echo.Context) error {

	query, err := c.FormParams()
	if err != nil {
		return err
	}
	q := query.Get("search-query")

	todos, err := h.store.GetStoreBySearchQuery(q)
	if err != nil {
		return err
	}

	return components.TodosList(todos).Render(c.Request().Context(), c.Response().Writer)

}

func (h *Handler) EditTodoGet(c echo.Context) error {

	id := c.Param("id")
	if id == "" {
		return c.String(http.StatusBadRequest, "id is required")
	}
	convertedId, err := strconv.Atoi(id)

	if err != nil {
		return c.String(http.StatusBadRequest, "id must be a number")
	}

	todo, err := h.store.GetTodoByID(convertedId)
	if err != nil {
		return err
	}

	return components.EditTodoForm(&todo).Render(c.Request().Context(), c.Response().Writer)
}

func (h *Handler) EditTodoPut(c echo.Context) error {
	v, err := c.FormParams()
	if err != nil {
		return err
	}
	id := c.Param("id")
	if id == "" {
		return c.String(http.StatusBadRequest, "id is required")
	}
	convertedId, err := strconv.Atoi(id)
	if err != nil {
		return c.String(http.StatusBadRequest, "id must be a number")
	}
	title := v.Get("title")
	completed := v.Get("completed")

	todo := types.Todo{
		ID:        convertedId,
		Title:     title,
		Completed: completed == "on",
	}
	_, err = h.store.EditTodo(todo)
	if err != nil {
		return err
	}
	//return c.Redirect(http.StatusFound, "/todos")
	return components.TodoCard(&todo).Render(c.Request().Context(), c.Response().Writer)

}

func (h *Handler) AddTodo(c echo.Context) error {
	body := new(types.Todo)

	if err := c.Bind(&body); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(&body); err != nil {

		//return echo.NewHTTPError(http.StatusBadRequest, errors.New("ensure all fields are filled"))
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	v, err := c.FormParams()
	if err != nil {
		return err
	}
	title := v.Get("title")
	completed := v.Get("completed") == "on"

	todo := types.Todo{
		Title:     title,
		Completed: completed,
	}

	createdTodo, err := h.store.CreateNewTodo(todo)
	if err != nil {
		log.Println(err)
		return err
	}

	c.Response().Header().Set("Content-Type", "text/html")

	return components.TodoCard(&createdTodo).Render(c.Request().Context(), c.Response().Writer)

}
