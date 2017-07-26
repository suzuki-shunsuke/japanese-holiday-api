package lib

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/suzuki-shunsuke/japanese-holiday-api/types"
	"strconv"
)

func GetConnection(config *types.Config) *gorm.DB {
	protocol := "tcp(" + config.RDB.Host + ":" + strconv.Itoa(config.RDB.Port) + ")"
	CONNECT := config.RDB.User + ":" + config.RDB.Password + "@" + protocol + "/" + config.RDB.DBName + "?parseTime=true"
	db, err := gorm.Open(config.RDB.Dbms, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	return db
}
