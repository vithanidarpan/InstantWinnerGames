package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRandomDrawPlayersApi(router *gin.RouterGroup, db *gorm.DB) {
	router.GET("/api/randomDrawPlayers", ListRandomDrawPlayers)
	router.GET("/api/randomDrawPlayers/:id", ReadRandomDrawPlayer)
	router.POST("/api/randomDrawPlayers", CreateRandomDrawPlayer)
	router.PUT("/api/randomDrawPlayers/:id", UpdateRandomDrawPlayer)
	router.DELETE("/api/randomDrawPlayers/:id", DeleteRandomDrawPlayer)
}
func ListRandomDrawPlayers(context *gin.Context) {
	var randomDrawPlayers []RandomDrawPlayer
	ListModels(context, &randomDrawPlayers)
}
func ReadRandomDrawPlayer(context *gin.Context) {
	var randomDrawPlayer RandomDrawPlayer
	ReadModel(context, &randomDrawPlayer)
}
func CreateRandomDrawPlayer(context *gin.Context) {
	var randomDrawPlayer RandomDrawPlayer
	CreateModel(context, &randomDrawPlayer)
}
func UpdateRandomDrawPlayer(context *gin.Context) {
	var randomDrawPlayer RandomDrawPlayer
	UpdateModel(context, &randomDrawPlayer, setRandomDrawPlayerId)
}
func DeleteRandomDrawPlayer(context *gin.Context) {
	DeleteModel(context, &RandomDrawPlayer{})
}
func setRandomDrawPlayerId(id uint64, model interface{}) {
	randomDrawPlayer := model.(*RandomDrawPlayer)
	randomDrawPlayer.ID = id
}
