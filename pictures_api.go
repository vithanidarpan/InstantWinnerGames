package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitPicturesApi(router *gin.Engine, db *gorm.DB) {
	router.GET("/api/pictures", ListPictures)
	router.GET("/api/pictures/:id", ReadPicture)
	router.POST("/api/pictures", CreatePicture)
	router.PUT("/api/pictures/:id", UpdatePicture)
	router.DELETE("/api/pictures/:id", DeletePicture)
}
func ListPictures(context *gin.Context) {
	var pictures []Picture
	ListModels(context, &pictures)
}
func ReadPicture(context *gin.Context) {
	var picture Picture
	ReadModel(context, &picture)
}
func CreatePicture(context *gin.Context) {
	var picture Picture
	CreateModel(context, &picture)
}
func UpdatePicture(context *gin.Context) {
	var picture Picture
	UpdateModel(context, &picture, setPictureId)
}
func DeletePicture(context *gin.Context) {
	DeleteModel(context, &Picture{})
}
func setPictureId(id uint64, model interface{}) {
	picture := model.(*Picture)
	picture.ID = id
}
