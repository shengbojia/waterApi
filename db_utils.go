package waterApi

import (
	"database/sql"
	"fmt"
)

func CreateDatabase() (*sql.DB, error) {
	serverName := "aa1nmgs2nqru5jj.cx4nbvplvoaq.us-east-2.rds.amazonaws.com:3306"
	user := "test"
	password := "1234abcd"
	dbName := "ebdb"

	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true&multiStatements=true", user, password, serverName, dbName)

	db, err := sql.Open("mysql", connectionString)

	return db, err
}