package main

import (
	"github.com/labstack/echo"
	"github.com/suzuki-shunsuke/japanese-holiday-api/lib"
)

func main() {
	e := echo.New()
	e.GET("/", lib.Hello)
	e.GET("/holidays", lib.GetHolidays)
	e.Logger.Fatal(e.Start(":1323"))
}
