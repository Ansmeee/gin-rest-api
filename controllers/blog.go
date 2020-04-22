package controllers

import (
	"errors"
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
	params := new(blog.Params)

	params.BlogType = context.Query("type")
	params.Page, _ = strconv.Atoi(context.Query("page"))
	params.Keywords = context.Query("keywords")

	list, error := blog.List(params)

	if error != nil {
		response.Error(500, error, context)
		return
	}

	responseData := make(response.Response)
	responseData["blogs"] = list
	response.Success(responseData, context)

	return
}

func Detail(context *gin.Context) {
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

func Create(context *gin.Context) {
	blogForm := makeForm(context)

	validateErr := validateForm(blogForm)
	if validateErr != nil {
		response.Error(400, validateErr, context)
		return
	}

	res, error := blog.Create(blogForm)
	if error == nil {
		responseData := make(response.Response)
		responseData["id"] = res
		response.Success(responseData, context)
	}

	response.Error(500, error, context)
	return
}

func makeForm(context *gin.Context) *blog.Blog {
	blogForm := new(blog.Blog)
	blogForm.Title = context.PostForm("title")
	blogForm.Summary = context.PostForm("summary")
	blogForm.Class = context.PostForm("class")
	blogForm.Content = context.PostForm("content")

	return blogForm
}

func validateForm(form *blog.Blog) error {
	if form.Title == "" {
		return errors.New("验证失败：标题不能为空")
	}

	if form.Class == "" {
		return errors.New("验证失败：分类选择不能为空")
	}

	_, isPresent := blog.BlogType[form.Class]
	if !isPresent {
		return errors.New("验证失败：选择的分类不存在")
	}

	if form.Content == "" {
		return errors.New("验证失败：内容不能为空")
	}

	return nil
}
