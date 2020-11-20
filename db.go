package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	//"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var AppDb *gorm.DB = nil

func InitDb() (*gorm.DB, bool) {
	var err error
	dbName := "/Users/michaelmerlange/Downloads/sqlite-tools-osx-x86-3330000/gifts_winner.db"
	AppDb, err = gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	//dsn := "user=isystech password=itsmeilhem dbname=gifts_winner port=5432 sslmode=disable TimeZone=Asia/Jerusalem"
	//AppDb, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	AppDb.LogMode(true)
	if err != nil {
		fmt.Println("InitDb(): Unable to Connect to Database: error=%v", err)
		return nil, false
	}
	return AppDb, true
}
