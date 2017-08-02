package main

import (
	"github.com/suzuki-shunsuke/japanese-holiday-api/models"
	"github.com/suzuki-shunsuke/japanese-holiday-api/services"
)

func main() {
	config, _ := services.GetConfig()
	db, _ := models.GetConnection(config)
	db.AutoMigrate(&models.Holiday{})
	holidays, _ := services.ReadHolidayCsv("data/syukujitsu.csv")
	for _, holiday := range holidays {
		db.Create(&holiday)
	}
}
