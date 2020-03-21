package controllers

import (
	file "gin-rest-api/lib"
	"github.com/gin-gonic/gin"
)


func Signup(context *gin.Context)  {
	// request form
	userName := context.DefaultPostForm("userName", "")
	userPass := context.DefaultPostForm("userPass", "")

	// init response data
	response := make(map[string]interface{})


	// validate user name and password
	if userName == "" {
		response["code"] = 400
		response["message"] = "请输入用户名！"
		context.JSON(200, response)
		return
	}

	if userPass == "" {
		response["code"] = 400
		response["message"] = "请输入密码！"
		context.JSON(200, response)
		return
	}

	// save user info
	createRes := file.CreateUserInfo(userName, userPass)
	if createRes == false {
		response["code"] = 500
		response["message"] = "注册失败！"
		context.JSON(200, response)
		return
	}

	// response with success
	response["message"] = "注册成功！"
	context.JSON(200, response)
	return
}