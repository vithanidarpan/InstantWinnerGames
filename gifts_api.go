package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitGiftsApi(router *gin.Engine, db *gorm.DB) {
	router.GET("/api/gifts", ListGifts)
	router.GET("/api/gifts/:id", ReadGift)
	router.POST("/api/gifts", CreateGift)
	router.PUT("/api/gifts/:id", UpdateGift)
	router.DELETE("/api/gifts/:id", DeleteGift)
}
func ListGifts(context *gin.Context) {
	var gifts []Gift
	ListModels(context, &gifts)
}
func ReadGift(context *gin.Context) {
	var gift Gift
	ReadModel(context, &gift)
}
func CreateGift(context *gin.Context) {
	var gift Gift
	CreateModel(context, &gift)
}
func UpdateGift(context *gin.Context) {
	var gift Gift
	UpdateModel(context, &gift, setGiftId)
}
func DeleteGift(context *gin.Context) {
	DeleteModel(context, &Gift{})
}
func setGiftId(id uint64, model interface{}) {
	gift := model.(*Gift)
	gift.ID = id
}
