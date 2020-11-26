package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ModelIdGetter func(model interface{}) uint64
type ModelIdSetter func(id uint64, model interface{})

func ListModels(context *gin.Context, models interface{}) {
	AppDb.Find(models)
	context.JSON(http.StatusOK, models)
}

func FetchListModels(models interface{}) {
	AppDb.Find(models)
}

func ListModelsWith(context *gin.Context, models interface{}, preloads []string) {
	var result *gorm.DB = nil
	for _, preload := range preloads {
		if result == nil {
			result = AppDb.Preload(preload)
		} else {
			result = result.Preload(preload)
		}
	}
	if result == nil {
		result = AppDb.Find(models)
	} else {
		result = result.Find(models)
	}
	context.JSON(http.StatusOK, models)
}
func ReadModel(context *gin.Context, model interface{}) {
	var id uint64
	var ok bool
	if id, ok = GetIDOrError(context); !ok {
		return
	}
	result := AppDb.First(model, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Model Not Found"})
		return
	}
	context.JSON(http.StatusOK, model)
}
func ReadModelWith(context *gin.Context, model interface{}, preloads []string) {
	var id uint64
	var ok bool
	if id, ok = GetIDOrError(context); !ok {
		return
	}
	var result *gorm.DB = nil
	for _, preload := range preloads {
		if result == nil {
			result = AppDb.Preload(preload)
		} else {
			result = result.Preload(preload)
		}
	}
	if result == nil {
		result = AppDb.First(model, id)
	} else {
		result = result.First(model, id)
	}
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Model Not Found"})
		return
	}
	context.JSON(http.StatusOK, model)
}
func CreateModel(context *gin.Context, model interface{}) {
	if err := context.BindJSON(model); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if AppDb.Create(model) == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create Model"})
		return
	}
	context.JSON(http.StatusOK, model)
}

func FetchCreateModel(model interface{}) error {
	var err error
	if AppDb.Create(model) == nil {
		return err
	}

	return nil
}

func UpdateModel(context *gin.Context, model interface{},
	idSetter ModelIdSetter) {
	var id uint64
	var ok bool
	if id, ok = GetIDOrError(context); !ok {
		return
	}
	if err := context.BindJSON(model); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	idSetter(id, model)
	if AppDb.Save(model) == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot Update Model"})
		return
	}
	context.JSON(http.StatusOK, model)
}
func FetchUpdateModel(model interface{}, id uint64, idSetter ModelIdSetter) error {
	var err error
	idSetter(id, model)
	if AppDb.Save(model) == nil {
		return err
	}

	return nil
}
func DeleteModel(context *gin.Context, model interface{}) {
	var id uint64
	var ok bool
	if id, ok = GetIDOrError(context); !ok {
		return
	}
	if AppDb.Delete(model, id) == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot Delete Model"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"success": "Model Deleted"})
}
