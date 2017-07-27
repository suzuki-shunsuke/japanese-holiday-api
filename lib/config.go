package lib

import (
	"github.com/BurntSushi/toml"
	"github.com/suzuki-shunsuke/japanese-holiday-api/types"
	"os"
	"strconv"
)

func GetConfig() (*types.Config, error) {
	var config types.Config
	_, err := toml.DecodeFile("config.toml", &config)
	password := os.Getenv("RDB_PASSWORD")
	host := os.Getenv("RDB_HOST")
	dbName := os.Getenv("RDB_DBNAME")
	port_str := os.Getenv("RDB_PORT")
	app_port_str := os.Getenv("APP_PORT")
	if len(password) > 0 {
		config.RDB.Password = password
	}
	if len(host) > 0 {
		config.RDB.Host = host
	}
	if len(dbName) > 0 {
		config.RDB.DBName = dbName
	}
	if len(port_str) > 0 {
		config.RDB.Port, err = strconv.Atoi(port_str)
	}
	if len(app_port_str) > 0 {
		config.App.Port, err = strconv.Atoi(app_port_str)
	}

	return &config, err
}
