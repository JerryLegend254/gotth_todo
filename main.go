package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/JerryLegend254/gotth/configs"
	"github.com/JerryLegend254/gotth/db"
	"github.com/JerryLegend254/gotth/handlers"
	"github.com/JerryLegend254/gotth/utils"
	"github.com/alexedwards/scs/v2"
	"github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

type application struct {
	sessionManager *scs.SessionManager
}

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

	_ = &application{
		sessionManager: scs.New(),
	}

	router := echo.New()
	router.Validator = utils.NewValidator()

	svr := &http.Server{
		Addr:    fmt.Sprintf(":%s", configs.Envs.Port),
		Handler: handlers.RegisterRoutes(router, d),
	}

	log.Fatal(svr.ListenAndServe())

}
