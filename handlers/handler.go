package handlers

import (
	"database/sql"

	"github.com/JerryLegend254/gotth/services/todos"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Handler struct {
}

func New() *Handler {
	return &Handler{}
}

var handler = New()

func RegisterRoutes(router *echo.Echo, db *sql.DB) *echo.Echo {

	todoStore := todos.NewStore(db)
	todosHandler := todos.NewHandler(todoStore)

	router.Static("/public", "public")

	router.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[ECHO] ${time_rfc3339_nano} |${status}|         ${latency_human} |             ${path} | ${method}\n",
	}))

	router.GET("/", handler.Home)

	router.GET("/todos", todosHandler.Todos)
	router.DELETE("/todos/:id", todosHandler.DeleteTodos)
	router.GET("/todos/:id/edit", todosHandler.EditTodoGet)
	router.PUT("/todos/:id/edit", todosHandler.EditTodoPut)
	router.POST("/todos", todosHandler.AddTodo)

	router.GET("/search", todosHandler.GetStoreBySearchQuery)

	return router
}
