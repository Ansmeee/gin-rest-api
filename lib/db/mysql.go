package db

import (
	"database/sql"
	"errors"
	"fmt"
	database "gin-rest-api/config"
	"gin-rest-api/util"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

const ConErr = "数据库连接失败，请重试"
const QueryErr = "数据查询失败，请重试"

func Connection() (*sql.DB, error) {

	userName := database.Mysql.UserName
	passWord := database.Mysql.PassWord
	address := database.Mysql.Address
	port := database.Mysql.Port
	database := database.Mysql.Database

	dns := userName + ":" + passWord + "@tcp(" + address + ":" + port + ")/" + database

	connection, conError := sql.Open("mysql", dns)

	if conError != nil {
		util.Error(conError, "Database Connection Failed")
		return nil, errors.New(ConErr)
	}

	return connection, nil
}

func Query(con *sql.DB, query string, args ...interface{}) (*sql.Rows, error) {

	queryRows, queryError := con.Query(query, args...)

	if queryError != nil {
		util.Error(queryError, "Query Failed")
		return queryRows, errors.New(QueryErr)
	}

	return queryRows, nil
}

func PrepareInArgs(query string, args []interface{}) string {

	in := " IN (%s)"
	query = query + in
	inArgs := strings.Repeat("?,", len(args)-1) + "?"
	inQuery := fmt.Sprintf(query, inArgs)

	return inQuery
}
