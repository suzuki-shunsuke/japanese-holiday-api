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

func getNationalHolidays(req *types.Request, startDate *time.Time, endDate *time.Time, config *types.Config) (holidays_ []models.Holiday, err *types.AppError) {
	db := lib.GetConnection(config)
	query := db.Debug().Select("name, type, date, day_of_week")
	if len(req.From) > 0 {
		query = query.Where("date >= ?", startDate.Format("2006-01-02"))
	}
	if len(req.To) > 0 {
		query = query.Where("date < ?", endDate.Format("2006-01-02"))
	}
	query.Find(&holidays_)
	return holidays_, nil
}

func parseQuery(q string, startDate *time.Time, endDate *time.Time) (*types.Request, *types.AppError) {
	var err error
	req := new(types.Request)
	if len(q) > 0 {
		if json.Unmarshal(([]byte)(q), req) != nil {
			return nil, &types.AppError{Code: http.StatusBadRequest, Message: "The format of 'q' parameter is invalid"}
		}
		// Convert From string to time.Time
		if len(req.From) > 0 {
			*startDate, err = time.Parse("2006-01-02", req.From)
			if err != nil {
				return nil, &types.AppError{Code: http.StatusBadRequest, Message: "The format of 'from' parameter is invalid"}
			}
		}
		if len(req.To) > 0 {
			*endDate, err = time.Parse("2006-01-02", req.To)
			if err != nil {
				return nil, &types.AppError{Code: http.StatusBadRequest, Message: "The format of 'to' parameter is invalid"}
			}
		}
	}
	return req, nil
}

func GetHolidays(c echo.Context) error {
	q := c.QueryParam("q")
	startDate := time.Date(2016, 1, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)
	config, _ := lib.GetConfig()
	req, err := parseQuery(q, &startDate, &endDate)
	if err != nil {
		return c.JSON(err.Code, map[string]string{"message": err.Message})
	}
	holidays_, err := getNationalHolidays(req, &startDate, &endDate, config)
	if err != nil {
		return c.JSON(err.Code, map[string]string{"message": err.Message})
	}
	holiday_list := getHolidayList(&holidays_, &startDate, &endDate)

	sort.Sort(holiday_list)
	return c.JSON(http.StatusOK, holiday_list)
}
