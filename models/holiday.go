package models

import (
	"github.com/suzuki-shunsuke/japanese-holiday-api/types"
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

func GetNationalHolidaysByRDB(startDate *time.Time, endDate *time.Time, config *types.Config) (holidays Holidays, err *types.AppError) {
	db, app_err := GetConnection(config)
	if app_err != nil {
		return nil, app_err
	}
	query := db.Select("name, type, date, day_of_week").Where("date >= ? AND date < ?", startDate.Format("2006-01-02"), endDate.Format("2006-01-02"))
	if config.RDB.Debug {
		query = query.Debug()
	}
	query.Find(&holidays)
	return holidays, nil
}

func GetHolidayList(holidays_ *Holidays, startDate *time.Time, endDate *time.Time) (holiday_list Holidays) {
	holidays := map[string]Holiday{}
	var prev_holiday time.Time
	for i, h := range *holidays_ {
		if i == 0 {
			prev_holiday = h.Date
		}
		holidays[h.StringDate()] = h
		holiday_list = append(holiday_list, h)
		// http://www8.cao.go.jp/chosei/shukujitsu/gaiyou.html
		// 3.その前日及び翌日が「国民の祝日」である日（「国民の祝日」でない日に限る。）は、休日とする。
		if prev_holiday.AddDate(0, 0, 2).Equal(h.Date) {
			ht_ := prev_holiday.AddDate(0, 0, 1)
			ht := Holiday{
				Name:      "",
				Type:      3,
				Date:      ht_,
				DayOfWeek: int(ht_.Weekday())}
			holiday_list = append(holiday_list, ht)
			holidays[ht.StringDate()] = ht
		}
		prev_holiday = h.Date
	}
	// add sunday
	for date := startDate.AddDate(0, 0, (7-int(startDate.Weekday()))%7); date.Before(*endDate); date = date.AddDate(0, 0, 7) {
		date_str := date.Format("2006-01-02")
		_, ok := holidays[date_str]
		if !ok {
			ht := Holiday{
				Name:      "",
				DayOfWeek: 0,
				Type:      0,
				Date:      date}
			holiday_list = append(holiday_list, ht)
			holidays[date_str] = ht
			continue
		}
		// http://www8.cao.go.jp/chosei/shukujitsu/gaiyou.html
		// 2.「国民の祝日」が日曜日に当たるときは、その日後においてその日に最も近い「国民の祝日」でない日を休日とする。
		for alter_date := date.AddDate(0, 0, 1); alter_date.Before(date.AddDate(0, 0, 7)) && alter_date.Before(*endDate); alter_date = alter_date.AddDate(0, 0, 1) {
			date_str := alter_date.Format("2006-01-02")
			_, ok := holidays[date_str]
			if !ok {
				ht := Holiday{
					Name:      "",
					DayOfWeek: int(alter_date.Weekday()),
					Type:      2,
					Date:      alter_date}
				holiday_list = append(holiday_list, ht)
				holidays[date_str] = ht
				break
			}
		}
	}
	return holiday_list
}
