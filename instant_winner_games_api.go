package main

import (
	//"errors"
	//"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	//"strconv"
	"time"
)

func InitInstantWinnerGamesApiRestricted(router *gin.RouterGroup, db *gorm.DB) {
	router.POST("/api/instantWinnerGames", CreateInstantWinnerGame)
	router.PUT("/api/instantWinnerGames/:id", UpdateInstantWinnerGame)
	router.DELETE("/api/instantWinnerGames/:id", DeleteInstantWinnerGame)
}

func InitInstantWinnerGamesApiPublic(router *gin.RouterGroup, db *gorm.DB) {
	router.GET("/api/instantWinnerGames", ListInstantWinnerGames)
	router.GET("/api/openedInstantWinnerGames", OpenedInstantGames)
}

func ListInstantWinnerGames(context *gin.Context) {
	var instantWinnerGames []InstantWinnerGame
	ListModels(context, &instantWinnerGames)
}

type InstWinGamesQuery struct {
	Longitude  float32 `form:"longitude" json:"longitude"`
	Latitude   float32 `form:"latitude" json:"latitude"`
}


func OpenedInstantGames(context *gin.Context) {
	var instantWinnerGames []InstantWinnerGame
	var input InstWinGamesQuery
	var err error

	var serverTime = time.Now().Format("2006-01-02 15:04:05")

	if err = context.Bind(&input); err != nil {
		context.AbortWithStatus(http.StatusBadRequest)
	}

	err = AppDb.Joins("JOIN places on places.id=instant_winner_games.place_id").
		Where("won = ? AND start_date < ? AND end_Date > ? AND places.latitude BETWEEN ? AND ? AND places.longitude BETWEEN ? AND ?", false, serverTime, serverTime, input.Latitude - 0.1, input.Latitude + 0.1, input.Longitude - 0.1, input.Longitude + 0.1 ).
		Preload("Place").Find(&instantWinnerGames).Error

	if err != nil {
		context.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	context.JSON(http.StatusOK, instantWinnerGames)
}

func ReadInstantWinnerGame(context *gin.Context) {
	var instantWinnerGame InstantWinnerGame
	ReadModel(context, &instantWinnerGame)
}
func CreateInstantWinnerGame(context *gin.Context) {
	var instantWinnerGame InstantWinnerGame
	CreateModel(context, &instantWinnerGame)
}
func UpdateInstantWinnerGame(context *gin.Context) {
	var instantWinnerGame InstantWinnerGame
	UpdateModel(context, &instantWinnerGame, setInstantWinnerGameId)
}
func DeleteInstantWinnerGame(context *gin.Context) {
	DeleteModel(context, &InstantWinnerGame{})
}
func setInstantWinnerGameId(id uint64, model interface{}) {
	instantWinnerGame := model.(*InstantWinnerGame)
	instantWinnerGame.ID = id
}
