package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func InitApp(router *gin.Engine) bool {
	db, ok := InitDb()
	if !ok {
		fmt.Println("Error Init Database: ")
		return false
	}
	InitModels(router, db)
	InitUsersApi(router, db)
	InitPlacesApi(router, db)
	InitPicturesApi(router, db)
	InitCampaignsApi(router, db)
	InitGiftsApi(router, db)
	InitInstantWinnerGamesApi(router, db)
	InitRandomDrawGamesApi(router, db)
	InitInstantWinnerPlayersApi(router, db)
	InitRandomDrawPlayersApi(router, db)

	return true
}
