package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const ALREADY_WON = 1
const CAN_NOT_PLAY = 2
const GAME_OVER = 3
const ALREADY_PLAYED = 4
const WON_GAME = 5
const LOST_GAME = 6

func InitInstantWinnerPlayersApi(router *gin.RouterGroup, db *gorm.DB) {
	router.GET("/api/instantWinnerPlayers", ListInstantWinnerPlayers)
	router.GET("/api/getListOfInstantWinnerPlayers", GetListOfInstantWinnerPlayers)
	router.GET("/api/instantWinnerPlayers/:id", ReadInstantWinnerPlayer)
	
	//unnecessary endpoint
	router.POST("/api/instantWinnerPlayers", CreateInstantWinnerPlayer)
	
	router.PUT("/api/instantWinnerPlayers/:id", UpdateInstantWinnerPlayer)
	router.DELETE("/api/instantWinnerPlayers/:id", DeleteInstantWinnerPlayer)
	router.POST("/api/instantWinnerPlayer", CreateInstantWinner)
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

type InstantWinnerPlayerData struct {
	InstantWinnerGameID uint64  `binding:"required" form:"instant_winner_game_id" json:"instant_winner_game_id"`
	Latitude            float32 `binding:"required" form:"latitude" json:"latitude"`
	Longitude           float32 `binding:"required" form:"longitude" json:"longitude"`
	Email               string  `binding:"required" form:"email" json:"email"`
	FingerPrint         string  `binding:"required" form:"finger_print" json:"finger_print"`
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

// unnecessary block
func CreateInstantWinnerPlayer(context *gin.Context) {
	var instantWinnerPlayer InstantWinnerPlayer
	CreateModel(context, &instantWinnerPlayer)
}

func CreateInstantWinner(context *gin.Context) {
	input := InstantWinnerPlayerData{}
	err := context.Bind(&input)
	var game InstantWinnerGame

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message" : err.Error()})
		return
	}

	// check for instantWinnerGamePlayer existence
	var player InstantWinnerPlayer
	result := AppDb.Where(`instant_winner_game_id = ? AND (fingerprint = ? OR email = ?)`, input.InstantWinnerGameID, input.FingerPrint, input.Email).First(&player)

	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": result.Error})
		return
	}

	if player.ID != 0 {
		context.AbortWithStatusJSON(http.StatusOK, gin.H{"result" : false, "code": ALREADY_PLAYED, "message" : "You had played the game already"})
		return
	}

	//save instant winner game player
	err = AppDb.Create(&InstantWinnerPlayer{
		CreatedAt:           time.Time{},
		UpdatedAt:           time.Time{},
		InstantWinnerGameID: input.InstantWinnerGameID,
		InstantWinnerGame:   nil,
		IPAddress:           "",
		Fingerprint:         input.FingerPrint,
		Email:               input.Email,
		Time:                nil,
		Result:              false,
	}).Error
	if err != nil {
		context.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	err = AppDb.Joins(`JOIN places on places.id = instant_winner_games.place_id`).Where(`instant_winner_games.id = ?`, input.InstantWinnerGameID).Preload("Place").First(&game).Error
	if err != nil {
		context.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// check whether game is won or not
	if game.Won {
		context.AbortWithStatusJSON(http.StatusOK, gin.H{"result" : false, "code": ALREADY_WON, "message" : "Game is won"})
		return
	}

	// check for location
	if (game.Place.Latitude > (input.Latitude - 0.1) && game.Place.Latitude < (input.Latitude + 0.1) &&
		game.Place.Longitude > (input.Longitude - 0.1) && game.Place.Longitude < (input.Longitude + 0.1)) == false {
		context.AbortWithStatusJSON(http.StatusOK, gin.H{"result" : false, "code": CAN_NOT_PLAY, "message" : "You can't play from here"})
		return
	}

	now := time.Now()
	if now.After(*game.EndDate) {
		context.AbortWithStatusJSON(http.StatusOK, gin.H{"result" : false, "code": GAME_OVER, "message" : "Game is over"})
		return
	}

	// check serverTime passed secretTime
	if now.After(*game.PlayTime) {
		response := AppDb.Table(`instant_winner_games`).Where("id = ?", input.InstantWinnerGameID).Update("won", true)
		if response.Error != nil && !errors.Is(response.Error, gorm.ErrRecordNotFound){
			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": response.Error})
			return
		}
		context.JSON(http.StatusOK, gin.H{"result" : true, "code": WON_GAME,"message" : "Congratulations !! you won the game"})
		return
	}

	// if serverTime not passed secretTime
	context.JSON(http.StatusOK, gin.H{"result" : false, "code": LOST_GAME, "message" : "You lost the game"})
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