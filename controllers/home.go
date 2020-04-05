package controllers

import (
	"gin-rest-api/models"
	"gin-rest-api/util/response"
	"github.com/gin-gonic/gin"
)

func Index(context *gin.Context)  {
	context.JSON(500, gin.H{})
}


func Info(context *gin.Context)  {
	blogClasses, error := models.ClassTotal()

	if error != nil {
		response.Error(500, "获取失败", context)
		return
	}

	responseData := make(response.Response)
	responseData["blogs"] = blogClasses
	response.Success(responseData, context)

	return
}