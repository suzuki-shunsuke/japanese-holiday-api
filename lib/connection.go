package lib

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/suzuki-shunsuke/japanese-holiday-api/types"
	"net/http"
	"strconv"
)

func GetConnection(config *types.Config) (*gorm.DB, *types.AppError) {
	protocol := "tcp(" + config.RDB.Host + ":" + strconv.Itoa(config.RDB.Port) + ")"
	CONNECT := config.RDB.User + ":" + config.RDB.Password + "@" + protocol + "/" + config.RDB.DBName + "?parseTime=true"
	db, err := gorm.Open(config.RDB.Dbms, CONNECT)

	if err != nil {
		return nil, &types.AppError{Code: http.StatusInternalServerError, Message: "Internal Server Error"}
	}
	return db, nil
}
