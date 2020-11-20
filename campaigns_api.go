package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitCampaignsApi(router *gin.Engine, db *gorm.DB) {
	router.GET("/api/campaigns", ListCampaigns)
	router.GET("/api/campaigns/:id", ReadCampaign)
	router.POST("/api/campaigns", CreateCampaign)
	router.PUT("/api/campaigns/:id", UpdateCampaign)
	router.DELETE("/api/campaigns/:id", DeleteCampaign)
}
func ListCampaigns(context *gin.Context) {
	var campaigns []Campaign
	ListModels(context, &campaigns)
}
func ReadCampaign(context *gin.Context) {
	var campaign Campaign
	ReadModel(context, &campaign)
}
func CreateCampaign(context *gin.Context) {
	var campaign Campaign
	CreateModel(context, &campaign)
}
func UpdateCampaign(context *gin.Context) {
	var campaign Campaign
	UpdateModel(context, &campaign, setCampaignId)
}
func DeleteCampaign(context *gin.Context) {
	DeleteModel(context, &Campaign{})
}
func setCampaignId(id uint64, model interface{}) {
	campaign := model.(*Campaign)
	campaign.ID = id
}
