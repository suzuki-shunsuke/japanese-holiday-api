package lib

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GetConnection() *gorm.DB {
	DBMS := "mysql"
	USER := "root"
	PASS := "password"
	PROTOCOL := "tcp(localhost:4306)"
	DBNAME := "app"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	return db
}
