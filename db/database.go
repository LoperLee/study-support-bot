package db

// this package using database

import (
	"fmt"

	cng "github.com/LoperLee/study-support-bot/config"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Connection struct {
	gorm.DB
}

// db connection string format
func getDBConnectionAddress(cofing cng.Config) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true", cofing.Database.User, cofing.Database.Pass, cofing.Database.Addr, cofing.Database.Port, cofing.Database.Name)
}

// initialize connection
func InitDbConnection() *gorm.DB {
	config := cng.LoadConfigration("./config.yml")
	db, err := gorm.Open(mysql.Open(getDBConnectionAddress(config)), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

// get users
func GetAllUser() {

}
