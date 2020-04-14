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

const ConErr = "连接失败，请重试"
const QueryErr = "查询失败，请重试"

func Connection() (*sql.DB, error) {
	connection, conError := sql.Open(database.Dns.Driver, makeDNS())

	if conError != nil {
		util.Error(conError, "Database Connection Failed")
		return nil, errors.New(ConErr)
	}

	return connection, nil
}

func Insert(con *sql.DB, query string, args ...interface{}) (int, error) {

	return 0, nil
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

func makeDNS() string {
	userName := database.Dns.UserName
	passWord := database.Dns.PassWord
	address := database.Dns.Address
	port := database.Dns.Port
	database := database.Dns.Database

	return userName + ":" + passWord + "@tcp(" + address + ":" + port + ")/" + database
}
