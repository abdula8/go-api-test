package main

import (
	"database/sql"
	"fmt"
	"os"

	mysqldriver "github.com/go-sql-driver/mysql"
)

var _connection *sql.DB

func init() {
	mysqlConfig := &mysqldriver.Config{
		User:                 os.Getenv("MYSQL_USER"),
		Passwd:               os.Getenv("MYSQL_PASS"),
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%s:%s", os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT")),
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	connection, err := sql.Open("mysql", mysqlConfig.FormatDSN())
	if err != nil {
		panic(err)
	}

	_, err = connection.Exec(`CREATE DATABASE IF NOT EXISTS internship`)
	if err != nil {
		panic(err)
	}

	mysqlConfig.DBName = `internship`

	if err = connection.Close(); err != nil {
		panic(err)
	}
	// Check and grant necessary privileges
	if err := grantPrivileges(connection); err != nil {
		panic(err)
	}

	connection, err = sql.Open("mysql", mysqlConfig.FormatDSN())
	if err != nil {
		panic(err)
	}

	_connection = connection

	schema()
}

func schema() {
	_, err := _connection.Exec(`CREATE TABLE IF NOT EXISTS stuff (
		id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
		created_at DATETIME NOT NULL
	)`)

	if err != nil {
		panic(err)
	}
}

func grantPrivileges(db *sql.DB) error {
	// Check if the user 'go' has the necessary privileges
	rows, err := db.Query("SHOW GRANTS FOR ?@?", os.Getenv("MYSQL_USER"), "%")
	if err != nil {
		return err
	}
	defer rows.Close()

	// If no rows are returned, grant necessary privileges
	if !rows.Next() {
		_, err := db.Exec("GRANT ALL PRIVILEGES ON ?.* TO ?@?;", "internship", os.Getenv("MYSQL_USER"), "%")
		if err != nil {
			return err
		}
		_, err = db.Exec("FLUSH PRIVILEGES;")
		if err != nil {
			return err
		}
	}

	return nil
}