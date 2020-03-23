package controllers

import (
	"database/sql"
	"fmt"
	response "gin-rest-api/lib"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func LatestBlog(context *gin.Context) {

	con, err := sql.Open("mysql", "dev:ansme007@tcp(192.168.33.10:3306)/blog")

	if err != nil {
		fmt.Printf("connect to mysql failed: %s", err)
		return
	}

	rows, errors := con.Query("select * from `user` where id = ?", 1)

	if errors != nil {
		fmt.Printf("sql error: %s", errors)
		return
	}

	fmt.Println(rows, errors)
	type Blog = map[string]interface{}

	blog := make(Blog)

	blog["id"] 		= 1
	blog["title"] 	= "mysql 性能优化策略"
	blog["date"] 	= "2019-03-23 22:23:22"
	blog["type"] 	= "学习笔记"
	blog["summary"] = "msyql 性能优化策略"
	blog["content"] = "msyql 性能优化策略"

	response := make(response.Response)
	response["code"] = 200
	response["data"] = blog
	context.JSON(200, response)
	return
}
