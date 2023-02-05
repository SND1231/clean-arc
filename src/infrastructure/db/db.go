package db

import (
	"database/sql"
	"ddd/infrastructure/setting"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func GetDBconnection(dbSetting setting.DB) (*sql.DB, error) {
	var dataSource string = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		dbSetting.User,
		dbSetting.Password,
		dbSetting.Host,
		dbSetting.Port,
		dbSetting.Name,
	)
	dataSource = dataSource + "&loc=Asia%2FTokyo"
	db, err := sql.Open(dbSetting.Type, dataSource)
	return db, err
}
