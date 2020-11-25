package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func InitApp(unAuthorized *gin.RouterGroup, authorized *gin.RouterGroup) bool {
	db, ok := InitDb()
	if !ok {
		fmt.Println("Error Init Database: ")
		return false
	}
	//restricted access
	InitModels(unAuthorized, db)
	InitUsersApi(unAuthorized, db)
	InitPlacesApi(unAuthorized, db)
	InitPicturesApi(unAuthorized, db)
	InitCampaignsApi(unAuthorized, db)
	InitGiftsApi(unAuthorized, db)

	//Use authorized group for authentication
	InitInstantWinnerGamesApiRestricted(unAuthorized, db)
	InitInstantWinnerGamesApiPublic(authorized, db)

	InitRandomDrawGamesApi(unAuthorized, db)
	InitInstantWinnerPlayersApi(unAuthorized, db)
	InitRandomDrawPlayersApi(unAuthorized, db)

	return true
}
