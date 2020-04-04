package db

import (
	"database/sql"
	"gin-rest-api/util"
	_ "github.com/go-sql-driver/mysql"
)

func Connection() (*sql.DB, error) {

	driver 		:= "mysql"
	userName 	:= "dev"
	passWord 	:= "ansme007"
	address 	:= "192.168.33.10"
	port 		:= "3306"
	database 	:= "blog"

	dns := userName + ":" + passWord + "@tcp(" + address + ":" + port + ")/" + database

	connection, conError := sql.Open(driver, dns)

	if conError != nil {
		util.Error(conError, "Database connection failed")
		return nil, conError
	}

	return connection, nil
}

func Query(con *sql.DB, query string) (*sql.Rows, error) {
	queryRows, queryError := con.Query(query)

	if queryError != nil {
		util.Error(queryError, "Query failed")
	}

	return queryRows, queryError
}
