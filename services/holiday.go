package services

import (
	"github.com/suzuki-shunsuke/japanese-holiday-api/models"
	"github.com/suzuki-shunsuke/japanese-holiday-api/types"
	"net/http"
	"time"
)

func GetNationalHolidays(startDate *time.Time, endDate *time.Time, config *types.Config) (holidays models.Holidays, app_err *types.AppError) {
	if config.App.Storage == "rdb" {
		return GetNationalHolidaysByRDB(startDate, endDate, config)
	}
	if config.App.Storage == "sjis_csv" {
		holidays_, app_err := ReadHolidayCsv(config.SjisCsv.Path)
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

func GetNationalHolidaysByRDB(startDate *time.Time, endDate *time.Time, config *types.Config) (holidays models.Holidays, err *types.AppError) {
	db, app_err := models.GetConnection(config)
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
