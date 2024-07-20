package main

import (
	"github.com/JerryLegend254/gotth/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	router := echo.New()
	router.Static("/public", "public")
	handler := handlers.New()
	router.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[ECHO] ${time_rfc3339_nano} |${status}|         ${latency_human} |             ${path} | ${method}\n",
	}))

	router.GET("/", handler.Home)
	router.GET("/todos", handler.Todos)
	router.Logger.Fatal(router.Start(":8080"))
}
