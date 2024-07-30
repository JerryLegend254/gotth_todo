package handlers

import (
	"fmt"

	"github.com/JerryLegend254/gotth/configs"
	"github.com/JerryLegend254/gotth/types"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Handler struct {
	store types.TodoStore
}

func New(s types.TodoStore) *Handler {
	return &Handler{store: s}
}

func RegisterRoutes(router *echo.Echo, handler *Handler) {

	router.Static("/public", "public")

	router.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[ECHO] ${time_rfc3339_nano} |${status}|         ${latency_human} |             ${path} | ${method}\n",
	}))

	router.GET("/", handler.Home)

	router.GET("/todos", handler.Todos)
	router.DELETE("/todos/:id", handler.DeleteTodos)
	router.GET("/todos/:id/edit", handler.EditTodoGet)
	router.PUT("/todos/:id/edit", handler.EditTodoPut)
	router.POST("/todos", handler.AddTodo)

	router.GET("/search", handler.GetStoreBySearchQuery)
	router.Logger.Fatal(router.Start(fmt.Sprintf(":%s", configs.Envs.Port)))
}
