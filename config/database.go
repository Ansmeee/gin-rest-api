package config

var Mysql = struct{ UserName, PassWord, Address, Port, Database string }{
	"dev",
	"ansme007.blog",
	"127.0.0.1",
	"3306",
	"blog",
}