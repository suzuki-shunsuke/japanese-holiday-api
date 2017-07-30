package lib

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/suzuki-shunsuke/japanese-holiday-api/models"
	"github.com/suzuki-shunsuke/japanese-holiday-api/types"
	"sort"
	"time"
)

func GetHolidays(c echo.Context) error {
	config, _ := GetConfig()
	db := GetConnection(config)
	var holidays_ []models.Holiday
	holidays := map[string]types.Holiday{}
	var holiday_list types.Holidays
	db.Select("name, type, date, day_of_week").Find(&holidays_)
	var startDate time.Time
	var endDate time.Time
	for i, h := range holidays_ {
		if i == 0 {
			startDate = time.Date(h.Date.Year(), 1, 1, 0, 0, 0, 0, time.UTC)
			endDate = time.Date(h.Date.Year(), 1, 1, 0, 0, 0, 0, time.UTC)
		}
		ht := h.ToType()
		holidays[ht.Date] = ht
		holiday_list = append(holiday_list, ht)
		endDate = h.Date
	}
	endDate = time.Date(endDate.Year()+1, 1, 1, 0, 0, 0, 0, time.UTC)
	// add sunday
	for date := startDate.AddDate(0, 0, (7-int(startDate.Weekday()))%7); date.Before(endDate); date = date.AddDate(0, 0, 7) {
		date_str := date.Format("2006-01-02")
		_, ok := holidays[date_str]
		if !ok {
			ht := types.Holiday{Name: "", DayOfWeek: 0, Type: 0, Date: date_str}
			holiday_list = append(holiday_list, ht)
			holidays[date_str] = ht
			continue
		}
		// http://www8.cao.go.jp/chosei/shukujitsu/gaiyou.html
		// 2.「国民の祝日」が日曜日に当たるときは、その日後においてその日に最も近い「国民の祝日」でない日を休日とする。
		for alter_date := date.AddDate(0, 0, 1); alter_date.Before(date.AddDate(0, 0, 7)); alter_date = alter_date.AddDate(0, 0, 1) {
			date_str := alter_date.Format("2006-01-02")
			_, ok := holidays[date_str]
			if !ok {
				ht := types.Holiday{Name: "", DayOfWeek: int(alter_date.Weekday()), Type: 2, Date: date_str}
				holiday_list = append(holiday_list, ht)
				holidays[date_str] = ht
				break
			}
		}
	}
	sort.Sort(holiday_list)
	return c.JSON(http.StatusOK, holiday_list)
}

func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
