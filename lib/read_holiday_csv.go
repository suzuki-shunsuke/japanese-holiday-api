package lib

import (
	"encoding/csv"
	"github.com/suzuki-shunsuke/japanese-holiday-api/models"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"io"
	"log"
	"os"
	"time"
)

func ReadHolidayCsv(path string) []models.Holiday {
	file, err := os.Open(path)
	failOnError(err)
	defer file.Close()
	reader := csv.NewReader(transform.NewReader(file, japanese.ShiftJIS.NewDecoder()))
	reader.LazyQuotes = true // ダブルクオートを厳密にチェックしない
	// remove header
	_, err = reader.Read()
	if err == io.EOF {
		failOnError(err)
	}
	failOnError(err)
	var holidays []models.Holiday

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		failOnError(err)
		date, _ := time.Parse("2006-01-02", record[0])
		holidays = append(
			holidays,
			models.Holiday{
				Date:      date,
				Name:      record[1],
				Type:      1,
				DayOfWeek: int(date.Weekday())})
	}
	return holidays
}

func failOnError(err error) {
	if err != nil {
		log.Fatal("Error:", err)
	}
}
