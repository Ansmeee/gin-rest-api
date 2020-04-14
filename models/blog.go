package models

import (
	"errors"
	"fmt"
	"gin-rest-api/lib/db"
	"gin-rest-api/util"
	"strings"
)

type Blog struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Summary string `json:"summary"`
	Content string `json:"content"`
	Class   string `json:"class"`
	Ctime   string `json:"ctime"`
}

var BlogType = map[string]string{
	"blog":  "日志分享",
	"study": "学习笔记",
	"photo": "摄影日记",
}

func Create(blog *Blog) (int, error) {
	con, conError := db.Connection()
	if conError != nil {
		return 0, conError
	}

	defer con.Close()

	query := "INSERT INTO `blog` (`title`, `summary`, `content`, `class``) values (?, ?, ?, ?)"
	queryRes, queryError := db.Insert(con, query, blog.Title, blog.Summary, blog.Content, BlogType[blog.Class])
	if queryError != nil {
		return 0, queryError
	}

	return queryRes, nil
}

func Find(id int) (*Blog, error) {
	con, conError := db.Connection()
	if conError != nil {
		return nil, conError
	}

	defer con.Close()

	query := "SELECT * FROM `blog` WHERE `id` = ? limit 1"
	queryRes, queryError := db.Query(con, query, id)
	if queryError != nil {
		return nil, queryError
	}

	blog := new(Blog)
	for queryRes.Next() {
		error := queryRes.Scan(&blog.Id, &blog.Title, &blog.Summary, &blog.Content, &blog.Class, &blog.Ctime)
		if error != nil {
			util.Error(error, "Scan Data Failed")
			return nil, errors.New("数据获取失败，请重试")
		}
	}

	defer queryRes.Close()

	return blog, nil
}

func LatestOne() (*Blog, error) {
	con, conError := db.Connection()
	if conError != nil {
		return nil, conError
	}

	defer con.Close()

	query := "SELECT * FROM `blog` ORDER BY `id` DESC LIMIT 1"
	queryRes, queryErr := db.Query(con, query)

	if queryErr != nil {
		return nil, queryErr
	}

	blog := new(Blog)
	for queryRes.Next() {
		error := queryRes.Scan(&blog.Id, &blog.Title, &blog.Summary, &blog.Content, &blog.Class, &blog.Ctime)
		if error != nil {
			util.Error(error, "Scan Data Failed")
			return nil, errors.New("数据获取失败，请重试")
		}
	}

	defer queryRes.Close()

	return blog, nil
}

type Class struct {
	Name  string `json:"name"`
	Total int    `json:"total"`
}

var class Class
var classes []*Class // 定义一个切片
func ClassTotal() ([]*Class, error) {

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

	classes = []*Class{}
	for queryRes.Next() {
		class := new(Class)
		error := queryRes.Scan(&class.Name, &class.Total)
		if error != nil {
			util.Error(error, "Scan Data Failed")
			return nil, errors.New("数据获取失败，请重试")
		}

		classes = append(classes, class)
	}

	defer queryRes.Close()

	return classes, nil
}

var lists []*Blog

func List(page int, blogType string) ([]*Blog, error) {
	con, conError := db.Connection()

	if conError != nil {
		return nil, conError
	}

	defer con.Close()

	ids := make([]interface{}, 0)

	wheres := make([]string, 0)
	idsQuery := "SELECT `id` FROM `blog`"

	currentType, isPresent := BlogType[blogType]
	if isPresent {
		wheres = append(wheres, fmt.Sprintf("`class` = %s", currentType))
	}

	if len(wheres) > 0 {
		idsQuery = idsQuery + " WHERE " + strings.Join(wheres, "AND")
	}

	idsRes, idsError := db.Query(con, idsQuery)

	if idsError != nil {
		return nil, idsError
	}

	for idsRes.Next() {
		var id int
		error := idsRes.Scan(&id)
		if error != nil {
			util.Error(error, "Scan Data Failed")
			return nil, errors.New("数据获取失败，请重试")
		}

		ids = append(ids, id)
	}

	total := len(ids)
	lists := make([]*Blog, 0)
	if total > 0 {

		query := "SELECT id, title, summary, class, ctime FROM `blog` WHERE `id`"
		query = db.PrepareInArgs(query, ids)
		query += " ORDER BY `id` DESC"

		queryRes, queryError := db.Query(con, query, ids...)

		if queryError != nil {
			return nil, queryError
		}

		for queryRes.Next() {
			blog := new(Blog)
			error := queryRes.Scan(&blog.Id, &blog.Title, &blog.Summary, &blog.Class, &blog.Ctime)
			if error != nil {
				util.Error(error, "Scan Data Failed")
				return nil, errors.New("数据获取失败，请重试")
			}

			lists = append(lists, blog)
		}

		queryRes.Close()
	}

	return lists, nil
}
