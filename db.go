package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var AppDb *gorm.DB = nil

func InitDb() (*gorm.DB, bool) {
	var err error
	dsn := "root:root@tcp(127.0.0.1:3306)/instant_winner_games?parseTime=true&&loc=Local"
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