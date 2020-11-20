package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitUsersApi(router *gin.Engine, db *gorm.DB) {
	router.GET("/api/users", ListUsers)
	router.GET("/api/users/:id", ReadUser)
	router.POST("/api/users", CreateUser)
	router.PUT("/api/users/:id", UpdateUser)
	router.DELETE("/api/users/:id", DeleteUser)
}
func ListUsers(context *gin.Context) {
	var users []User
	ListModels(context, &users)
}
func ReadUser(context *gin.Context) {
	var user User
	ReadModel(context, &user)
}
func CreateUser(context *gin.Context) {
	var user User
	CreateModel(context, &user)
}
func UpdateUser(context *gin.Context) {
	var user User
	UpdateModel(context, &user, setUserId)
}
func DeleteUser(context *gin.Context) {
	DeleteModel(context, &User{})
}
func setUserId(id uint64, model interface{}) {
	user := model.(*User)
	user.ID = id
}
