package types

import (
	"fmt"
)

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
	RDB RDBConfig
	App AppConfig
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
