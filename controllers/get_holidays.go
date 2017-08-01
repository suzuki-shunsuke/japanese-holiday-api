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

func getNationalHolidays(startDate *time.Time, endDate *time.Time, config *types.Config) (holidays models.Holidays, app_err *types.AppError) {
	if config.App.Storage == "rdb" {
		return models.GetNationalHolidaysByRDB(startDate, endDate, config)
	}
	if config.App.Storage == "sjis_csv" {
		holidays_, app_err := lib.ReadHolidayCsv(config.SjisCsv.Path)
		if app_err != nil {
			return holidays_, app_err
		}
		for _, holiday := range holidays_ {
			if !holiday.Date.Before(*startDate) && holiday.Date.Before(*endDate) {
				holidays = append(holidays, holiday)
			}
		}
		return holidays, nil
	}
	return nil, &types.AppError{Code: http.StatusInternalServerError, Message: "Internal Server Error"}
}

func parseQuery(q string, startDate *time.Time, endDate *time.Time) *types.AppError {
	var err error
	req := new(types.Request)
	if len(q) > 0 {
		if json.Unmarshal(([]byte)(q), req) != nil {
			return &types.AppError{Code: http.StatusBadRequest, Message: "The format of 'q' parameter is invalid"}
		}
		// Convert From string to time.Time
		if len(req.From) > 0 {
			*startDate, err = time.Parse("2006-01-02", req.From)
			if err != nil {
				return &types.AppError{Code: http.StatusBadRequest, Message: "The format of 'from' parameter is invalid"}
			}
		}
		if len(req.To) > 0 {
			*endDate, err = time.Parse("2006-01-02", req.To)
			if err != nil {
				return &types.AppError{Code: http.StatusBadRequest, Message: "The format of 'to' parameter is invalid"}
			}
		}
	}
	return nil
}

func GetHolidays(c echo.Context) error {
	config, _ := lib.GetConfig()
	startDate, err := time.Parse("2006-01-02", config.App.StartDate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal Server Error"})
	}
	endDate, err := time.Parse("2006-01-02", config.App.EndDate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal Server Error"})
	}
	if err := parseQuery(c.QueryParam("q"), &startDate, &endDate); err != nil {
		return c.JSON(err.Code, map[string]string{"message": err.Message})
	}
	holiday_list, app_err := getNationalHolidays(&startDate, &endDate, config)
	if app_err != nil {
		return c.JSON(app_err.Code, map[string]string{"message": app_err.Message})
	}
	if !config.RDB.IsOtherHolidaysStored {
		holiday_list = models.GetHolidayList(&holiday_list, &startDate, &endDate)
	}
	sort.Sort(holiday_list)
	var holidays []map[string]interface{}
	for _, h := range holiday_list {
		ht := h.Map([]string{"name", "date", "type", "day_of_week"})
		ht["date"] = h.StringDate()
		holidays = append(holidays, ht)
	}
	return c.JSON(http.StatusOK, holidays)
}
