package main

import (
	"github.com/suzuki-shunsuke/japanese-holiday-api/lib"
	"github.com/suzuki-shunsuke/japanese-holiday-api/models"
)

func main() {
	config, _ := lib.GetConfig()
	db := lib.GetConnection(config)
	db.AutoMigrate(&models.Holiday{})
	holidays := lib.ReadHolidayCsv("syukujitsu.csv")
	for _, holiday := range holidays {
		models.CreateHoliday(db, holiday)
	}
}
