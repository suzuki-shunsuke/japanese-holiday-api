package services

import (
	"encoding/csv"
	"github.com/suzuki-shunsuke/japanese-holiday-api/models"
	"github.com/suzuki-shunsuke/japanese-holiday-api/types"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"io"
	"net/http"
	"os"
	"time"
)

func ReadHolidayCsv(path string) ([]models.Holiday, *types.AppError) {
	file, err := os.Open(path)
	if err != nil {
		return nil, &types.AppError{Code: http.StatusInternalServerError, Message: "Internal Server Error"}
	}
	defer file.Close()
	reader := csv.NewReader(transform.NewReader(file, japanese.ShiftJIS.NewDecoder()))
	reader.LazyQuotes = true // ダブルクオートを厳密にチェックしない
	// remove header
	_, err = reader.Read()
	if err != nil {
		return nil, &types.AppError{Code: http.StatusInternalServerError, Message: "Internal Server Error"}
	}
	var holidays []models.Holiday

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, &types.AppError{Code: http.StatusInternalServerError, Message: "Internal Server Error"}
		}
		date, _ := time.Parse("2006-01-02", record[0])
		holidays = append(
			holidays,
			models.Holiday{
				Date:      date,
				Name:      record[1],
				Type:      1,
				DayOfWeek: int(date.Weekday())})
	}
	return holidays, nil
}
