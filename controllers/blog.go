package controllers

import (
	response "gin-rest-api/lib"
	"github.com/gin-gonic/gin"
)

func LatestBlog(context *gin.Context)  {
	response := make(response.Response)
	response["code"] = 400
	response["message"] = "验证失败：请输入用户名！"
	context.JSON(200, response)
	return
}

