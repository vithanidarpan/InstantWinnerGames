package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRandomDrawGamesApi(router *gin.Engine, db *gorm.DB) {
	router.GET("/api/randomDrawGames", ListRandomDrawGames)
	router.GET("/api/randomDrawGames/:id", ReadRandomDrawGame)
	router.POST("/api/randomDrawGrams", CreateRandomDrawGame)
	router.PUT("/api/randomDrawGames/:id", UpdateRandomDrawGame)
	router.DELETE("/api/randomDrawGames/:id", DeleteRandomDrawGame)
}
func ListRandomDrawGames(context *gin.Context) {
	var randomDrawGames []RandomDrawGame
	ListModels(context, &randomDrawGames)
}
func ReadRandomDrawGame(context *gin.Context) {
	var randomDrawGame RandomDrawGame
	ReadModel(context, &randomDrawGame)
}
func CreateRandomDrawGame(context *gin.Context) {
	var randomDrawGame RandomDrawGame
	CreateModel(context, &randomDrawGame)
}
func UpdateRandomDrawGame(context *gin.Context) {
	var randomDrawGame RandomDrawGame
	UpdateModel(context, &randomDrawGame, setRandomDrawGameId)
}
func DeleteRandomDrawGame(context *gin.Context) {
	DeleteModel(context, &RandomDrawGame{})
}
func setRandomDrawGameId(id uint64, model interface{}) {
	randomDrawGame := model.(*RandomDrawGame)
	randomDrawGame.ID = id
}
