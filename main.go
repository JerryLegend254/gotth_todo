package main

import (
	"log"

	"github.com/JerryLegend254/gotth/configs"
	"github.com/JerryLegend254/gotth/db"
	"github.com/JerryLegend254/gotth/handlers"
	"github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

func main() {
	config := mysql.Config{
		User:                 configs.Envs.DBUser,
		Passwd:               configs.Envs.DBPassword,
		Net:                  "tcp",
		Addr:                 configs.Envs.DBAddress,
		DBName:               configs.Envs.DBName,
		AllowNativePasswords: true,
		ParseTime:            true,
	}
	d, err := db.NewSQLStorage(config)
	if err != nil {
		log.Fatal(err)
	}

	db.InitDB(d)

	router := echo.New()

	handlers.RegisterRoutes(router, d)

}
