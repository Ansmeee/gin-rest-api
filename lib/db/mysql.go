package db

import (
	"database/sql"
	"gin-rest-api/lib"
)

func Connection() *sql.DB {

	driver 		:= "mysql"
	userName 	:= "dev"
	passWord 	:= "ansme007"
	address 	:= "192.168.33.10"
	port 		:= "3306"
	database 	:= "blog"

	dns := userName + ":" + passWord + "@tcp(" + address + ":" + port + ")/" + database

	connection, conError := sql.Open(driver, dns)

	if conError != nil {
		lib.Audit("Database connection failed", conError)
		return nil
	}

	return connection
}
