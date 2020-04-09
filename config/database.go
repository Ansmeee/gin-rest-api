package config

// 数据库配置
var Dns = struct{ Driver, UserName, PassWord, Address, Port, Database string }{
	"mysql",
	"dev",
	"ansme007.blog",
	"127.0.0.1",
	"3306",
	"blog",
}