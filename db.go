package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var AppDb *gorm.DB = nil

func InitDb() (*gorm.DB, bool) {
	dbUsername := os.Getenv("MYSQL_ROOT_USER")
	dbPassword := os.Getenv("MYSQL_ROOT_PASSWORD")
	dbHost := os.Getenv("MYSQL_HOST")
	dbPort := os.Getenv("MYSQL_PORT")
	dbName := os.Getenv("MYSQL_DATABASE_NAME")

	var err error

	dsn := dbUsername + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?parseTime=true&&loc=Local"
	AppDb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//dsn := "user=isystech password=itsmeilhem dbname=gifts_winner port=5432 sslmode=disable TimeZone=Asia/Jerusalem"
	//AppDb, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	AppDb.LogMode(true)
	if err != nil {
		fmt.Println("InitDb(): Unable to Connect to Database: error=%v", err)
		return nil, false
	}
	return AppDb, true
}
