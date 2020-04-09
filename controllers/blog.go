package controllers

import (
	blog "gin-rest-api/models"
	"gin-rest-api/util/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

// 获取最新增加的日志
func Latest(context *gin.Context) {

	blog, error := blog.LatestOne()
	if error != nil {
		response.Error(500, error, context)
		return
	}

	responseData := make(response.Response)
	responseData["blog"] = blog
	response.Success(responseData, context)

	return
}

// 获取日志列表
func List(context *gin.Context) {
	blogType := context.Query("type")
	page, _ := strconv.Atoi(context.Query("page"))

	list, error := blog.List(blogType, page)

	if error != nil {
		response.Error(500, error, context)
		return
	}

	responseData := make(response.Response)
	responseData["blogs"] = list
	response.Success(responseData, context)

	return
}

func Detail(context *gin.Context)  {
	id, _ := strconv.Atoi(context.Query("id"))

	blog, error := blog.Find(id)

	if error != nil {
		response.Error(500, error, context)
		return
	}

	responseData := make(response.Response)
	responseData["blog"] = blog
	response.Success(responseData, context)

	return
}