package models

type Blog struct {
	id int
	title string
	date string
	summary string
	content string
}

func LatestOne() *Blog {

	blog := new(Blog)
	blog.id = 1
	blog.title = "mysql 性能优化策略"
	blog.date = "2019-03-23 22:23:22"
	blog.summary = "msyql 性能优化策略"
	blog.content = "msyql 性能优化策略"

	return blog
}