package controllers

import (
	"../models"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

type updateParams struct {
	Approved []int64 `json:"approved"`
	Declined []int64 `json:"declined"`
}

func CommentsUpdate(ctx *gin.Context) {
	var data updateParams
	err := ctx.BindJSON(&data)
	if err != nil {
		ctx.AbortWithError(422, err)
		return
	}

	err = models.DB.Model(&models.Comment{}).
		Where("id in (?)", data.Approved).
		Update("approval", sql.NullBool{Bool: true, Valid: true}).
		Error
	if err != nil {
		ctx.AbortWithError(422, err)
		return
	}
	err = models.DB.Model(&models.Comment{}).
		Where("id in (?)", data.Declined).
		Update("approval", sql.NullBool{Bool: false, Valid: true}).
		Error
	if err != nil {
		ctx.AbortWithError(422, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"approved": data.Approved,
		"declined": data.Declined,
	})
}

func CommentsIndex(ctx *gin.Context) {
	comments := []*models.Comment{}
	err := models.DB.Model(&models.Comment{}).Where("approval IS NULL").Order("id asc").Limit(1000).Find(&comments).Error
	if err != nil {
		ctx.AbortWithError(422, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   comments,
	})
}
