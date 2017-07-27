package main

import (
	"github.com/labstack/echo"
	"github.com/suzuki-shunsuke/japanese-holiday-api/lib"
	"strconv"
)

func main() {
	config, _ := lib.GetConfig()
	e := echo.New()
	e.GET("/", lib.Hello)
	e.GET("/holidays", lib.GetHolidays)
	e.Logger.Fatal(e.Start(":" + strconv.Itoa(config.App.Port)))
}
