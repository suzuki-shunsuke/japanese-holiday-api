package main

import (
	"github.com/labstack/echo"
	"github.com/suzuki-shunsuke/japanese-holiday-api/controllers"
	"github.com/suzuki-shunsuke/japanese-holiday-api/lib"
	"strconv"
)

func main() {
	config, _ := lib.GetConfig()
	e := echo.New()
	e.GET("/health-check", controllers.CheckHealth)
	e.GET("/holidays", controllers.GetHolidays)
	e.Logger.Fatal(e.Start(":" + strconv.Itoa(config.App.Port)))
}
