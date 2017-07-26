package lib

import (
	"net/http"

	"github.com/labstack/echo"
)

func GetHolidays(c echo.Context) error {
	return c.JSON(http.StatusOK, ReadHolidayCsv("syukujitsu.csv"))
}

func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
