package main

import (
	"github.com/labstack/echo"
	"github.com/suzuki-shunsuke/japanese-holiday-api/controllers"
	"github.com/suzuki-shunsuke/japanese-holiday-api/services"
	"strconv"
)

func main() {
	config, _ := services.GetConfig()
	e := echo.New()
	e.GET("/health-check", controllers.CheckHealth)
	e.GET("/holidays", controllers.GetHolidays)
	e.Logger.Fatal(e.Start(":" + strconv.Itoa(config.App.Port)))
}
