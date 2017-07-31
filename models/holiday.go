package models

import (
	"time"
)

type Holiday struct {
	ID        int       `json:"id"`
	Name      string    `gorm:"not null;type:varchar(100)" json:"name"`
	Date      time.Time `gorm:"not null;unique_index;type:date" json:"date"`
	Type      int       `gorm:"not null" json:"type"`
	DayOfWeek int       `gorm:"not null" json:"day_of_week"`
}

func (h Holiday) StringDate() string {
	return h.Date.Format("2006-01-02")
}

type Holidays []Holiday

func (h Holidays) Len() int {
	return len(h)
}

func (h Holidays) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h Holidays) Less(i, j int) bool {
	return h[i].Date.Before(h[j].Date)
}

func (h Holiday) Map(keys []string) (ret map[string]interface{}) {
	ret = map[string]interface{}{}
	for _, key := range keys {
		if key == "id" {
			ret["id"] = h.ID
		} else if key == "name" {
			ret["name"] = h.Name
		} else if key == "date" {
			ret["date"] = h.Date
		} else if key == "type" {
			ret["type"] = h.Type
		} else if key == "day_of_week" {
			ret["day_of_week"] = h.DayOfWeek
		}
	}
	return ret
}
