package models

import (
	"github.com/jinzhu/gorm"
	"github.com/suzuki-shunsuke/japanese-holiday-api/types"
	"time"
)

type Holiday struct {
	ID   int       `json:"id"`
	Name string    `gorm:"not null;type:varchar(100)" json:"name"`
	Date time.Time `gorm:"not null;type:date" json:"date"`
	Type int       `gorm:"not null" json:"type"`
}

func CreateHoliday(db *gorm.DB, holiday types.Holiday) {
	date, _ := time.Parse("2006-01-02", holiday.Date)
	item := Holiday{Name: holiday.Name, Type: holiday.Type, Date: date}
	db.Create(&item)
}

func (h Holiday) ToType() types.Holiday {
	return types.Holiday{Name: h.Name, Type: h.Type, Date: h.Date.Format("2006-01-02")}
}
