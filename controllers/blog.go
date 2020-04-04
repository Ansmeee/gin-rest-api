package controllers

import (
	"gin-rest-api/lib"
	"gin-rest-api/lib/db"
	"gin-rest-api/lib/response"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func LatestBlog(context *gin.Context) {

	con := db.Connection()

	rows, sqlError := con.Query("select id, name from `user` where id = ?", 1)

	if sqlError != nil {
		lib.Audit("sql error", sqlError)
		response.Error(500, "查询失败", context)
		return
	}

	defer rows.Close()

	type Blog = map[string]interface{}

	blog := make(Blog)

	response.Success(blog, context)
	return
}
