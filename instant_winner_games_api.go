package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

func InitInstantWinnerGamesApi(router *gin.RouterGroup, db *gorm.DB) {
	router.GET("/api/instantWinnerGames", ListInstantWinnerGames)

	//unnecessary endpoint
	router.GET("/api/openedInstantWinnerGames/:lat/:lon", GetListOfOpenedInstantGames)

	router.GET("/api/instantWinnerGames/:id", ReadInstantWinnerGame)
	router.POST("/api/instantWinnerGames", CreateInstantWinnerGame)
	router.PUT("/api/instantWinnerGames/:id", UpdateInstantWinnerGame)
	router.DELETE("/api/instantWinnerGames/:id", DeleteInstantWinnerGame)
	router.GET("/api/GetListOfOpenedInstantGames", OpenedInstantGames)
}
func ListInstantWinnerGames(context *gin.Context) {
	var instantWinnerGames []InstantWinnerGame
	ListModels(context, &instantWinnerGames)
}

type InstWinGamesQuery struct {
	Longitude  float32 `form:"longitude" json:"longitude"`
	Latitude   float32 `form:"latitude" json:"latitude"`
}

//unnecessary block
func GetListOfOpenedInstantGames(context *gin.Context) {
	//var instantWinnerGames []InstantWinnerGame
	var opened []InstantWinnerGame
	//var places []Place
/* 	AppDb.Find(&instantWinnerGames)
	context.JSON(http.StatusOK, instantWinnerGames) */
	var query InstWinGamesQuery
	var serverTime = time.Now()
	var result *gorm.DB = nil
	fmt.Printf("serverTime=%v", serverTime)
	//var instantWinnerGames []InstantWinnerGame
	
	if context.ShouldBindQuery(&query) != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": query})
		return
	}
	context.Bind(&query)
	//var c=context.Request.URL.Query()
/* 	if query.Date == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid Date"})
		return
	} */
	fmt.Printf("myquery2=%v", query)
	
/* 	var lat=ParseFloat(query.latitude,32)
	var lon=ParseFloat(query.longitude,32) */
 	var lat,_=strconv.ParseFloat(context.Param("lat"), 32)
	var lon,_=strconv.ParseFloat(context.Param("lon"), 32) 
	//fmt.Printf("lat=%v", lat)
	//id := c.Param("id")
/* 	AppDb.Find(&places)
	for i, s := range places {
		var pLat =s.Latitude
		var pLon =s.Longitude
		fmt.Println(i,pLat, pLon)
	} */
	//
	//latitude=11.1&longitude=22.3
	//AppDb.Find(&places)
/* 	var	lat,_ = strconv.ParseFloat("11.1", 32) 
	var lon,_ =strconv.ParseFloat("22.3", 32)   */
	fmt.Printf("lat=%v", lat)
	fmt.Printf("lon=%v", lon)
	
	//result =AppDb.Where("won = ? AND start_date < ? AND end_date > ?", false,serverTime, serverTime).Find(&opened)
	//AppDb.Where("won = ? AND start_date < ? AND end_date > ?", false,serverTime, serverTime).Find(&opened)
	fmt.Printf("step0")
	 //AppDb.Where("latitude < ? AND latitude > ? AND longitude < ? AND longitude > ?", 11.2,11.1, 22.4, 22.2).Find(&places)
	 //result = AppDb.Where("id > ?",0).Find(&places)
	 //AppDb.Where("id > ?",0).Find(&places)
	 fmt.Printf("step1")
	 //serverTime,_ = time.Parse("2006-01-02","2020-11-19")
	 //result = AppDb.Where("won = ? AND start_date < ? AND end_date > ?", false,serverTime, serverTime).Find(&opened)
	 //fmt.Printf(result)
	//serverTime,_ = time.Parse("2006-01-02","2020-11-19")
	
	//result =AppDb.Where("CreatedAt <= ?", serverTime).Find(&opened)
	//result =AppDb.Where("Won = ?", false).Find(&opened)
	//result =AppDb.Find(&opened)
	//context.JSON(http.StatusOK, opened)

	AppDb.Preload("instant_winner_games").Joins("JOIN places on places.id=instant_winner_games.place_id").Where("places.latitude < ? AND places.latitude > ? AND places.longitude < ? AND places.longitude > ? AND instant_winner_games.won = ? AND instant_winner_games.start_date < ? AND instant_winner_games.end_date > ?", (lat+0.1),(lat-0.1), (lon+0.1), (lon-0.1),false,serverTime, serverTime).Find(&opened)
		
	for _, ar := range opened {
		fmt.Println(ar.Name)
	}

	/* if err = db.Joins("JOIN artist_movies on artist_movies.artist_id=artists.id").
	Joins("JOIN movies on movies.id=artist_movies.movie_id").
	Joins("JOIN languages on movies.language_id=languages.id").
	Where("languages.name=?", "english").
	Group("artists.id").Preload("Movies"). */
		
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Model Not Found"})
		return
	}
	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	}
	context.JSON(http.StatusOK, opened)
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
