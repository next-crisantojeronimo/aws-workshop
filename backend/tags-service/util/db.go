package util

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// GetConnection gives you a DB connection
func GetConnection() *gorm.DB {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("GO_APP_DB_USR"),
		os.Getenv("GO_APP_DB_PASSWD"),
		os.Getenv("GO_APP_DB_HOST"),
		os.Getenv("GO_APP_DB_PORT"),
	)
	connection, err := gorm.Open("mysql", dataSource)
	connection.Exec(fmt.Sprintf("USE `%s`", os.Getenv("GO_APP_DB_SCHEME")))
	connection.Exec("SET FOREIGN_KEY_CHECKS=0;")

	if err != nil {
		log.Fatal(err)
	}

	return connection
}
