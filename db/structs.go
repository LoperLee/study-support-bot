package db

import (
	"time"

	"gorm.io/gorm"
)

// this file is collect all structs
// using by database

// save commit history struct
type CommitHistory struct {
	gorm.Model
	Id         int
	UserId     int
	CommitDate time.Time
}

// user table struct
type Users struct {
	gorm.Model
	Id    int
	Name  string
	Email string
}
