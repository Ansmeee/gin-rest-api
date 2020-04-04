package models

import (
	"fmt"
	"gin-rest-api/lib/db"
	"gin-rest-api/util"
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
	data := make(map[string]interface{})

	data["id"] = blog.Id
	data["title"] = blog.Title
	data["content"] = blog.Content
	data["class"] = blog.Class
	data["ctime"] = blog.Ctime

	return data
}

func LatestOne() (*Blog, error)  {
	var blog Blog

	con, conError := db.Connection()
	if conError != nil {
		return &blog, conError
	}

	defer con.Close()

	query := "SELECT * FROM `blog` ORDER BY `id` DESC LIMIT 1";
	queryRes, queryErr := db.Query(con, query)

	if queryErr != nil {
		return &blog, queryErr
	}

	for queryRes.Next() {
		error := queryRes.Scan(&blog.Id, &blog.Title, &blog.Content, &blog.Class, &blog.Ctime)
		if error != nil {
			util.Error(error, "failed")
		}
	}

	defer queryRes.Close()

	return &blog, nil
}

func ClassTotal() (map[string]interface{}, error) {
	data := make(map[string]interface{})

	con, conError := db.Connection()
	if conError != nil {
		return nil, conError
	}

	defer con.Close()

	query := "SELECT `class`, COUNT(`id`) AS total FROM `blog` GROUP BY `class`"
	queryRes, queryErr := db.Query(con, query)

	if queryErr != nil {
		return nil, queryErr
	}

	type Class struct {
		Name string
		Total int
	}

	var class Class
	var classes []Class

	for queryRes.Next(){
		error := queryRes.Scan(&class.Name, &class.Total)
		if error != nil {
			util.Error(error, "failed")
		}

		classes = append(classes, class)
	}

	defer queryRes.Close()

	data["blogs"] = classes
	fmt.Println(classes)
	return data, nil
}