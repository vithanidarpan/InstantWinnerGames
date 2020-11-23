package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetIDOrError(context *gin.Context) (uint64, bool) {
	var id uint64
	var err error
	if id, err = strconv.ParseUint(context.Param("id"), 10, 64); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return 0, false
	}
	return id, true
}
