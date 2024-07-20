package main

import (
	"log"

	"github.com/JerryLegend254/gotth/db"
	"github.com/JerryLegend254/gotth/handlers"
	"github.com/JerryLegend254/gotth/services/todos"
	"github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config := mysql.Config{
		User:                 "root",
		Passwd:               "",
		Net:                  "tcp",
		Addr:                 "localhost:3306",
		DBName:               "gotth_todo",
		AllowNativePasswords: true,
		ParseTime:            true,
	}
	d, err := db.NewSQLStorage(config)
	if err != nil {
		log.Fatal(err)
	}

	db.InitDB(d)

	todosStore := todos.NewStore(d)

	router := echo.New()
	router.Static("/public", "public")
	handler := handlers.New(todosStore)
	router.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[ECHO] ${time_rfc3339_nano} |${status}|         ${latency_human} |             ${path} | ${method}\n",
	}))

	router.GET("/", handler.Home)
	router.GET("/todos", handler.Todos)
	router.DELETE("/todos/:id", handler.DeleteTodos)
	router.Logger.Fatal(router.Start(":8080"))
}
