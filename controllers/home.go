package controllers

import "github.com/gin-gonic/gin"

func Index(context *gin.Context)  {


	context.JSON(500, gin.H{})
}