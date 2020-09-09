package db
// this package using database

import (
	"fmt"
	"time"

	cng "github.com/LoperLee/study-support-bot/config"
	_ "github.com/go-sql-driver/mysql"
)

// save commit history struct
type CommitHistory struct {
	Id         int
	UserId     int
	CommitDate time.Time
}

// user table struct
type Users struct {
	Id    int
	Name  string
	Email string
}

// db connection string format
func getDBConnectionAddress(cofing cng.Config) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true"), cofing.Database.User, cofing.Database.Pass, cofing.Database.Addr, cofing.Database.Port, cofing.Database.Name)
}

// initialize connection
func InitDbConnection(cofing cng.Config) *sql.DB {
	db, err := sql.Open("mysql", getDBConnectionAddress(config))
	if err != nil {
		panic(err)
	}
	return db
}
