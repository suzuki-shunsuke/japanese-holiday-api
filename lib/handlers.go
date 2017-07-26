package lib

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func GetHolidays(c echo.Context) error {
	bytes, err := ioutil.ReadFile("db.json")
	if err != nil {
		log.Fatal(err)
	}
	var db Db
	if err := json.Unmarshal(bytes, &db); err != nil {
		log.Fatal(err)
	}
	return c.JSON(http.StatusOK, db.Holidays)
}

func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
