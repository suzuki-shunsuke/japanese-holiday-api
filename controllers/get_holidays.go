package controllers

import (
	"net/http"

	"encoding/json"
	"github.com/labstack/echo"
	"github.com/suzuki-shunsuke/japanese-holiday-api/lib"
	"github.com/suzuki-shunsuke/japanese-holiday-api/models"
	"github.com/suzuki-shunsuke/japanese-holiday-api/types"
	"sort"
	"time"
)

func getHolidayList(holidays_ *[]models.Holiday, startDate *time.Time, endDate *time.Time) (holiday_list types.Holidays) {
	holidays := map[string]types.Holiday{}
	var prev_holiday time.Time
	for i, h := range *holidays_ {
		if i == 0 {
			prev_holiday = h.Date
		}
		ht := h.ToType()
		holidays[ht.Date] = ht
		holiday_list = append(holiday_list, ht)
		// http://www8.cao.go.jp/chosei/shukujitsu/gaiyou.html
		// 3.その前日及び翌日が「国民の祝日」である日（「国民の祝日」でない日に限る。）は、休日とする。
		if prev_holiday.AddDate(0, 0, 2).Equal(h.Date) {
			ht_ := prev_holiday.AddDate(0, 0, 1)
			ht = types.Holiday{Name: "", Type: 3, Date: ht_.Format("2006-01-02"), DayOfWeek: int(ht_.Weekday())}
			holiday_list = append(holiday_list, ht)
			holidays[ht.Date] = ht
		}
		prev_holiday = h.Date
	}
	// add sunday
	for date := startDate.AddDate(0, 0, (7-int(startDate.Weekday()))%7); date.Before(*endDate); date = date.AddDate(0, 0, 7) {
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
	return holiday_list
}

func GetHolidays(c echo.Context) error {
	req := new(types.Request)
	q := c.QueryParam("q")
	var holidays_ []models.Holiday
	startDate := time.Date(2016, 1, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)
	config, _ := lib.GetConfig()
	db := lib.GetConnection(config)
	query := db.Debug().Select("name, type, date, day_of_week")
	if len(q) > 0 {
		if err := json.Unmarshal(([]byte)(q), req); err != nil {
			ret := map[string]string{"message": "The format of 'q' parameter is invalid"}
			return c.JSON(http.StatusBadRequest, ret)
		}
		// Convert From string to time.Time
		if len(req.From) > 0 {
			from_time, err := time.Parse("2006-01-02", req.From)
			if err != nil {
				ret := map[string]string{"message": "The format of 'from' paramater is invalid"}
				return c.JSON(http.StatusBadRequest, ret)
			}
			query = query.Where("date >= ?", from_time.Format("2006-01-02"))
			startDate = from_time
		}
		if len(req.To) > 0 {
			to_time, err := time.Parse("2006-01-02", req.To)
			if err != nil {
				ret := map[string]string{"message": "The format of 'to' paramater is invalid"}
				return c.JSON(http.StatusBadRequest, ret)
			}
			query = query.Where("date < ?", to_time.Format("2006-01-02"))
			endDate = to_time
		}
	}
	query.Find(&holidays_)
	holiday_list := getHolidayList(&holidays_, &startDate, &endDate)

	sort.Sort(holiday_list)
	return c.JSON(http.StatusOK, holiday_list)
}
