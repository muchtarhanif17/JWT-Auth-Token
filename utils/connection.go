package utils

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var Conn *sql.DB

func InitDB() error {
	var err error
	Conn, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/siqurban2025")
	if err != nil {
		return err
	}
	return Conn.Ping()
}
