package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitInstantWinnerPlayersApi(router *gin.Engine, db *gorm.DB) {
	router.GET("/api/instantWinnerPlayers", ListInstantWinnerPlayers)
	router.GET("/api/getListOfInstantWinnerPlayers", GetListOfInstantWinnerPlayers)
	router.GET("/api/instantWinnerPlayers/:id", ReadInstantWinnerPlayer)
	router.POST("/api/instantWinnerPlayers", CreateInstantWinnerPlayer)
	router.PUT("/api/instantWinnerPlayers/:id", UpdateInstantWinnerPlayer)
	router.DELETE("/api/instantWinnerPlayers/:id", DeleteInstantWinnerPlayer)
}
func ListInstantWinnerPlayers(context *gin.Context) {
	var instantWinnerPlayers []InstantWinnerPlayer
	ListModels(context, &instantWinnerPlayers)
}

type InstWinPlayersQuery struct {
	CampaignID int64
	IPAddress  string
	Longitude  float64
	Latitude   float64
	Date       *time.Time
}

func GetListOfInstantWinnerPlayers(context *gin.Context) {
	var instantWinnerPlayers []InstantWinnerPlayer
	var query InstWinPlayersQuery
	if context.ShouldBindQuery(&query) != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": query})
		return
	}
	if query.Date == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid Date"})
		return
	}
	fmt.Printf("query=%v", query)
	result := AppDb.Joins("InstantWinnerGame.Place").Joins("InstantWinnerGame").Find(
		&instantWinnerPlayers, "IP_Address=? AND InstantWinnerGame__campaign_id=?", query.IPAddress, query.CampaignID)
	// query.CampaignID, query.Date, query.Date).Preload(
	// "InstantWinnerGame.Place", "Longitude=? AND Latitude =?",
	// query.Longitude, query.Latitude).Find(&instantWinnerPlayers)*/
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Model Not Found"})
		return
	}
	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	}
	context.JSON(http.StatusOK, instantWinnerPlayers)
}

func ReadInstantWinnerPlayer(context *gin.Context) {
	var instantWinnerPlayer InstantWinnerPlayer
	ReadModel(context, &instantWinnerPlayer)
}
func CreateInstantWinnerPlayer(context *gin.Context) {
	var instantWinnerPlayer InstantWinnerPlayer
	CreateModel(context, &instantWinnerPlayer)
}
func UpdateInstantWinnerPlayer(context *gin.Context) {
	var instantWinnerPlayer InstantWinnerPlayer
	UpdateModel(context, &instantWinnerPlayer, setInstantWinnerPlayerId)
}
func DeleteInstantWinnerPlayer(context *gin.Context) {
	DeleteModel(context, &InstantWinnerPlayer{})
}
func setInstantWinnerPlayerId(id uint64, model interface{}) {
	instantWinnerPlayer := model.(*InstantWinnerPlayer)
	instantWinnerPlayer.ID = id
}
