package db

import (
	"database/sql"
	"ddd/infrastructure/setting"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func GetDBconnection(dbSetting setting.DB) (*sql.DB, error) {
	var dataSource string = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		dbSetting.User,
		dbSetting.Password,
		dbSetting.Host,
		dbSetting.Port,
		dbSetting.Name,
	)
	log.Println(dataSource)
	db, err := sql.Open(dbSetting.Type, dataSource)
	return db, err
}
