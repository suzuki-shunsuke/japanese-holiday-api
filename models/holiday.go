package models

import (
	"github.com/jinzhu/gorm"
	"github.com/suzuki-shunsuke/japanese-holiday-api/lib"
	"time"
)

type Holiday struct {
	ID   int
	Name string    `gorm:"not null;type:varchar(100)"`
	Date time.Time `gorm:"not null;type:date"`
	Type int       `gorm:"not null"`
}

func CreateHoliday(db *gorm.DB, holiday lib.Holiday) {
	layout := "2006-01-02"
	date, _ := time.Parse(layout, holiday.Date)
	item := Holiday{Name: holiday.Name, Type: holiday.Type, Date: date}
	db.Create(&item)
}
