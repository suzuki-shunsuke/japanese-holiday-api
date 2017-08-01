package types

import (
	"fmt"
)

type SjisCsvConfig struct {
	Path string
}

type AppConfig struct {
	Port      int
	StartDate string
	EndDate   string
	Storage   string
}

type RDBConfig struct {
	Dbms                  string
	User                  string
	Password              string
	Host                  string
	Port                  int
	DBName                string
	Protocol              string
	IsOtherHolidaysStored bool
	Debug                 bool
}

type Config struct {
	RDB     RDBConfig
	App     AppConfig
	SjisCsv SjisCsvConfig `toml:"sjis_csv"`
}

type Request struct {
	From  string `json:"from" form:"from" query:"from"`
	To    string `json:"to" form:"to" query:"to"`
	Types []int  `json:"types" form:"types" query:"types"`
}

type AppError struct {
	Message string
	Code    int
}

func (err *AppError) Error() string {
	return fmt.Sprintf("err %s [code=%d]", err.Message, err.Code)
}
