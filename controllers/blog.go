package controllers

import (
	"gin-rest-api/models"
	"gin-rest-api/util/response"
	"github.com/gin-gonic/gin"
)

func LatestBlog(context *gin.Context) {

	blog, error := models.LatestOne()
	if error != nil {
		response.Error(500, "请求失败", context)
		return
	}

	data := blog.MakeData()

	response.Success(data, context)

}
