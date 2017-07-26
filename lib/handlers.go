package lib

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/suzuki-shunsuke/japanese-holiday-api/models"
	"github.com/suzuki-shunsuke/japanese-holiday-api/types"
)

func GetHolidays(c echo.Context) error {
	config, _ := GetConfig()
	db := GetConnection(config)
	var holidays_ []models.Holiday
	var holidays []types.Holiday
	db.Select("name, type, date").Find(&holidays_)
	for _, h := range holidays_ {
		holidays = append(holidays, h.ToType())
	}
	return c.JSON(http.StatusOK, holidays)
}

func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
