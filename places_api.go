package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitPlacesApi(router *gin.Engine, db *gorm.DB) {
	router.GET("/api/places", ListPlaces)
	router.GET("/api/places/:id", ReadPlace)
	router.POST("/api/places", CreatePlace)
	router.PUT("/api/places/:id", UpdatePlace)
	router.DELETE("/api/places/:id", DeletePlace)
}
func ListPlaces(context *gin.Context) {
	var places []Place
	ListModels(context, &places)
}
func ReadPlace(context *gin.Context) {
	var place Place
	ReadModel(context, &place)
}
func CreatePlace(context *gin.Context) {
	var place Place
	CreateModel(context, &place)
}
func UpdatePlace(context *gin.Context) {
	var place Place
	UpdateModel(context, &place, setPlaceId)
}
func DeletePlace(context *gin.Context) {
	DeleteModel(context, &Place{})
}
func setPlaceId(id uint64, model interface{}) {
	place := model.(*Place)
	place.ID = id
}
