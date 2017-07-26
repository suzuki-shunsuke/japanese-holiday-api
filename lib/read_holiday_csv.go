package lib

import (
	"encoding/csv"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"io"
	"log"
	"os"
)

func ReadHolidayCsv(path string) []Holiday {
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
	var holidays []Holiday

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		failOnError(err)
		holidays = append(holidays, Holiday{Date: record[0], Name: record[1], Type: 1})
	}
	return holidays
}

func failOnError(err error) {
	if err != nil {
		log.Fatal("Error:", err)
	}
}
