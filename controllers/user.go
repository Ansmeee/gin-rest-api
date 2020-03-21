package controllers

import (
	file "gin-rest-api/lib"
	user "gin-rest-api/models"
	"github.com/gin-gonic/gin"
)


type response = map[string]interface{}

func Signup(context *gin.Context)  {
	// request form
	userName := context.DefaultPostForm("userName", "")
	userPass := context.DefaultPostForm("userPass", "")

	// init response data
	response := make(response)

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

	// response with success
	if createRes == true {
		response["message"] = "注册成功！"
		context.JSON(200, response)
		return

	}


	// response with error
	response["code"] = 500
	response["message"] = "注册失败！"
	context.JSON(200, response)
	return
}

func Signin(context *gin.Context)  {
	// get request form from post form
	userName := context.DefaultPostForm("userName", "")
	userPass := context.DefaultPostForm("userPass", "")

	// init response data
	response := make(response)
	if userName == "" {
		response["code"] = 400
		response["message"] = "验证失败：请输入用户名！"
		context.JSON(200, response)
		return
	}

	if userPass == "" {
		response["code"] = 400
		response["message"] = "验证失败：请输入密码！"
		context.JSON(200, response)
		return
	}

	// check user exists first
	userExists := file.GetUser(userName)
	if !userExists {
		response["code"] = 400
		response["message"] = "验证失败：该用户不存在！"
		context.JSON(200, response)
		return
	}


	// encrypt userPass second
	encryptUserPass := string(user.EncryptUserPass(userPass))

	// get user origin password third
	originUserPass := file.GetUserPass(userName)

	// validate userName and userPass fourth
	if encryptUserPass == originUserPass {
		response["code"] = 200
		response["message"] = "登陆成功！"
		context.JSON(200, response)
		return
	}

	// return error
	response["code"] = 500
	response["message"] = "用户名或密码错误！"
	context.JSON(200, response)
	return
}