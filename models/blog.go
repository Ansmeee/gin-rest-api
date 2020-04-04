package models

import (
	"gin-rest-api/lib/db"
	"gin-rest-api/util"
	"gin-rest-api/util/response"
)

type Blog struct {
	Id int
	Title string
	Content string
	Class string
	Ctime string
}

type Bloger interface {
	MakeData() map[string]interface{}
}

func (blog *Blog) MakeData() map[string]interface{} {
	data := response.Response{}

	data["id"] = blog.Id
	data["title"] = blog.Title
	data["content"] = blog.Content
	data["class"] = blog.Class
	data["ctime"] = blog.Ctime

	return data
}

func LatestOne() (Blog, error)  {
	var blog Blog

	con := db.Connection()
	if con == nil {
		return blog, nil
	}

	defer con.Close()

	query := "SELECT * FROM `blog` ORDER BY `id` DESC LIMIT 1";
	queryRes, queryErr := con.Query(query)

	if queryErr != nil {
		util.Error(queryErr, "query failed")
	}


	for queryRes.Next() {
		error := queryRes.Scan(&blog.Id, &blog.Title, &blog.Content, &blog.Class, &blog.Ctime)
		if error != nil {
			util.Error(error, "failed")
		}
	}

	defer queryRes.Close()

	return blog, nil
}